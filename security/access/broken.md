Broken Access Control

https://www.owasp.org/index.php/Broken_Access_Control

Access control enforces policy such that users cannot act outside of their intended permissions. Failures typically lead to unauthorized information disclosure, modification or destruction of all data, or performing a business function outside of the limits of the user. Common access control vulnerabilities include:

    -Bypassing access control checks by modifying the URL, internal application state, or the HTML page, 
    or simply using a custom API attack tool
    -Allowing the primary key to be changed to another's users record, permitting viewing 
    or editing someone else's account.
    -Elevation of privilege. Acting as a user without being logged in, or acting 
    as an admin when logged in as a user.
    -Metadata manipulation, such as replaying or tampering with a JSON Web Token (JWT) 
    access control token or a cookie or hidden field manipulated to elevate privileges, or abusing JWT invalidation
    -CORS misconfiguration allows unauthorized API access.
    -Force browsing to authenticated pages as an unauthenticated user or to privileged 
    pages as a standard user. Accessing API with missing access controls for POST, PUT and DELETE.
    
    Broken Access Control examples
Example #1: The application uses unverified data

The application uses unverified data in a SQL call that is accessing account information:


        pstmt.setString(1, request.getParameter("acct"));
        ResultSet results = pstmt.executeQuery();
      

An attacker simply modifies the acct parameter in the browser to send whatever account 
number they want. If not properly verified, the attacker can access any user's account.


        http://example.com/app/accountInfo?acct=notmyacct
      


Example #2: An attacker simply force browses to target URLs

Admin rights are required for access to the admin page.


        http://example.com/app/getappInfo
        http://example.com/app/admin_getappInfo
      

If an unauthenticated user can access either page, it’s a flaw. If a non-admin 
can access the admin page, this is a flaw.

Access control is only effective if enforced in trusted server-side code or server-less API,
where the attacker cannot modify the access control check or metadata.

    Deny access to functionality by default.
    Use Access control lists and role-based authentication mechanisms.
    Do not just hide functions.



How To Defend Broken Access Control Vulnerability?

    Explicitly verify the access controls needs for every chunk of application 
    operation and document it. This requires comprising who can legitimately 
    allow performing an action and what resources the user may access through the action

    Drive entire access control decisions from the lower privileged user’s session

    Employ a central application component for verifying access control

    Verify every single request with this central application component 
    in order to decide whether the request from the user is permitted to access the resources

    Employ programmatic techniques to guarantee that there are no case of exceptions

    For more sensitive functionalities like accessing administrative pages, 
    add additional access restriction with IP address to enforce only users 
    from certain network are permitted to access the resources, irrespective of their login status

    Log each event where sensitive operation is performed that will help to 
    detect and investigate if any access control breaches happen

    Ensure prevention from forced browsing by providing access rights only 
    to the users equal with their privileges.

    Always test unprivileged roles or low-level access based on the information 
    under separation of duties. You can capture and replay the privileged request to test the same.

    Deny all by default that will treat everything not explicitly allowed is banned

    Review the server/application from time to time to detect the holes in the access controls
