Broken Authentication and Session Management

    Application functions related to authentication and session management are often not implemented correctly, allowing attackers to compromise passwords, keys, or session tokens, or to exploit other implementation flaws to assume other users’ identities.

When we talk about authentication, we’re concerned with verifying the identity of the user. Typically, we employ a username and password, with the assumption our user does not share their password with anyone else. At this point, we’re not concerned with what the user is allowed to do, only establishing their identity.

Any application dealing with data faces the challenge of ensuring only the right parties ever have access to the data. Unfortunately, “data” is an abstract concept. It could be user information. Data could be posts in a blog’s database. It could be customer banking details. Fundamentally, dealing with application “data” abstractly makes it easy to gloss over the implications of an unauthorized party having access as equally abstract.

To make things easier to understand, we’ll narrow in on your users’ identities. The wrong party having access to your identity can be a nightmare. In the brick-and-mortar world, this is most easily visualized as a thief having physical possession of your credit card. They can pretend to be you to vendors anywhere and run up debt in your name without easily being caught.

A digital thief impersonating you to a server has similar repercussions. If that server is your bank, they could conduct outgoing transfers from your account. If that server belongs to the Secretary of State, they could change your voter registration or redirect your ballot during election season. If that server is your travel agent’s office, they could reschedule or cancel your upcoming vacation.

The backbone of the internet is built on the assumption a remote party can, beyond the shadow of a doubt, identify you and allow you to conduct business as yourself securely. It’s the responsibility of web applications to ensure they are capable of both identifying and verifying the identity of their users. Failing to do so is a fundamental breach of the trust our users give us when they sign up for our services and applications.

Unfortunately, it is very easy to implement client authentication wrong. We’ll look at four different issues modern application developers face building applications today. Some are failures in underlying tools to protect developers from faulty implementations properly. Others are architectural or conceptual issues arising from mass amounts of incorrect information or poorly constructed tutorials circulating amongst the development community.
Issues Facing Authentication
Session Management

Recall, HTTP is a stateless protocol. That is, one request does not know anything about subsequent ones from the same client. To get around this and build useful applications, cookies and sessions allow you to persist some data between requests for a client.

Client-Side Sessions

One of the first mistakes a developer can make is to shift the responsibility for session management from the server to the client. This is a summarily bad idea.

Sessions can store both anonymous data and also the secure data about your application’s users. This might be shopping cart and purchase information for an ecommerce platform. It might be a user profile for a medical inquiry system, or it could merely be profile information for a blogging platform. The point is, a session stores the state for your user at any one point in time.

It’s a bad idea to allow users to control their own state for an application. If they have control, they can manipulate their own state out of band and submit a potentially invalid state back to the server. Or worse, an attacker could manipulate your user’s data without their knowledge. Since your application has trusted users to manage their own session, there is no way for it to validate the state is correct after it’s been resubmitted—or even if it’s coming from the user you expect!

Said another way—if a user session contains a summary of their account, asking the user to store and manage the session gives them the ability to manipulate it directly without any of the other permissions or requirements the application might otherwise dictate. Sessions—and user state—should always live on the server; the user should have access to their session identifier only and should supply that with every request, so the server can populate its state engine and pick up where things left off in the transaction.
Insecure Session Cookies

Another issue facing sessions, whether they’re built from the ground up or implemented properly with PHP primitives, is the security of the client-side cookie containing the session ID. By default, PHP places no restrictions on this cookie. This means it can be sent over either HTTP or HTTPS by default. It can also be used for standard server requests or even read and manipulated by JavaScript in the browser.


Improper Usage of Primitives

This book’s introduction briefly covered the JavaScript Object Signing and Encryption (JOSE) standard. It’s one of the more popular primitives for creating, signing, and validating the integrity of messages, and is widely used in web applications. The JOSE standard defines ways to represent encryption and signing keys, encrypted data, signed data, and identity tokens. The elements of JOSE underlying identity tokens (also called JSON Web Tokens, or JWTs) are also vital to the OpenID Connect standard.

The presence of JWTs in OpenID Connect for representing identity tokens has encouraged many developers to use them as distributed identity assertions elsewhere as well. JWTs contain information identifying the user and are cryptographically signed by the server that issued them (making out-of-band verification of the signature very straightforward).

