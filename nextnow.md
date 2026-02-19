1. The Core Identity Flow

These are the non-negotiables. They must be fast, secure, and generic to prevent user enumeration.
    SignUp: Creates a new user. It hashes the password using Argon2id and triggers the "Email Verification" flow.

    SignIn: Validates credentials. On success, it returns a Dual-Token Pair (JWT Access Token + Opaque Refresh Token). It also sends the X-RateLimit headers.

    SignOut: Revokes the specific Refresh Token being used. This immediately stops that specific device from getting a new Access Token.

    SignOut (All Devices): This is your "Panic Button" feature. The engine deletes all refresh tokens associated with that userID in the database.

2. The Account Safety Flow

Since you want devs to "own" their users, you need to provide the hooks for them to manage those users safely.
    ChangePassword: Requires the old password + the new password. Once changed, it should automatically trigger a SignOut (All Devices) for maximum security.

    Email Verification: * Offline Mode: Prints a "Magic Link" or 6-digit code to the terminal.
    Online Mode: Sends a real email via the Notifier interface.

    ForgotPassword / Reset: Generates a short-lived (e.g., 15 mins), one-time-use token. Like the login, it must "fail securely" by not revealing if the email exists.

3. Developer "Quality of Life" Features

This is what makes your library better than just "rolling your own."

    The "Dev-Mode" Switch: A single config flag that swaps all interfaces to their In-Memory or Console versions.

    Standard Middleware: Pre-built code that a dev can drop into any Go web server to protect routes (e.g., OnlyAuthenticated(next)).

    Audit Trail: Every action above (failed login, password change, etc.) is sent to the AuditLogger interface in the background.

Define the 4 Key Interfaces: IdentityStore, AuditLogger, RateLimiter, and Notifier.

Feature, Description, Security Detail
SignUp, Register a new user., Uses Argon2id (memory-hard hashing).
SignIn, Authenticate and get tokens., Constant-time checks to stop timing attacks.
SignOut, Kill the current session., Revokes the specific Refresh Token.
SignOut (Global), Logout from everywhere., Deletes all tokens for that UserID.
Email Verify, Prove ownership of email., Works offline (prints to console) or online.
Change Pass, Update security credentials., Triggers a global logout for safety.
Rate Limiting, Prevent brute-force., Includes X-RateLimit headers for the frontend.