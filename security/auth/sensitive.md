Sensitive Data Exposure

    Many web applications do not adequately protect sensitive data, such as credit cards, tax IDs, and authentication credentials. Attackers may steal or modify such weakly protected data to conduct credit card fraud, identity theft, or other crimes. Sensitive data deserves extra protection such as encryption at rest or in transit, as well as special precautions when exchanged with the browser.

All of the application security risks listed until now are important, but the risk of exposing customers’ sensitive data is likely one of the most impactful and highly visible of the OWASP Top Ten. Every time an enterprise is hacked and, as a result, breaches the privacy of their customers, it makes the news immediately. Rarely is a breach or hack more visible or more talked about in the media or by consumers evaluating a company or brand.

    In the introduction, we talked about a high-profile breach that occurred at Yahoo! before their acquisition by Verizon. The breach exposed the personal information of several hundred thousand users. It also ultimately cost Yahoo!’s shareholders over $350 million by lowering the company’s acquisition value.

The reason this risk is so highly discussed is because of the direct impact it has on customers. Allowing an account breach or a loss of data due to configuration errors impacts business. Leaking credit card information, personally identifying information, or other sensitive data to the public sphere impacts the lives of affected customers.

Sometimes for many years after the breach as well.
What Are Some of the Practical Risks to Sensitive Data?
Insider Threat

The best developers are often the laziest; lazy developers are among the first to devise ways to automate away busywork or other routine maintenance tasks. As a result, development can focus on the more intricate elements of an application that keep developers more engaged.

Unfortunately, this laziness often leads to weakness when it comes to managing credentials or otherwise protecting sensitive information. It’s easier to share credentials than to manage individual logins for separate developers. It’s easier to email a new account password to an engineer than it is to use a proper password manager. It’s easier to bypass security measures put in place to protect customer information than it is to follow standards.

    Sometimes, security is easily and promptly bypassed by insiders who understand the workarounds. Never underestimate the threat posed to your application by those on the inside. Photo borrowed from Schneier on Security.

Figure 6.1
Figure 6.1

Taking the easy way out, however, can lead to severe issues if there’s ever a threat to the application from inside your organization. If every engineer has the same level of access—or uses the same credentials to access data—there is little to no way to detect who was the source of any breach.

Likewise, if all developers are allowed to wear many hats, any one engineer can wreak damage on the system. Imagine a world where a developer writes a piece of code, reviews their own code, then deploys that code to production in isolation. On a small development team, this might be common practice; it’s easy and fast for the developer to self-audit and immediately deploy. However, if that code is a back door, a page defacement, or the introduction of a malware distribution that one developer’s autonomy has crippled your application.

In a similar vein, shared credentials are as hard to audit as they are to deprecate—if you’re aiming for PCI compliance, credential sharing is also explicitly forbidden. If an engineer leaves the project, whether on good terms with the team or otherwise, they potentially take with them privileged access to the interior of the application. These credentials can be easily mismanaged or lost, opening the database and your customer’s data to attack.

If a developer leaves on poor terms with management, they could abuse their continued access to materially damage your organization or compromise the application’s data. Likewise, a disgruntled—or blackmailed—engineer within the team with more access than necessary can steal customer data and sell it on the open market.
External Breach (I.E., Database Backups)

Most PHP applications will use some form of database to store customer information and make it readily available at runtime. In many situations, this is a variant of MySQL. Customer information, product information, billing information, past or pending orders; all of this data lives in MySQL, ready for retrieval and use by PHP as needed.

The MySQL database might live on the same server that runs PHP. If the application is larger and serves a vast customer base, it might be split across multiple servers and leverage multiple database servers as well. Each of these will require authentication to serve data, and this is where the first risk rears its head.

Misconfiguring the database server is easy. However, if you’ve read through the chapter on ASR6: Security Misconfiguration, you already know how to take care of this on your own.

The second, trickier risk is related to how the data is actually stored. Preferably, your database will store information in two places: the disk of the server running the database application and the backup disk storing a long-term image. Both of these data stores are potentially vulnerable to attack.

