https://angular.io/guide/security

Angular Security

https://www.codemag.com/article/1805021/Security-in-Angular-Part-1

https://ordina-jworks.github.io/angular/2018/03/30/angular-security-best-practices.html


React js

https://medium.com/dailyjs/exploiting-script-injection-flaws-in-reactjs-883fb1fe36c1

https://medium.com/javascript-security/avoiding-xss-in-react-is-still-hard-d2b5c7ad9412


I just red the following post regarding Localstorage and security.
https://www.rdegges.com/2018/please-stop-using-local-storage/

Basically it warn against to save sensitive data on the Localstorage because 
it may be subject of XSS attack. His explanation makes sense, but loads of tutorials says the contrary.

What is your position regarding saving JWT on the LocalStorage of the browser?

it a good question, let me see if I can help. It all depends on the level of security 
we are looking for in our application. For example, the popular authentication 
provider Auth0 shows most of their examples using local storage, so I wouldn't say that it's a wrong solution.
In the Angular Security Course we also show how to store the JWT in a HTTP Only Secure cookie. 
This type of cookies is not accessible via Javascript, and its only sent by the browser over HTTPS connections.
That is indeed even better in terms of security that storing the token in local storage, 
because the JWT is not accessible by Javascript at all, including from our own code.
If we use a cookie this means that we need to add CSRF defenses to our application.
We also need a separate login page that sets the cookie after authenticating 
with the backend, etc. so the implementation is a bit more advanced but definitively doable.

In summary, I think for most applications receiving the JWT in a URL parameter 
after redirection from login page and storing it in local storage is OK, 
in the sense that it's a sufficiently good solution. If developing a banking site 
I would not use this, and use instead a secure HTTP Only cookie. So its a trade-off 
between security and ease of implementation, as it usual comes down to in this type 
of design decisions, there is not only one right solution. 

<br> dfgofigodifg

