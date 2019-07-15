Security Misconfiguration

    Good security requires having a secure configuration defined and deployed for the application, frameworks, application server, web server, database server, and platform. Secure settings should be defined, implemented, and maintained, as defaults are often insecure. Additionally, software should be kept up to date.

One of the easiest mistakes to make when you deploy a web application is to ship a rough prototype to production. The compromises we make when building a proof-of-concept quickly for management are easily discovered by a handful of quick inspections once they go live. In a prototype, you might cut corners or ignore security to prove out a concept; if you’re building an actual release, bake security in from the start!

Web requests report not only the server but also its version in headers to clients. PHP noisily exposes what version it’s running with every request. Often, the application layer we’re using inserts a <meta> into HTML content so the developers can readily track statistics downstream.

These might all seem innocuous, but they help highlight when we’re running old or insecure applications in production and help attackers choose their next target. Listing the version on PHP your server might feel like a badge of honor if you’re launching with a bleeding-edge version in production. But if X-Powered-by is instead bragging you’ve got an unpatched deployment of PHP 5.3.3 on the server, you’re shooting yourself in the foot. Of course, it’s even more important to keep your server software patched and up-to-date in the first place.

Similarly, the runtime configuration of each application in your stack is relatively easy to misconfigure. Most of us use one configuration in development—or even on a staging server—that makes it easier for us to debug problems as they come up. However, his configuration is a horrible mismatch for a production environment.

    By default, PHP ships with recommended settings in php.ini for both production and development. Save yourself time by leveraging the defaults already configured by the community to keep your stack secure.

A stack dump triggered by an uninitialized array index is a horrible user experience. It’s also a horrendous way to expose the internals of your application to the world if display_errors is enabled.

Similarly, some servers can allow the users to do things they otherwise shouldn’t be able to do at all. A server allowing directory traversal will permit an attacker to index locations they shouldn’t be able to access. A database that logs queries to disk for debugging might erroneously index otherwise secure information. To understand how deeper issues can then expose this information, see the chapter on ASR3.
How Would This Look in Production?

Every application has a different configuration file. As they each serve different purposes within the stack, there are different ways each can be configured that will led to potential issues within your application.

As previously mentioned, many of these settings and feature flags are things you might use in development to track down issues. However, it’s vital you understand what each flag does, why they present unnecessary risk to your application, and how to set things properly for a production instance.

    Note: This chapter does not expect you to become an expert systems administrator or know the ins and outs of each service running on the server. There are experts who specialize in this work, and your team is fortunate if you have even one qualified systems engineer. Every engineer should, however, at least have a passing familiarity with these services, their configuration, and the potential pitfalls faced by your application if things are not set properly.

Web servers (NGINX and Apache)

While it would be my preference that the entire world configured their server hosts the same way, that is not likely to happen. Some applications use NGINX as a proxy to PHP-FPM. Some use Apache with PHP embedded. Others use PHP’s built-in web server running inside a Docker container. There is no perfect host configuration, but there are specific features and options which are often used on servers that present configuration issues with your application. While we can’t address all of the various configuration details available, there are five specific to NGINX and Apache you should be aware of:

Both web servers will helpfully add server tokens with their names and versions in the headers returned with every request—Apache will even print version information in the HTML footer of default error pages. These values can be useful for debugging a staging or development server as they help narrow down issues to specific versions of the server application. However, they also expose the server version to inspection by potential attackers with lists of known exploits against older versions. If, for whatever reason, you’re unable to keep your servers running an up-to-date and supported release, this information could help an attacker breach your system.

When configuring a multitenant server—hosting multiple applications within the same VM and forwarding traffic to one or another virtual host—you’ll likely configure the server name directive for each virtual host. This helps NGINX and Apache act as a load balancer so they can serve the correct application in response to requests. On single-purpose servers, there’s no need to set the server name, as the “default” virtual host will resolve for any address that hits the server. Unfortunately, using the default host means the $_SERVER superglobal in PHP will load some of its information from the request headers sent by the user making the request. If your application relies at all on this superglobal (and many do), then users can inject data into the application by way of customized request headers!

NGINX and Apache can both be configured to permit directory traversal on the web server. In older file servers (usually running some form of FTP), this would allow end users to navigate from one directory to another in search of a specific file that would ultimately be served by the application itself. However, giving users the ability to traverse through your web root and view lists of files could expose them to content they are not permitted to access or, more dangerously, traverse out of the web root deeper into the server itself.