An attacker breaking into a running database is a significant issue. It’s, thankfully, also a straightforward risk to protect against. An attacker breaking into an idle server and copying the data directory to another location is also a significant issue. Sadly, it’s also one which is harder to detect when it happens.

Likewise, an attacker breaking into the server or machine housing the external backup of the database is also a risk. They can break in, copy the data to a local or network location, and peruse the full database at their own leisure without your knowledge. The idea of a third-party fully exfiltrating your data store for their own purposes is chilling. The fact that they can take the entire database, rather than running individual queries against a specific user or set of users, means they have access to far more data than most attackers breaching by other means.
Unnecessary Data Storage

It can be tempting for an application to store more data than is actually needed. The marketing team wants access to customer email addresses for future use with newsletters. The sales team wants customer phone numbers and mailing addresses for future up-selling activity through affiliates. The security team wants social security numbers for future credit checks. The business management team wants to store credit card information to make add-on purchases easier.

Depending on the application, this information might all be unnecessary.

Storing data unnecessarily exposes an application to additional risk if that data is later leaked or stolen by an unauthorized party. A hacker attacking a blog, for example, would be excited to discover the blog is tracking not just usernames but also email addresses and phone numbers. This information is valuable in trade among other malicious parties on the internet as it’s one more piece of information that can be used for identity theft.

    If your business uses credit card or other banking information, contract with an outside payment system like Stripe. Outside vendors specialize in storing this kind of data securely and measuring up to strict industry requirements like PCI. Do not try to reinvent the wheel and store this information yourself.

A personal encounter with unnecessary data storage was a client I worked with years ago. On one side of things, they were highly attuned to security issues and required all of their vendors, contractors, and staff to interact with the codebase and servers running it exclusively over an individually-authenticated VPN. Every access was logged for future audit so they could adequately protect against rogue engineers damaging the system.

Getting access to the VPN required an application process through the security team. Though I was part of a larger organization, I was required to complete a personal application to get access. This form required my name, email address, phone number, home address, and social security number. None of this information should have been required of me, though it might make sense for a freelancer.

My team pushed back.

It turned out, a team on their side had conflated other freelancer information requirements with the VPN authorization form. Email address and phone number were used to populate an external contact system used by the business development team to contact future potential hires. Home address was used for tax reporting purposes. Social security data was originally used in the same way but had been migrated to a proper W-9 disclosure. It turned out the security team was instead using SSNs as indexes in the authorization database.

In short—none of this information was necessary to grant me access to the VPN. As a result, I provided none of it, and the team followed suit. The moral of the story is requiring more information than necessary to perform a task results in your application housing more information than it actually needs. If this VPN database were ever stolen or otherwise made publicly accessible, the breach would do significant harm to anyone listed in it.

This is not the kind of risk your application developers need hanging over their heads on a day-to-day basis.
Using Insecure Cryptography

PHP is a fantastic language where just about anything is possible, and almost anyone can write code for a new project from scratch. PHP is a great language for learning about programming and a fantastic tool for anyone trying to prototype a new project quickly. It’s easy to write, very forgiving when you make a mistake, and there are loads of examples posted on the internet for any given problem.

Unfortunately, PHP is also easy for inexperienced engineers to write, overly forgiving when you make a mistake, and bad examples abound on the internet.

It’s a good thing developers dealing with sensitive customer information reach for encryption. When appropriately used, encryption is a solid way to protect customer data by making it unreadable and unusable by anyone without the proper level of access. Unfortunately, cryptography is very hard to do properly, and very easy to do poorly.

A good example is the conflation of the terms “encrypt” and “encode.” Base64 is a type of character encoding which converts binary into human-readable strings of text by converting three-byte chunks into two-character representations using a specific, 64-character alphabet. The fact that regular data looks obfuscated after it’s Base64-encoded makes it easy to confuse with encryption, which also obfuscates the characters or bytes being used.

Encoding is not encryption.

Likewise, it’s easy to confuse different primitives in the world of cryptography. All of the algorithms in use are ciphers which can turn a message (or piece of data) in plaintext into a matching ciphertext given a key. Two-way ciphers allow for converting the ciphertext back into plaintext given a key. One-way ciphers, or hashes, are irreversible.

    Note: While it is feasible to use brute force to determine the plaintext used to generate a specific, hashed value, cryptographic hashes are designed to work in one direction only. Some hash families even have a certain amount of resistance built-in to prevent reversal.

