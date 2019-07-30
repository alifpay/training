 Using Components With Known Vulnerabilities - Why You Can't Ignore It
 
 It’s estimated that well over 80% of all software includes, at least, some open source components. 
 As a result, because of their widespread use, third-party components make a tempting target for potential hackers. 

When the above factors are considered, it's easy to understand why OWASP 
updated their Top Ten list to include “a9: Using Components with Known 
Vulnerabilities.” While OWASP acknowledges that the simplest way to avoid 
known security risks is to avoid using components that were not written in-house, 
they further emphasize that this is an unrealistic option. Such a course of action 
would deprive a company of invaluable resources and significantly increase the cost 
and timeframe of any development projects.

In view of this, what steps should a company take to address A9 concerns, avoiding 
known vulnerabilities and continuing to reap the benefits of open source components?

https://www.cvedetails.com/vulnerability-list/vendor_id-14185/Golang.html

Components, such as libraries, frameworks, and other software modules, almost 
always run with full privileges. If a vulnerable component is exploited, such 
an attack can facilitate serious data loss or server takeover. Applications 
using components with known vulnerabilities 
may undermine application defenses and enable a range of possible attacks and impacts.


A recent example of a vulnerable component is PHPMailer. PHPMailer is a utility 
used by many modern PHP applications to handle dispatching email. It’s also used 
by the larger CMS applications: WordPress, Drupal, and Joomla. Until recently, PHPMailer 
had incorrectly implemented a sanitization function meant to protect applications from unfiltered input.

OWASP dependency-check is a software composition analysis 
utility that detects publicly disclosed vulnerabilities in application dependencies.
https://github.com/jeremylong/DependencyCheck

The Six Dumbest Ideas in Computer Security
http://www.ranum.com/security/computer_security/editorials/dumb/
