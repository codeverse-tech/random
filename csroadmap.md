Perfect — here’s a **one-page Credensync MVP blueprint** you can use as your **step-by-step reference** to build the project yourself. Everything you need is organized for engine → storage → wrappers → dev/end-user flows.

---

# **Credensync MVP Blueprint**

## **1️⃣ Modules / Responsibilities**

| Module               | Responsibility                                                           | Notes                                                    |
| -------------------- | ------------------------------------------------------------------------ | -------------------------------------------------------- |
| **Engine / Core**    | Signup, login, sessions, password hash, opaque tokens, validation, hooks | Pure Go, storage-agnostic                                |
| **Storage Adapters** | CRUD for users & sessions                                                | Implement `Storage` interface; SQLite first, cloud later |
| **HTTP API**         | Fiber endpoints `/signup`, `/login`, `/logout`, etc.                     | Wrap engine; handle request/response                     |
| **CLI**              | Start dev server, migrate DB, seed users                                 | Optional for dev convenience                             |
| **Utils**            | Token generator, password hash/check, email validation                   | Pure Go, reusable                                        |
| **Config**           | Env vars or config file                                                  | Engine receives already-initialized storage + options    |

---

## **2️⃣ MVP Features**

**Engine / Core**

* Signup (email + password + name)

  * Validations: email format, uniqueness, password strength, name non-empty
* Login (email + password + deviceID) → creates opaque token
* Logout (single session/device)
* Global logout (all sessions)
* Multi-device sessions
* Session expiration & revocation
* Password change, reset, update profile
* Email verification (flag + hook)
* Hooks: `OnUserCreated`, `OnUserLogin`, `OnUserLogout`
* Storage interface: supports SQLite + cloud DB

**HTTP API Wrapper (Fiber)**

* Endpoints: `/signup`, `/login`, `/logout`, `/logout-all`, `/change-password`, `/update-profile`, `/verify-email`
* Return JSON responses from engine errors

**Dev Experience**

* Library + HTTP API usable
* Switch between local SQLite ↔ cloud DB
* See active sessions per user
* Export users anytime

---

## **3️⃣ Core Flows / Interactions**

### **Signup Flow**

1. Dev app → `/signup` (HTTP)
2. Engine: validate email, password, name
3. Engine: hash password → store user in storage
4. Trigger `OnUserCreated` hook
5. Return success to wrapper → HTTP response

### **Login Flow**

1. Dev app → `/login` with email/password/deviceID
2. Engine: validate credentials
3. Engine: create opaque token session → store in DB
4. Trigger `OnUserLogin` hook
5. Return token to wrapper → HTTP response

### **Logout Flow**

* Engine revokes token(s) → DB update

---

## **4️⃣ Error Handling / Validation**

* Email already exists
* Invalid email format
* Weak password (<8 chars, optional: rules)
* User not found / invalid credentials
* Session not found / expired
* Storage errors

> Engine returns typed Go errors; wrapper translates to HTTP/CLI messages.

---

## **5️⃣ Config Philosophy**

* Engine receives **storage instance + token/session options**
* Wrapper decides **env vars** or **config file**
* Example:

```json
{
    "storage": "sqlite",
    "db_path": "dev.db",
    "session_expiry_hours": 24
}
```

---

## **6️⃣ Testing Strategy**

* **Unit tests:** engine functions → signup, login, session validation
* **Mock storage:** test flows without DB
* **Integration tests:** SQLite adapter → engine → simulate signup/login

---

## **7️⃣ Notes / Next Steps**

* After MVP: MFA, OAuth/social login
* SDKs for other languages
* Analytics / audit logs
* CLI enhancements (seeding users, dev server modes)

---

✅ **Survival tip:**

* Stick to your **MVP blueprint**.
* Only implement what’s **in the flows + features**.
* Keep **engine pure Go + storage-agnostic**.
* Track next steps in a small board / doc — you’ll never need to guess what comes next.

---

If you want, I can make a **visual flow diagram** for Credensync next, showing **engine ↔ storage ↔ HTTP ↔ dev app ↔ end-user**, so it’s super clear before you write more code.

Do you want me to do that?