Given some of the poor or misleading documentation on the internet, it’s very easy to confuse the two families of algorithms. Given websites like md5decrypt.org exist and claim to decrypt certain hashes, it’s even easier to understand how some developers confuse the topics.

This confusion compounds into further security issues when developers attempt to implement encryption on their own. Reading an article on Wikipedia or taking a class on Coursera are great ways to learn the basics; neither approach builds the foundation necessary to fully implement a solid cryptographic scheme in isolation.

Taylor Hornby, known in the PHP world as Defuse, documents such failures to implement cryptograph on his blog, Crypto Fails. One of the more recent examples demonstrated someone who confused Base64 encoding with encryption and built his own set of methods for turning strings into “code” and back again:



This function would turn a string like this is a secret into the continuous string:

QVlRHZlbopUYxQWShRkTUR1aaVUWuB3UNdlR2NmR
WplUuJkVUxGcPFGbGVkVqp0VUJjUZdVVaNVTtVUP

It looks like it’s been turned into an encrypted message. Unfortunately,merely applying iterative rounds of Base64 encoding while reversing the direction of the string does not actually protect the message. There are multiple problems with this scheme. The most readily apparent is there is no secret key applied; anyone else who knows the algorithm itself can extract the secret message with no problem whatsoever.

Taylor’s site includes several other novel examples, most with clear explanations of the varying issues with each approach. In at least one case, an allegedly secure implementation of AES is called out for:

    Improperly ordering the authentication and encryption steps of the operation. In practice, a message should first be encrypted, then the ciphertext should be signed with a message authentication code (MAC). Performing the operation the other way around (signing the plaintext with a MAC then encrypting the plaintext and MAC together) is generally considered a bad practice. It exposes significant weaknesses towards truly verifying message integrity and should be avoided.
    Failing to use constant time comparisons, meaning the decryption algorithm is vulnerable to timing oracle attacks.

He even takes the time to demonstrate a proof-of-concept attack against the proposed algorithm. It’s an excellent, deep-dive into how an attacker would exploit such weaknesses in a production system and highlights how easy it is for a smart developer to bypass a poorly-designed system.
How Can These Risks Be Effectively Mitigated?
Staff Management

Sometimes, the risk of exposing sensitive data is greatest due to the intrinsic nature of how applications are developed: the human factor. An application can only be as secure as the weakest point of ingress by a potential attacker. Leaked authentication credentials, poorly-managed passwords, mistakenly opened ports, and insufficiently sanitized code inputs; these are all potentially easy ways for attackers to bypass a secure system.

Admitting the attacker might already be inside the system is the first step to realizing the most secure code in the world is still vulnerable.

Thankfully, there are three specific ways a development team can adequately protect an application from an insider threat.
Separation of Concerns

Individual members of the development team should operate in different roles. Said another way, the idea of critical members of the development team wearing different hats at the same time is contrary to the mission of building a critical infrastructure. While one engineer might be capable of writing, reviewing, and deploying code, they should never be permitted to perform more than one of these roles at a time.

The engineer who writes the code for a feature should never be the one signing off on that code as secure. Likewise, the author of the code should never be the one to deploy that code to production. Separating these roles and requiring checkpoints with varying members of the development team helps to audit code as it moves along the development pipeline.

When code lives on GitHub, requiring explicit reviews and approval on pull requests keeps track of who looked at the code when. GitLab and Bitbucket have similar ways to review pull requests. If code lives in a separate version control system, tools like Review Board make it easy to submit, review, and track pull requests in parallel with the system housing code.

Once code is approved and merged, continuous integrations systems like Travis CI or Jenkins can run any unit or integration tests and push changes out to a staging environment. A separate person should review the staging environment for correctness (i.e., Does the feature satisfy the business requirements that necessitated it?) and promote to production as necessary.

