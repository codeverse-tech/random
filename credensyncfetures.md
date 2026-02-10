Perfect — let’s consolidate everything we’ve discussed and make a **clear MVP feature list** for Credensync. I’ll split it by **core engine features**, **wrappers / API features**, and **dev/end-user experience**.

---

## **1️⃣ Core Engine (Go library) — MVP**

These are the **must-have auth/session features** in the engine:

* **User management**:

  * Create user (email + password)
  * Delete user
  * Update user profile (basic info: name, email, etc.)
  * Change password
  * Forgot password / reset password
  * Email verification

* **Session & token management**:

  * Login → create session with **opaque token**
  * Logout (single session/device)
  * Global logout (all sessions)
  * Device-specific logout
  * Session expiration & revocation
  * Multi-device sessions (1 token per device)

* **Security & validation**:

  * Password hashing (bcrypt / argon2)
  * Token validation (opaque tokens)
  * Optional hooks for devs:

    * `OnUserCreated`
    * `OnUserLogin`
    * `OnUserLogout`

* **Storage abstraction**:

  * Interface to plug SQLite (dev/local) or cloud DB (Postgres/Supabase)
  * Engine does **not depend on DB type**

---

## **2️⃣ Wrappers / API (MVP)**

* **HTTP API** (Fiber router) exposing endpoints for:

  * `/signup`
  * `/login`
  * `/logout`
  * `/logout-all`
  * `/change-password`
  * `/forgot-password`
  * `/update-profile`
  * `/verify-email`

* **CLI**:

  * `credensync start` → start local dev server
  * `credensync migrate` → create DB tables
  * Optional: `credensync seed-users`

* **Configuration support**:

  * Environment variables or config file to switch between:

    * Local SQLite
    * Cloud DB

---

## **3️⃣ Dev experience (MVP)**

* Devs can:

  * Use **library directly** (self-host)
  * Or use **HTTP API wrapper**
  * Easily switch between **local dev** and **cloud**
  * Access hooks (`OnUserCreated`, `OnUserLogin`)
  * See **active sessions per user**
  * Export users anytime (no vendor lock-in)

* End users can:

  * Sign up, login, logout
  * Reset / change password
  * Update profile
  * Be logged in on multiple devices (sessions tracked per device)
  * Trigger global logout via dev app

---

### ✅ Summary table

| Layer       | Features                                                                                                                          |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------- |
| Core Engine | User CRUD, login/logout, password management, sessions, token validation, multi-device, global logout, hooks, storage abstraction |
| HTTP API    | Endpoints for signup/login/logout/logout-all/change-password/forgot-password/update-profile/verify-email                          |
| CLI         | Start server, migrate DB, seed users                                                                                              |
| Dev UX      | Switch storage via config/env, access hooks, export users                                                                         |
| End User UX | Standard auth flows, multi-device sessions, global logout via dev app                                                             |

---

This **covers everything we talked about as MVP**, leaving advanced features like MFA, OAuth/social login, SDKs, or advanced audit/logging for **post-MVP / 1–2 year roadmap**.

---

If you want, I can **also make a visual MVP diagram showing core engine + wrappers + dev app + end users with all these features**, so you have it ready for docs or planning.

Do you want me to do that?