The biggest problem with JWTs is, unfortunately, also one of its most significant selling points: flexibility. The development team implementing a JWT-backed system can decide which algorithms and key lengths to use for signing. Because of how the underlying specification is designed, these choices are then encoded into headers within the JWT that are broadcast to remote servers. The idea is the server will then use this information to verify the signature on the JWT independently and either trust the data it contains (e.g., e user’s identity) or reject the request entirely.

This flexibility, also termed “algorithm agility,” is a fatal flaw in the JOSE standard.

In 2015, security researchers found a flaw in many different JOSE libraries which conflated the choice between RSA (public/private keys) and HMAC (pre-shared secret keys) for signing tokens. The specification says that both algorithm choices are perfectly valid and requires specifying which one was used in the header. In this way, an attacker could use a server’s RSA public key to generate an HMAC signature for a token, then provide the token to the server with HMAC identified as the algorithm choice.

    Most of the JWT libraries that I’ve looked at have an API like this:

    # sometimes called "decode"
    verify(string token, string verificationKey)
    # returns payload if valid token, else throws an error

    In systems using HMAC signatures, verificationKey will be the server’s secret signing key (since HMAC uses the same key for signing and verifying):

    verify(clientToken, serverHMACSecretKey)

    In systems using an asymmetric algorithm, verificationKey will be the public key against which the token should be verified:

    verify(clientToken, serverRSAPublicKey)

    Unfortunately, an attacker can abuse this. If a server is expecting a token signed with RSA, but actually receives a token signed with HMAC, it will think the public key is actually an HMAC secret key.

    How is this a disaster? HMAC secret keys are supposed to be kept private, while public keys are, well, public. This means that your typical ski mask-wearing attacker has access to the public key, and can use this to forge a token that the server will accept.

    –Critical Vulnerabilities in JSON Web Token Libraries, Auth0

Application developers are working to build a tool with very specific functionality. More often than not, they aren’t cryptographers or even security experts by trade. As such, they rely on the underlying frameworks used by their applications for authentication to be solid. JOSE doesn’t offer secure defaults, and it’s very easy to leave the permissive, inherently insecure defaults in place.



Password Management

Passwords are tricky business. As end users, we’re often given strict instructions by the various services and applications we use regarding password strength: both lowercase and capital characters, at least one number, at least one non-alphanumeric character, and some length between 8 and 16 characters.

The rules we’re given are endless. And, usually, they’re absolutely meaningless regarding increased security. Their only end effect is encouraging the reuse of easy-to-remember passwords.

As pointed out by the webcomic, xkcd.com, most existing password schemes and requirements are inherently hostile to users, see XKCD: Password Strength
Password Hashing

The problem with rules like those listed above is they highlight how passwords in an application are likely stored in an insecure manner. A maximum password length tells attackers that the passwords are potentially stored in plaintext in a fixed-length database column. Restricting the character set (i.e., prohibiting the use of certain special characters) suggests much of the same. In some cases, web applications will even offer to recover a password in plaintext; the only way this is effective is when the application has access to and could potentially leak the password in the first place!
Constant Time Comparisons

Password lookups are also fraught with issues. Even when passwords are properly hashed, the most common way to verify that a submitted password is accurate is through something like:

VULNERABLE

if (hash($_POST['password'])===$hashed_password) {
  // ...
}

To most developers, the conditional above looks perfectly acceptable. The reality, though, is that it’s a vulnerability. Internally, PHP checks to see if the two strings are equal in length. If they aren’t, it returns false immediately. If they are the same length, then it loops through a comparison between the two strings one character at a time, returning false immediately when it finds a difference or true if all of the characters match.

For the uninitiated, this is known as a “timing oracle attack.” If you can control the value on either side of the identity operator (===), you can gauge how close or how different it is to the other side by testing the length of time it takes to return from a function using the conditional. Given enough time, you can determine the secret/unknown value on the other side of the operator. See ircmaxell’s blog: It’s All About Time for a detailed explanation of how this works, complete with explanations of the C code underlying the PHP engine.
Insecure Database Lookups

A similar vulnerability occurs if you’re looking for secret information (e.g., a password hash or a password reset token) in the database. The database of choice for many PHP developers is MySQL. Like the PHP code above gauging equality, MySQL “exits early” from various queries performed on the underlying data.




Password Management

Passwords should never be stored in plaintext.

