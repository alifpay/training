Insufficient Logging and Monitoring

Exploitation of insufficient logging and monitoring is the bedrock of nearly every major incident. 
Attackers rely on the lack of monitoring and timely response to achieve their goals without being detected.

Unlike the previous nine security risks, the issue presented here is one of missed opportunity. In 2016, 
an average data breach was not identified for 191 days. Meaning an attacker had an average of 
six months to attack and attempt to penetrate a system before their infiltration was identified 
and blocked. Correctly logging application and system events—and routinely auditing these logs—is 
critical to detecting a breach early and preventing any subsequent damage it might cause.

Often, logging and monitoring can detect an attacker before they’ve had the chance to actually 
infiltrate the system. Most attackers will start scanning a system using automated 
attacks and scripts hoping to detect an out-of-date component, security misconfiguration, 
or other weakness in the application. Catching an attack in progress before the attacker 
has successfully breached your application’s security gives your team the time it needs to 
put proper protections in place to block any further attacks.

Why Logging Matters

One of my earliest clients qualifies as a success and a failure story when it comes to logging. 
They were a software company focused primarily on servicing a legacy desktop application. As a result, 
their systems administrator had limited experience dealing with and securing large networks. Or so that’s 
what he told me.

I helped launch a new, publicly accessible API to serve application updates, license registration, 
and integrations with their help and FAQ systems. Under test it was lightning fast. In production, 
it slowed to a crawl and failed to provide customers with the expected application speedup. Luckily, 
the systems team logged everything, so I had a place to start debugging.

The first thing I noticed was a slew of logs reporting authentication failures to the company’s LDAP endpoint. 
It seemed one IP address had been attempting to brute force the password for an admin account. 
Several thousand times per second. Every day. For six months. This large influx of traffic was 
clogging the service and preventing otherwise legitimate connections from being created.

I asked the systems administrator about these logs and the attempt at unauthorized access. “Oh, yeah, 
I noticed that a few weeks ago. Don’t worry about it. This is why we make everyone change their 
passwords every few months.”

The company leveraged a large team of remote salespeople who all authenticated against 
the same LDAP endpoint. It was public because they never logged in from a single, 
predictable network and needed consistent access from anywhere. That the system would permit 
so many failed login attempts from the same IP address, however, was a massive security failure. 
One that, thankfully, was quickly corrected.

While the systems team did have a significant security oversight (not blocking an attacker 
after detecting malicious activity), they did the right thing by implementing a comprehensive 
logging system. These logs were instrumental in not just blocking that one attacker from 
gaining unauthorized access; they also helped a few months later to detect and circumvent a 
large DDoS (Distributed Denial of Service) attack against the same platform!

What Events Should We Log?

Knowing we should log is useful; understanding what we should log, even more so. An application 
might trigger thousands of data events during routine operation. Some of these events are more 
useful from a security perspective than others.

For all of the following classes of events, it’s important to track:

    What happened (what was the nature of the event)
    When it happened.
    Where it happened (in terms of code and the IP of the application server).
    To who it happened (in terms of an authenticated user and a request IP).
    What input triggered the event (application state, user-defined input, etc.).



Input Validation Errors

Any application that accepts user input in any way should differentiate between valid and invalid input. 
If an application expects a string but a user provides an integer, log the discrepancy. If an 
application expects an XML document with certain entities and properties, but the user provides 
one with different entities, log the discrepancy.

In web applications leveraging a strong content security policy (CSP), the scripts 
embedded on the page can also be classified as “input” to the application. It’s possible 
to define a reporting URI so browsers can automatically report any violations of your CSP 
(i.e., when an extension or script attempts to inject another script that isn’t allowed by our policy). 
Anyone implementing a strong CSP should also implement a reporting endpoint so they can log and track 
potential attempts to bypass the security setting.


Output Validation Errors

Some applications might generate output based on stored data, either user input cached in a database or 
generated content stored in the same. If any of this stored content, or any programmatically-generated data, 
causes unexpected output in the application, that event should be logged and tracked. Often, this will be the 
result of a warning or unhandled error in PHP. If these errors aren’t caught at the application level, 
the report on the frontend might leak sensitive information to your end users. Logging unexpected output 
can help track where these errors are occurring so the development team can patch the application to 
handle things appropriately.

Authentication Events

Every time someone successfully logs in, your application should track the event. Likewise, 
track the event when they explicitly log out or when their session times out and implicitly 
logs them out. This provides your team with a solid audit trail for determining who was using the application when.