Separating the roles and responsibilities for getting code from conception to the server helps build a solid audit trail to later verify that best practices and required steps were followed throughout. It also ensures a rogue engineer is incapable of deploying malicious code to a production environment.
Credential Audits

Keeping track of authentication credentials can be tricky, particularly as development teams grow over time. What was once a shared password used between two privileged users is now a login used by a much larger team. Changing shared credentials when a member of that team leaves is critical to maintaining proper application security. However, it puts the remaining team in a sore spot having to roll credentials, update scripts, and reauthenticate to known services.

If at all possible, every user of every system should have a discrete login for that system. These logins should be tracked centrally, so when an engineer inevitably does leave the team, their login can be immediately disabled to prevent potential abuse.

To that end, utilities like SimpleSAMLphp enable the centralization of user credentials to one resource. Instead of maintaining separate logins for disparate resources (like staging servers, Google Docs, and other services), developers log in once via a central server and have their identity federated to whatever services require it. Disabling a user becomes a matter of disabling one credential rather than several.

Similarly, individual credentials like SSH and GPG keys should be tracked and audited frequently. It’s easy to remove a former employee’s access from tools used on a daily basis and leave them with access to older tools or those used less frequently. This oversight, while it seems minor, might present opportunity for an attacker (not necessarily the employee themselves) to breach another unprotected system.

Engineers should have a specific public key or set of keys assigned and tracked in a single location. Whenever those keys are added to a server, that addition should be tracked as well. When an engineer leaves the organization, their public key should be removed from any location where it had been added. Further, systems engineers need to perform regular audits to verify the user accounts and keys present on any server match those expected to exist.

If it’s at all possible, engineers should use SSH and GPG keys embedded on physical modules like Yubikeys. These modules provide additional physical protection from abuse by preventing the matching private keys from ever being extracted. Engineers leaving the team can surrender their physical key, providing a way to audit they no longer possess access to the servers.

Finally, utilities like HashiCorp’s Vault enable managing API keys, passwords, and even private keys as centralized resources controlled by the organization. Credentials can be issued as needed—even temporarily—and easily audited by members of management to ensure only authorized access is possible.
Principal of Least Privilege

At all times, engineers should only be granted access to perform tasks necessary to complete the task at hand. Giving a developer root access to a server, merely so they can update a single package, is a great way to get the job done quickly while also permitting too deep of access to a server. Systems engineers should manage the server; developers should manage the code; QA engineers should manage promotion to production.

These roles are separate, as are the job descriptions, management, responsibilities, and audit trail. By limiting any one employee’s access to a system, you’ve also limited the potential damage they could cause if their credentials are ever breached or if they ever decide to do harm on their own.

This is akin to the way engineers can lock down processes on a web server to prevent any one application from impacting or damaging the runtime environment of another (see the chapter on ASR6: Security Misconfiguration for more details). Consider at all times, the amount of exposure your data has to the engineering team. For example, someone writing frontend code to change the presentation of the website should not have access to the keys used to decrypt account information in the database.

Granting staff only the privileges they need, for only the time during which they need it, limits the impact any one person can have on the sensitive data within your application.
Encryption at Rest

Information should always be encrypted both in-transit and at rest. Encrypting data as it moves from point A to B helps protect it from a malicious eavesdropper who intends to either abscond with or manipulate the data as it travels. Encrypting data at rest protects it from a malicious actor who also has the capability to steal the data all at once.

In short: encrypt everything at all times. The only time data should ever be unencrypted is when it’s being actively used by the application.

If data is unencrypted at rest, a breach of your application’s backups could leak the entirety of your database and all of your users’ information. If the data is encrypted at rest, the breach will still be troubling, but any data protected via encryption is fundamentally useless to the attacker.

Hosting a database with a cloud provider like Amazon presents your team with the opportunity to transparently encrypt the entire database:

    Amazon RDS encrypted instances use the industry standard AES-256 encryption algorithm to encrypt your data on the server that hosts your Amazon RDS instance. Once your data is encrypted, Amazon RDS handles authentication of access and decryption of your data transparently with a minimal impact on performance. You don’t need to modify your database client applications to use encryption.

    Amazon RDS encrypted instances provide an additional layer of data protection by securing your data from unauthorized access to the underlying storage. You can use Amazon RDS encryption to increase data protection of your applications deployed in the cloud, and to fulfill compliance requirements for data-at-rest encryption.

    –Encrypting Amazon RDS Resources

