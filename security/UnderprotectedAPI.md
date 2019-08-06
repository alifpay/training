    Modern applications often involve rich client applications and APIs, such as JavaScript 
    in the browser and mobile apps, that connect to an API of some kind (SOAP/XML, REST/JSON, RPC, GWT, etc.). 
    These APIs are often unprotected and contain numerous vulnerabilities.

Most systems also extend a REST API allowing consumption of data and execution of routines via JSON. 
Custom applications might expose other protocols or interfaces as well.

As developers, it’s easy to place these interfaces in a different mental bucket than the application UI. 
They aren’t intended for end-user use, and security is often geared 
towards either protecting end users or preventing them from abusing the application. 
As a result, best practices like input validation, output escaping, or even data encryption 
often go unimplemented entirely.

The truth is, your application’s remote programming interface is just as visible, accessible, 
and exploitable as its user interface. Focusing on the application risks otherwise exposed 
by your application, but from a programmatic perspective, is a solid first step to keeping things safe.


    Secure traffic between clients and APIs, and between applications and any APIs on third party services. 
    KEMP LoadMaster can assist with this by taking on the burden of SSL/TLS encryption via the SSL Offloading feature.
    
    Make sure that there are robust authentication schemes in place and that they use secured credentials, keys, and tokens.
    
    Perform on access authorization & control and don’t just allow access to all APIs after successful authentication.
    
    Ensure that requests sent to APIs are in a secure format that can’t be hijacked or spoofed.
    
    Make sure there is comprehensive checking for, and protection against, injection attacks as these forms of 
    vulnerability are open via API calls just like they are via web interfaces.
    
    Deploy other network monitoring tools, that can inspect for suspicious activity at multiple layers of the 
    network stack, upstream from your applications. KEMP LoadMaster with Edge Security Pack (ESP) and
    
    Web Application Firewall (WAF) can provide this upstream protection with blocking and alerting on suspicious activity.
    This should be seen as an addition to applications also doing attack detection as each application will 
    know its expected usage patterns.

HMAC-SHA-256 

https://geekflare.com/securing-api-endpoint/

    - HMAC, which incorporates timestamps to limit the validity of the transaction to a defined time period
    - two-factor authentication
    - enabling a short-lived access token by using OAuth
    
    Use quotas and throttling. Place quotas on how often your API can be called and track its use over history. 
    More calls on an API may indicate that it is being abused. It could also be a programming mistake such as calling 
    the API in an endless loop. Make rules for throttling to protect your APIs from spikes and Denial-of-Service attacks.
    
 
