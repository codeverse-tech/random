⌨️ The csax CLI (CrydenSync Auth eXperience)
The csax CLI is the control center for your authentication engine. It is designed to be interactive-first, meaning you don't need to memorize complex flags to get started.
1. Installation
The fastest way to get csax on your machine:
bash

# via Curl (Modern One-Liner)
curl -sSfL https://get.crydensync.com | sh

# OR via Go (For Gophers)
go install ://github.com

Use code with caution.
2. Core Commands
csax init (The Setup Wizard)
This is the entry point for every new project. It launches an interactive setup to configure your environment.

    What it does: Creates a cryden.yaml file in your root directory.
    Interactive Prompts:
        Choose Storage (SQLite/Postgres/MongoDB)
        Set Master Encryption Key
        Enable/Disable Rate Limiting
        Set Session Expiry (JWT & Opaque tokens)

csax serve (Launch the Engine)
Starts the high-speed Go Auth Server.

    Default Port: 50051 (gRPC) and 8080 (Optional HTTP).
    Live Feedback: Displays a real-time log of connection attempts and health status.
    Example: csax serve --watch (Restarts the engine automatically if cryden.yaml changes).

csax user (Identity Management)
Manage your users directly from the terminal without opening a database tool.

    csax user create: Add a new user (it will prompt for email/pass).
    csax user list: View all registered users in a clean table.
    csax user reset-pass: Manually trigger a password reset for a user.

csax keys (Security Management)
Handle your Cryptographic keys safely.

    csax keys rotate: Generates new RSA/ED25519 keys for signing your JWTs.
    csax keys export: Safely prints your public key for use in your frontend/wrappers.

3. The "Kinder" Configuration (cryden.yaml)
csax prefers a configuration file over long terminal flags. Here is a detailed look at a standard setup:
yaml

version: "1.0"
project_name: "MyAwesomeApp"

# The Brain
engine:
  host: "0.0.0.0"
  grpc_port: 50051
  http_port: 8080

# The Pluggable Organs (Interfaces)
storage:
  driver: "sqlite"      # Switch to "postgres" for production
  path: "./cryden.db"

security:
  hasher: "argon2id"    # Industry standard password hashing
  access_token_ttl: 15m # Short-lived JWT
  refresh_token_ttl: 7d # Long-lived Opaque token

# The Guard
rate_limiter:
  enabled: true
  requests_per_minute: 60

Use code with caution.
4. Why use the CLI?

    Framework Agnostic: Whether you use Django, Express, or Laravel, the CLI runs independently as a "Security Sidecar."
    Developer Experience (DX): Built-in "Debug Mode" allows you to see exactly why a login failed (e.g., "Invalid Signature" or "Rate Limit Exceeded") in beautiful colored output.
    Zero-Config Development: Run csax init and csax serve and you have a production-grade auth system ready for your local JS/Python code to call.