Until recently, SSL certificates have been expensive, difficult to obtain, and tricky to install in production. This has lead to developers quickly spinning up production servers that present content unencrypted over HTTP. Poor documentation about the underlying cryptography options or proper configuration has also lead to many developers who do add SSL to their applications doing so incorrectly and exposing their application to further security issues.

Error handling is a component of software development many engineers skip right over. PHP developers can silence application errors entirely using runtime configuration (settings in their php.ini file), so they often form a habit of not thinking about edge cases that might occur in their application and otherwise impact user behavior. Worse, some developers will use the @ to silence errors or warnings in the code. Unfortunately, not handling errors is incredibly unsafe. An unhandled fatal error can trigger a crash on the server, a dump of sensitive memory, or otherwise cause the application to behave in unexpected and unanticipated ways. Some silent errors, while they don’t trigger a crash, might allow the application to continue behaving even though it should have stopped entirely.
PHP

Obviously, we want to ensure our PHP service is configured correctly to serve our PHP application securely. Whether you’re running PHP-FPM and serving content through NGINX as a proxy or using the PHP module built into Apache, your application will rely on a proper php.ini file to behave as expected.

There are eleven directives within php.ini that we need to look at to start:

expose_php is a Boolean flag that determines whether or not the X-Powered-By header is presented in a request. Running curl -I https://mysite.com shows whether or not this is enabled. On a generic Ubuntu 12.04 installation, the header will likely report PHP/5.5.9-1ubuntu4.21.

From this information, a would-be attacker can identify which operating system you’re running and potentially target a known vulnerability if things are unmatched. They can also directly see which version of PHP is powering the server. For example, PHP 5.5.9 was released in February of 2014, has several documented vulnerabilities, and is the default installation in Ubuntu 12.04. Equally unfortunate is the PHP 5.5 branch was end-of-life in July of 2016, so even updating to the latest version in that branch (5.5.37) wouldn’t be enough to protect such a server.

allow_url_fopen and allow_url_include are both Boolean options developers use to interact with remote scripts. The first variant allows for opening remote resources as if they’re local files via fopen(). In some cases, this might be used to interact with remote APIs or download files into a local cache. However, fopen() is also used often by attackers to download remote backdoor scripts into a server. Though its availability alone isn’t a vulnerability, it’s a tool available to anyone who gains runtime access of the server to escalate their level of control or access.

The second variant allows PHP’s standard include() and require() mechanisms to load remote PHP scripts for execution on the server. Read that sentence one more time, and I’m sure you’ll immediately know why it’s a bad idea.

Loading scripts from a remote location that can then execute on your server opens you up to malicious third parties taking control of your machine!

display_errors is a Boolean flag that determines whether or not PHP will print errors (and warnings and notices) to the front-end of your site. This is a very useful constant to set to on in development as you’ll be able to identify various issues with your code that, while they don’t halt execution, are probably causing more trouble than you’d like on the server side. It’s also a good idea to set error_reporting to be as verbose as possible (like E_ALL to report all errors), so you can keep track of what’s going on.

Assuming you support file uploads on your server (as most content management systems do), you’ll want to keep an eye on the upload_max_filesize and post_max_size settings as well. These will help PHP control how much buffer is allocated to read in and handle an incoming request. The post_max_size setting also contains how much data is allowed in standard form submissions. On many servers, these are set to ridiculously high levels—or even infinite—to allow for arbitrary file uploads of unknown size.

The tuple of max_execution_time, max_input_time, and memory_limit control how long you’ll let PHP run before committing seppuku. The defaults that ship with PHP are quite reasonable, but as applications grow in complexity (given a server that doesn’t grow with the application it hosts), developers might feel the need to increase or even disable the limits.

disable_functions is a handy directive that accepts a comma-delimited list of core PHP functions to disable entirely. Many servers do not have anything in this directive at all by default, meaning the whole gamut of PHP’s API is available to the application. Depending on how your application is using system resources, this may or may not be a requirement.

Finally, the PHP application can be locked down to only access certain parts of the filesystem using open_basedir. This is not a setting enabled by default, meaning the PHP process can talk to the entire filesystem, potentially editing or reading files your application has no business even knowing exist.
MySQL

