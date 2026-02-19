### 1. Storage (`IdentityStore`)
We decouple your identity logic from data persistence. This allows you to start with zero-config storage and scale to production databases seamlessly.

* **In-Memory (Dev):** Perfect for local testing and prototyping. No database setup required.
* **SQLite (MVP):** A local file-based database. Ideal for small-to-medium apps where you want "offline-first" capability.
* **Postgres/MySQL (Scale):** High-performance adapters for production environments.

> **Why this matters:** You own your data. This library doesn't "lock in" your users; it simply provides the tools to manage them in whatever database you choose.


### 2. Rate Limiting (`Shield`)
Our Rate Limiter acts as a circuit breaker, rejecting malicious traffic before it ever touches your database or expensive business logic.

* **IP-Based:** Prevents bots from brute-forcing sign-ins or spamming account creation.
* **Frontend-Friendly Headers:** Every response includes professional `X-RateLimit` headers:
* `X-RateLimit-Limit`: Your total request capacity.
* `X-RateLimit-Remaining`: How many tries you have left.
* `X-RateLimit-Reset`: Seconds remaining until your limit refreshes.


**UX Tip:** Developers can use these headers to show countdown timers or disable "Login" buttons, making the app feel incredibly polished.


### 3. Audit Logs (`The Black Box`)

Security is only as good as your ability to investigate it. Our Engine records every significant event (Logins, Failed Attempts, Password Changes) asynchronously.

* **Zero Latency:** Logging happens in a background thread (Goroutine), so it never slows down the user's login experience.
* **Standardized Format:** Logs capture the "Five W's":
* **Who:** UserID/Email.
* **What:** Action (e.g., `SIGN_IN_SUCCESS`, `PASSWORD_RESET_REQUEST`).
* **When:** UTC Timestamp.
* **Where:** IP Address and User Agent.
* **Outcome:** Success or Failure reason.


### 4. Messaging & Notifications (`Notifier`)

Handle email verification and password resets without being locked into a single provider.

* **Offline Mode (Terminal):** In development, verification codes are printed directly to your console. This allows you to build and test your entire flow without an internet connection or an API key.
* **Online Mode (Production):** Plug in adapters for SendGrid, Twilio, or SMTP to send real emails to your users.


### 5. Dual-Token System

We implement the industry-standard "Stateless + Stateful" security model:

1. **Access Tokens (JWT):** Short-lived (15 min) signed tokens used for high-speed API authorization without database hits.
2. **Refresh Tokens (Opaque):** Long-lived, random strings stored in your database. These allow you to implement:
* **Sign-Out All Devices:** By deleting all refresh tokens for a UserID, you instantly invalidate every session they have active.
* **Token Rotation:** Every time a refresh token is used, it’s replaced, making it much harder for stolen tokens to be useful.