Likewise, every time a user attempts to authenticate and fails, log that attempt. If an attacker 
is attempting to break into a site, your first indication might be a series of authentication 
failures in the logs. Also, repeated failures that are properly logged (with the request IP address) 
can be fed into systems like Fail2ban to block potential abusers proactively.



Authorization (Access Control) Failures

Sometimes, the attack comes from a user with otherwise 
legitimate access to the application. Privilege escalation attacks are when an authenticated
user attempts to breach the sandbox in which the sandbox lets them play. Imagine a blog where
a minimally-privileged “author” account is somehow able to gain control of administrative 
access to install third-party code. This could be tragic.

Not only should developers keep an eye on access control, they should keep track 
of every attempt a user makes to escalate their own access to a system. Did an authenticated 
user attempt to access a web page they’re not permitted to see? Log it. Was an unprivileged 
user cookie submitted with an API request to delete or update a piece of data they aren’t 
supposed to access? Log it.



Application Errors

Many newer developers will entirely disable error reporting in their application to hide notices, 
warnings, and other errors triggered not only by their code but by poorly-architected, third-party libraries. 
This is a mistake.

Every error triggered by your application needs to end up in a log the development team reviews regularly. 
These errors might feel like noise, but can sometimes highlight an attack in progress that’s attempting to 
trick your application to behave in a way you didn’t intend.
Application Startup/Shutdown

One of the most sophisticated hacks I’ve ever encountered in production involved an attacker who 
replaced the NGINX binary on a server with one of their own creation. It still served the website 
as expected, but would also phone home and alert a command-and-control server of every incoming request. 
The attacker used this control to inject his own malicious content into the page whenever he wanted to. 
This led the development team as they scoured thousands of lines of PHP code looking for where he’d 
injected a malware payload.

It took one of the developers scanning system logs to figure out what had happened. He noticed 
NGINX had been stopped and started multiple times and the output in the system log was different 
after the last startup. This led the team to look more closely at NGINX and, after comparing the 
file’s SHA hash with what it was supposed to be, flagged the binary as the source of the issue.

While starting and stopping is a routine event for any application, keeping an eye on when the 
various services running your system reboot will help identify potential problems downstream. 
Is the application rebooting because of a memory issue? Did some code trigger a fatal exception 
that triggered a reload? Has the system applied a security update? Has an undetected attacker 
replaced a key system binary?

Logging when your application starts and stops will help identify all of these events. Logging 
when other applications on your server do the same will help keep track of their behavior as well.
High-risk Operations

Your application will likely have multiple functions and serve different purposes for different users. 
There are, however, some classes of operations which are considered “risky” and therefore should be logged. 
There isn’t anything inherently nefarious going on, but logging these operations will provide the team with 
a solid audit trail should they detect an actual attack elsewhere:

    When new users are added or deleted.
    When users’ permissions or access levels change.
    Any time a user performs an administrative action (i.e. changing settings or updating third-party API credentials)
    Whenever data is either imported to or exported from the application’s store
    Whenever a file is uploaded, either anonymously or by an authenticated user.

Each of these operations is something that will happen with relative frequency during the 
lifetime of your application. At the same time, each is also something an attacker might 
attempt after gaining access to your system. These events could indicate, respectively:

    Creating a hidden user account for future access.
    Elevating the privileges of an otherwise unprivileged account.
    Changing backup settings to store data on an unauthorized system.
    Exfiltrating sensitive customer information.
    Uploading a script to enable remote backdoor access to the application

What Data Should We Log?

Knowing who did what, when is the first step to separating legitimate operations 
from malicious ones. Keeping a log of all such activity, particularly that which might be easily 
abused by an attacker, is necessary to proactively protecting your system.

Logs should contain all of the information necessary for your team to identify:

    A timestamp.
    The agent who performed an action
    The type of action performed.
    The type of the log (i.e., informative or error or warning).
    The application’s name (to differentiate between multiple event sources).
    The server where the application is running (i.e., an IP address).
    The location of the agent performing an action (i.e., an IP address for a remote request).
    Any other information your team deems necessary for evaluating the log entry

Some logs are merely informative: an application started up, or a connection to a remote 
service was established. Other logs are useful for debugging: a stack trace identifying a 
faulty method call. Routinely reviewing logs and identifying those that are useful versus 
those that clutter the logs with noise will help your team keep track of both the application 
and your early warning system.