One of the easiest ways to ensure your database server is secure is to move beyond hosting it yourself. Services like Amazon RDS allow for hosting highly-scalable MySQL databases that leverage all of the identity and access management infrastructure built into the rest of Amazon’s cloud. Unfortunately, this isn’t always the best fit for applications and, more often than not, your PHP application will be living alongside MySQL on the server. If that’s the case, there are some settings to keep in mind when trying to maintain proper security and access control in your own system.

    If at all possible, your database (MySQL) and application (PHP) should run on separate servers. This helps to enforce strict security between them, and also makes it easier to scale up (for better performance) when your application is under load.

The following four settings, two of which occur in the server’s my.cnf configuration and two of which are in the database itself, will help keep your data locked-down when it ships to production.

The first my.cnf directive to watch is bind-address, which determines the IP address that MySQL can listen on. When connections are made to the MySQL server, an internal firewall will determine whether or not the connection is allowed to proceed to an authentication step. It’s common to set this directive to a 0.0.0.0 wildcard, which allows access to the database from anywhere. Unfortunately, this means anyone can prod the database and, if passwords are insecure (or leaked), breach into your datastore.

    Note: For the uninitiated, an IP address like 0.0.0.0 (or 0:0:0:0:0:0:0:0 or [::] for IPv6) will match any IP address. When building services locally, these wildcards allow services to attach to either your local 127.0.0.1 address or your machine’s physical IP. They effectively allow a service to listen on a port without filtering the IP address being requested at all. Wildcards are hugely powerful, but can easily lead to security issues if not used intentionally. Wherever possible, they should be avoided.

MySQL comes with a helpful feature which allows developers to load data into a table based on entries in local files on the same system. If this feature is left enabled, an attacker can trick the database into loading data from other locations on the server, potentially adding themselves to any user stores or even escalating the privileges of existing users. The Boolean local-infile option is usually disabled by default, but some developers will enable it in development to aid in bootstrapping default data and relationships.

Every MySQL user is associated with a host—the IP address from which they’re authorized. While these settings do not exist in my.cnf, they can be viewed by enumerating users directly when authenticated with a root-level account.

SELECT user, host FROM mysql.user;

Most users will have a host set to either localhost or 127.0.0.1. While these are equivalent to web implementations, MySQL treats both as different hosts. This behavior often leads developers to add users with a wildcard host value of % to allow authentication from anywhere. It’s a neat shortcut but means there is no security within MySQL itself regarding the source of a user’s connection.

Finally, it’s common practice to use a single MySQL server to host multiple databases—particularly in shared hosting environments where the number of allotted databases is limited on a per-account basis. It’s also relatively simple to reuse MySQL user credentials to create, manage, and interact with these various databases. An engineer might never create anything beyond a root user in development. This means if one application using your database server is breached, all applications using your database server are breached.
How Would This Code Look If Patched?
Web servers (NGINX and Apache)

Disabling the server tokens returned by NGINX and Apache doesn’t necessarily increase the security of either. However, if you are running an unpatched version in production for any reason, it will make it harder to identify and target the version being used. Both servers use similar directives to control the output of their data. On NGINX, setting server_tokens off will disable the header entirely. On Apache, settings ServerTokens Prod will reduce the output to just reporting that Apache is being used.

Apache does not give you the option to disable disclosing the server in use entirely.

Apache also will print its name and version in the footers of any error page it prints unless you set the ServerSignature directive to Off as well. It’s an added step but helps to lock things down so you aren’t leaking configuration data.

One of the injection examples used in discussing ASR1 was how older versions of PHPMailer failed to properly sanitize the value of the sender’s email address before passing it on to Sendmail. Because of the way Sendmail works, this value is passed as a command line flag; allowing unsanitized input permits an attacker to execute arbitrary system code through those broken versions of PHPMailer. What’s more, the default sender address used by many systems is derived from the superglobal $_SERVER['HTTP_HOST'] value.

If the server name is not configured in NGINX or Apache, they’ll use the value of the HOST header sent by the user when making the initial request. If your server is running in multitenant mode, a malicious HOST header will only affect the default site; other virtual hosts will likely have their name explicitly set already. Regardless of whether your system is serving multiple sites or a single application, you should never use the default virtual host. On NGINX, always set the server_name directive. On Apache, always set the ServerName directive. These static values will then be used to populate $_SERVER and other constants or functions that depend on knowing the name of the host.

Directory traversal can potentially expose a wealth of information to someone inspecting your site. They might see static files intended for other users. They might determine the nature of the software running on the system. They might even be able to see otherwise sensitive configuration files erroneously placed in the web root. Your server should not allow end users to navigate aimlessly through its files.