Passwords should never be stored with encryption.

Passwords should be stored using one-way hashes.

    Many people use the term encrypted and hashed interchangeably, but each has a specific meaning. Encryption implies the encrypted data can be decrypted in some fashion. Hashed means transforming the data with a function such that you can not recover the original but will always get the same hash given the same input.

Passwords are the first—in many cases, only—line of defense protecting your users’ identities within the application. Therefore, protecting passwords and using them in the safest way possible should warrant particular attention from your development team.
Password Hashing

Every user password in your database should be individually salted and hashed.

In a cryptographically secure application, passwords are hashed with random salts. A secure algorithm makes it impossible to convert from a hash back into a plaintext password. If for any reason your datastore is ever breached or leaked, the hashed versions of passwords you’ve stored are unusable by the attacker and your users’ authentication information is still safe.

Hashing a password also removes the necessity to limit the length of passwords, to require specific character sets (or prohibit others), or really even to restrict users to a minimum password length. Proper hash implementations use a random salt for every password and store the salt along with the hash in the database for later verification.

PHP ships with its own set of password interaction functions. To create a password hash, use the aptly-named password_hash(). This function will generate a random salt on each usage, use the strong bcrypt algorithm, and even allows for increasing the “cost” of generating a hash.

    Note: The cost of a cryptographic hash function is related to the amount of processor time required to generate new hash values. As computers increase in performance and efficiency, older hashing algorithms become easier to brute force and the processors can attempt to “guess” the password faster each time. Modern hashing algorithms summary the ability to increase the cost of use by internally performing several iterations of the hash (usually thousands). Generating a single hash is still incredibly fast on modern hardware, so authenticating is still easy. Attempting to guess a password by creating multiple hashes for comparisons, however, becomes prohibitively expensive as the cost of the operation grows.

When a user creates an account, the application hashes the password as follows:

$hashed = password_hash( 
    $plaintext, PASSWORD_DEFAULT, ['cost' => 12]
);

This would hash the password using the default bcrypt algorithm with a cost of 12. Depending on your system architecture, this may take significant time—note that a cost of 10 is the default. At the time of this writing, the default hashing algorithm in PHP consistently produced hashes 60 characters in length, but it should be noted this could change in the future. Any storage system for these hashes (e.g., a database column) should support up to at least 255 characters to allow for future growth.
Constant Time Comparisons

Comparing the stored hash with a newly-generated password hash is the source of potential timing attacks due to PHP’s underlying use of memcmp for string comparison. To avoid an attacker using a clock to detect flaws in and potentially bypass your application’s authentication scheme, you should always use constant time string comparisons when validating passwords.

Luckily, PHP ships with a function to validate hashes securely and in constant time without requiring you to implement the behavior on your own. Given a plaintext password (submitted through a login form) and a stored password hash (Generated by password_hash() above), your application can merely:

if (password_verify($password, $hash)) {
  // ...
}

The hashes generated by PHP contain information about the algorithm used, the cost factor applied, and even the random salt that had been initially used. PHP’s password_verify uses this to recalculate the hash internally and then compare the derived value with the reference string passed to the invocation. It does so by checking the equality of every byte of each string to avoid the “early exit” flaw of a === comparison that exposes a timing attack.

    Note: Developers supporting projects running in environments older than PHP 5.5 can’t use the password_* functions natively. However, Anthony Ferrara supports a compatibility library that extends support for these libraries to use the newer functionality where needed.

Insecure Database Lookups

Constant time comparisons are tricky with MySQL. The advantage of an external database is it can query and return data quickly; tying up the connection so it can perform constant time evaluations turns the database into a bottleneck for your application. This is often a bad idea.

Instead, your application can partition the information it needs into two components—a lookup and the string to be eventually compared. In our query example before, this would mean retrieving both the user’s ID and the reset token instead of using the reset token itself to perform the query.

$q = 'SELECT user_id, reset_token 
      FROM users 
      WHERE user_id = %d AND NOW() < expires';

If the data is returned from MySQL, the stored reset token can be compared to the user-provided reset token using PHP’s constant time string comparison function, hash_equals():

if (hash_equals($reset_token, $submitted_token)) {
  // ...
}

This change moves the comparison logic out of MySQL into PHP where it belongs—the database should be a data store, not the business logic of your application.