So far as the application is concerned, the data is being processed in the clear. Queries still operate as usual, and there are no programmatic changes required within the application. The data itself, though, is stored on disk in an encrypted format, protecting it if, somehow, anyone ever breaches the physical machine where it’s written.

Standard MySQL installations (including MariaDB and Percona) natively support table-level encryption which provides many of the same benefits, but with the addition of hosting the data yourself. Again, everything is transparent to the application while it’s running, but a breach of the server (to copy the files as written to disk) or the backup will not leak any plaintext user information.
Mission-Critical Data

The easiest way to minimize the impact of a data breach is to minimize the data being maintained. If your application doesn’t need user contact information, don’t collect user contact information. In many cases, your application can avoid even storing authentication information by leveraging an OAuth or OpenID Connect provider for authentication.

Every piece of personally identifying information requested and stored by your application increases the scale of impact of a data breach. Minimizing requirements to only store the minimum information necessary for the application to function decreases the application’s security footprint and helps protect sensitive customer data.

The best way to keep sensitive user data secure is to not have the user data in the first place.
Cryptographic Best Practices

Never build your own cryptography.

That sentence alone is the single best advice any engineer, security-focused or otherwise, can give in terms of properly implementing encryption in an application. Unless the team employs a mathematician who specializes in cryptography (and maybe even if you do) implementing your own encryption algorithms is fraught with danger. Cryptography is hard to get right and very easy to get wrong.

Instead, use well-established, publicly-audited libraries that implement algorithms and other cryptographic primitives for you. The Sodium cryptography library—also known as Libsodium—is an industry-standard, easy-to-use library that supports encryption, decryption, signing, and password hashing. It makes a few very opinionated choices in how it implements the underlying primitives and has been independently, professionally audited for security and correctness.

For versions of PHP up through the 7.1 branch, a PECL module brings native support for libsodium into userland. Thanks to the work of security-minded contributors to PHP, libsodium will be shipping as a native component of PHP starting with version 7.2.

If the developer called out on Crypto Fails had used libsodium instead of devising his own functionality, the code could have looked something like:


Assuming the secret key is stored in a configuration-level constant, this would use a secure, random nonce and the key to encrypt the plaintext string properly. libsodium would use the Salsa20 stream cipher to encrypt the message, attaching a Poly1305 MAC at the same time to help authenticate the validity of the message.

Decrypting will use the same key (and the same random nonce generated during encryption) to extract the message. Since this model uses authenticated encryption, Sodium will also check the validation of the MAC embedded with the ciphertext—if the MAC is invalid, decryption will automatically fail.

The cryptographic primitives here are sophisticated and would be a challenge to implement from scratch in a new project. Thankfully, they’re available through the Sodium project and allow for:

    Authenticated symmetric encryption
    Authenticated asymmetric encryption with elliptic curves
    Anonymous asymmetric encryption through the use of temporary ephemeral keys
    Message signing
    Password hashing

Using Sodium is a solid way to include community-validated, industry-standard, strong, modern encryption in your application. It’s resistant to the attacks which make home-grown and outdated cryptosystems vulnerable to attack. Using Sodium is the proper way to encrypt and thus protect sensitive user data from potential exposure to third parties.

Likewise, be sure to use standard, well-vetted patterns and practices for hashing and storing authentication passwords. Don’t ever store in plaintext, and ensure any password retrievals use safe string comparison techniques. Password hashing and proper storage is discussed at length in the chapter on ASR2: Broken Authentication and Session Management.
Conclusion

In this chapter, we discussed both what not to do with respect to storing sensitive data and how to properly protect user information. This includes:

    Protecting against insider (employee) abuse.
    Ensuring data stores are backed up, remotely, using secure methods.
    Only storing the data your application actually needs to function (and leveraging secure, external vendors for things like payment processing)
    Using secure, industry-standard cryptographic libraries.