NGINX can serve a 403 Forbidden status in response to requests for empty directories if autoindex off is set in the location directive for those directories. Setting it on the root path will trigger a 403 error whenever a default index file (specified with the index directive) is not found in the directory. Apache can disable directory browsing by adding Options None or Options -Indexes to the <Directory> block controlling the location.

Thanks to the Internet Security Research Group’s Let’s Encrypt project, SSL certificates are no longer a burden to acquire and configure. Certificates from Let’s Encrypt are entirely free for any domain. The project also published an automated configuration tool which integrates with both NGINX and Apache to install newly-provisioned certificates into the correct locations. Automated tests like Qualys’ SSL Labs can dynamically test your server configuration to identify out-of-date or weak cryptographic primitives in use and help guide you towards a properly secure system.

The security team behind the website for a major political candidate came under fire last year for skipping out on error handling and instead, permitting any requested URL on their domain to resolve on their campaign website. This lead to several hilarious Twitter screenshots showing inappropriate URLs that resolved to real web pages. In a few cases, the URL was also used to populate the text in the header of the page! This behavior kept the site from serving up an error, but the resulting media coverage of inflammatory, antagonistic, or otherwise immature headlines being printed at the top of a real website was a major embarrassment to the team.

If a document in your web application is not found, it should always return a 404 Not Found error. If a page requires authentication, it should always return either a 401 Unauthorized or a 403 Forbidden error. If a page is an interface to a teapot, brewing coffee should always return a 418 error.

    Note: Whether your application should serve a 403 or a 404 in response to an invalid request is up to you. The advice in this section in generic, but your actual needs will better guide the behavior of the application. If it’s a blog or newsletter, serving a 404 in response to an invalid but otherwise reasonable URL is acceptable. If the application is, instead, a REST interface laid atop actual resources and paths, a 403 will be a safer way to prevent directory traversal and accidental discovery of protected resources. Think back to the introductory section on threat modeling and consider which approach is best for your specific implementation.

In many cases, these error messages are vital for clients to infer proper behavior. A 404 Not Found being treated as a 200 OK actually breaks certain indexing utilities and package managers. Allowing your application to run smoothly along instead of stopping when presented with garbage inputs also exposes you to potential exploit paths. In the case above, an adversary could inject content into your website or make it look like your web application is doing things it’s not meant to.
PHP

Though every application and server requires a different configuration, many of these sensitive directives can be disabled by default to create a more secure installation. Setting expose_php to “off” will prevent the application from printing the PHP (or operating system) version in headers. Setting both allow_url_fopen and allow_url_include to “off” will lock your server down to working with just local files.

display_errors is a finicky setting. Regularly on production, this should always be set to “off” to prevent PHP from dumping potentially sensitive information to the browser when things go awry. However, it can be just as easy to set this to “off” in development and ignore warnings, notices, and errors while you write your code. I would strongly urge you to disable error reporting in production, but would even more strenuously encourage you to enable the most verbose error reporting possible in development.

    In production, display_errors should be set to the literal 0 to disable it. If you want to completely silence error reporting, set error_reporting to 0 to disable it. Alternatively, leave error_reporting set to a verbose level, but also set log_errors to 1 and error_log to a specific file (such as /var/log/php-error.log). This will allow you to keep track of any issues reported.

All of the file size and resource limiting directives should be set to reasonable levels for your application. No one can dictate what the best upload_max_filesize or memory_limit will be for your application. The thing to keep in mind is that unbounded uploads and memory make it easier for an attacker to trigger a denial-of-service attack in an application. They could merely attempt to stream /dev/random to your server. Or craft a request to trigger a runaway loop which exhausts system memory. If you can avoid uploading files at all, do so. It’s safer to use an external resource (like Microsoft Azure or Amazon S3) for hosting. If you can profile appropriate memory usage within your application, do so. Then set these directives such that PHP only has access to the resources it truly needs to serve the application.

There are some functions that, unless you are very careful and intentional with their use, are more trouble for your system than they’re worth; functions like exec, fopen, and show_source. The first (and other, similar functions) allows for the execution of arbitrary system-level commands on your server. The second allows for the loading of arbitrary files on the system (perhaps even root-level configuration files). The third allows for inspection of PHP code that’s being used to run your application.

The functions aren’t security risks unto themselves, but using them in an insecure fashion is remarkably easy and can lead to major vulnerabilities within your application. If you’re not using them—and any use should be limited and explicitly intentional—you should use the disable_functions directive to turn them off entirely. The list of functions you should consider disabling include:

    exec
    passthru
    shell_exec
    system
    proc_open
    popen
    parse_ini_file
    show_source
    eval
    create_function

Again, there are legitimate arguments that can be made for each and every one of these functions. Some utilities require them (the PHPMailer library, for example, uses popen() to communicate with Sendmail). Legacy code might use them because certain features were nonexistent in earlier versions of PHP. The point is these functions can be easily misused and should be avoided. If you can safely disable them outright, you should do so to prevent accidental misuse in the future.

Finally, PHP should be locked down such that your application has access to itself and nothing else. You can do this by specifying open_basedir to be the root of your application’s installation. It will be able to access files from its own directory and nothing else on the system. If you need functions like fopen() for your application to behave properly, this one change will help prevent your application from potentially accessing—or even modifying—files outside of its purview.
MySQL

If the MySQL server is colocated with the PHP engine rendering the site, there is no reason it should be listening for connections on any IP address other than locally. Setting bind-address = 127.0.0.1 will block any remote connections and ensure only PHP (or other applications on the server) can interact with your database.

    It is highly preferable that your MySQL server not be colocated with the PHP engine. It’s better in terms of security that the applications run in dedicated environments (and helps to increase the overall performance of each as they won’t share resources). In a multi-server environment, you will need to allow remote connections, but can explicitly whitelist the addresses of those remote connections.

The unfortunate effect of this change is that you can no longer connect to the database remotely, either. Once MySQL is locked down to local connections only, you’ll have limited methods with which to interact with the server. For command line aficionados, connecting to the server over ssh and invoking mysql directly might be enough. Those who prefer a graphical interface to the database can use tools like MySQL Workbench, which also support connecting to a MySQL instance via an SSH tunnel.

Other users might prefer installing a phpMyAdmin instance on the same server and using PHP again to interact with the database. This works well in local or even shared development environments but is inadvisable for a production installation. The entire point of locking the database down to only local connections is to prevent remote access. Adding a server tool enabling remote access negates any protection gained through bind-address.

As said earlier, the default value for local-infile is already 0. That said, it’s a setting sometimes used by developers to tweak a development environment and streamline the bootstrapping of default or starter data for testing. Ensure this tweak never makes it to production and your server will be protected from loading data via local files.

Once bind-address is properly configured, you’ll likely notice some issues connecting to MySQL from PHP unless the hostname in use matches the local machine. You’ll also run into issues if the user with which you’re authenticating is associated with a host other than the local machine.

    Note: Adding a new user associated with localhost will also associate them with the literal IP addresses 127.0.0.1 and ::1. At this stage, MySQL is only listening for local connections, so these two additional associations are both normal and expected.

Locking down the server such that all users are tightly associated with only the local host and not a remote location is a doubly-redundant way to prevent unauthorized remote access to your database. You’ll also want to be sure the root MySQL user is protected with a strong password to prevent abuse. To ensure everything else is safe and secure, you should audit the users table to ensure no users are associated with a wildcard (%) and, if any are, updating them accordingly:

UPDATE mysql.user SET host='localhost' WHERE host='%';

Proper separation of concerns would dictate that every discrete PHP application on your server would use distinct MySQL user credentials to access its own database. Further, the application’s database should be the only data store it is able to access with those credentials. The risk of reusing logins between applications is that a breach in one (even if due to some other application vulnerability) can lead to a breach in others.

When creating users, you can explicitly grant privileges on only one database, preventing them from accessing or manipulating data elsewhere.

CREATE DATABASE IF NOT EXISTS `my_database_name`;
GRANT ALL PRIVILEGES ON `my_database_name`.*
  TO 'thisuser'@'localhost' IDENTIFIED BY 'thatpass';
FLUSH PRIVILEGES;

This new user will be set with a specific password, associated directly with the local machine, and only has privileges on the specified database. In the example above, the user will have all privileges, including the ability to alter, create, or drop columns. If your application does not require this functionality, it would serve you well to create a user with more fine-grained access. Take special care that your database users only have the access they need to act on behalf of the PHP application.
Conclusion

In this chapter, we looked at the different properties and settings required to configure a secure server properly:

    Locking down MySQL to a specific set of users and trusted access locations.
    Preventing PHP from leaking configuration data or allowing insecure access.
    Protecting Apache or NGINX from third-party abuse.

