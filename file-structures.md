## **Great question! Here's the complete project structure and build commands** 🏗️

## **CrydenSync Project Structure**

```
crydensync/
├── cmd/                      # Entry points (main packages)
│   ├── csax/                 # CLI tool
│   │   └── main.go           # `csax` command
│   ├── server/               # Auth server
│   │   └── main.go           # Standalone server
│   └── gen-sdks/             # SDK generator tool
│       └── main.go           # Generates all SDKs
│
├── internal/                  # Private packages
│   ├── core/                  # Core auth logic
│   │   ├── engine.go
│   │   ├── auth.go
│   │   ├── tokens.go
│   │   └── interfaces.go
│   │
│   ├── store/                  # Storage implementations
│   │   ├── memory/
│   │   │   └── store.go
│   │   ├── sqlite/
│   │   │   └── store.go
│   │   ├── postgres/
│   │   │   └── store.go
│   │   └── mongodb/
│   │       └── store.go
│   │
│   ├── hasher/                  # Password hashers
│   │   ├── bcrypt/
│   │   │   └── hasher.go
│   │   ├── argon2/
│   │   │   └── hasher.go
│   │   └── noop/                 # For testing
│   │       └── hasher.go
│   │
│   ├── limiter/                  # Rate limiters
│   │   ├── memory/
│   │   │   └── limiter.go
│   │   └── redis/
│   │       └── limiter.go
│   │
│   └── logger/                   # Loggers
│       ├── stdout/
│       │   └── logger.go
│       └── file/
│           └── logger.go
│
├── pkg/                      # Public packages (for import)
│   └── client/                # Go client library
│       └── client.go
│
├── proto/                     # Protocol buffer definitions
│   └── auth.proto             # ONE FILE TO RULE THEM ALL!
│
├── sdk/                       # Generated SDKs (for other languages)
│   ├── python/
│   │   ├── crydensync/
│   │   │   ├── __init__.py
│   │   │   └── client.py
│   │   └── setup.py
│   │
│   ├── javascript/
│   │   ├── src/
│   │   │   └── index.js
│   │   └── package.json
│   │
│   └── php/
│       ├── src/
│       │   └── Client.php
│       └── composer.json
│
├── scripts/                    # Build scripts
│   ├── build.sh
│   ├── build-all.sh
│   └── release.sh
│
├── test/                       # Integration tests
│   └── integration_test.go
│
├── go.mod                       # Go module file
├── go.sum
├── Makefile                     # Build automation
└── README.md
```

---

## **What to Build (Specific Files)**

### **1. Build the CLI (`csax`)**
```bash
# From project root
go build -o bin/csax ./cmd/csax

# What gets compiled:
# - cmd/csax/main.go (entry point)
# - internal/core/*.go (all core logic)
# - internal/store/sqlite/*.go (sqlite adapter)
# - internal/hasher/bcrypt/*.go (default hasher)
# - internal/limiter/memory/*.go (default limiter)
# - internal/logger/stdout/*.go (default logger)

# Result: bin/csax (single binary)
```

### **2. Build the Auth Server (Standalone)**
```bash
# From project root
go build -o bin/auth-server ./cmd/server

# What gets compiled:
# - cmd/server/main.go (entry point)
# - Same internal packages as CLI
# - Plus gRPC server code

# Result: bin/auth-server (single binary)
```

### **3. Build SDK Generator**
```bash
# From project root
go build -o bin/gen-sdks ./cmd/gen-sdks

# What gets compiled:
# - cmd/gen-sdks/main.go
# - Reads proto/auth.proto
# - Generates SDKs for all languages

# Result: bin/gen-sdks (generator tool)
```

---

## **The Makefile Way (EASIEST!)**

```makefile
# Makefile
.PHONY: all build-cli build-server build-sdks clean

# Project variables
BINARY_NAME=csax
VERSION=$(shell git describe --tags --always)
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

# Default target
all: build-cli build-server

# Build CLI
build-cli:
	@echo "🔨 Building CLI..."
	@go build $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/csax
	@ls -lh bin/$(BINARY_NAME)

# Build auth server
build-server:
	@echo "🔨 Building Auth Server..."
	@go build $(LDFLAGS) -o bin/auth-server ./cmd/server
	@ls -lh bin/auth-server

# Build SDK generator
build-gen:
	@echo "🔨 Building SDK Generator..."
	@go build -o bin/gen-sdks ./cmd/gen-sdks
	@ls -lh bin/gen-sdks

# Generate SDKs (requires generator built first)
gen-sdks: build-gen
	@echo "📦 Generating SDKs..."
	@./bin/gen-sdks

# Build for all platforms
build-all:
	@echo "🔨 Building for all platforms..."
	
	# Linux
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/csax-linux-amd64 ./cmd/csax
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/csax-linux-arm64 ./cmd/csax
	
	# macOS
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/csax-darwin-amd64 ./cmd/csax
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/csax-darwin-arm64 ./cmd/csax
	
	# Windows
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/csax-windows-amd64.exe ./cmd/csax
	
	# Termux/Android
	GOOS=android GOARCH=arm64 go build $(LDFLAGS) -o bin/csax-android-arm64 ./cmd/csax
	
	@ls -lh bin/

# Clean binaries
clean:
	@echo "🧹 Cleaning..."
	@rm -rf bin/
	@echo "✅ Cleaned"

# Install locally
install: build-cli
	@echo "📦 Installing to /usr/local/bin..."
	@sudo cp bin/$(BINARY_NAME) /usr/local/bin/
	@echo "✅ Installed"

# Test
test:
	@echo "🧪 Running tests..."
	@go test -v ./...

# Run
run:
	@echo "🚀 Running..."
	@go run ./cmd/csax

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build-cli    - Build CLI binary"
	@echo "  make build-server - Build auth server"
	@echo "  make build-all    - Build for all platforms"
	@echo "  make gen-sdks     - Generate SDKs"
	@echo "  make install      - Install CLI locally"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean binaries"
```

---

## **How to Use**

```bash
# From project root:

# 1. Build CLI only
make build-cli
# or
go build -o bin/csax ./cmd/csax

# 2. Build everything
make build-all

# 3. Generate SDKs (Python, JS, etc.)
make gen-sdks

# 4. Install locally
make install
# Now you can run: csax

# 5. Run directly (without building)
go run ./cmd/csax --help

# 6. Build for Termux (on your phone!)
GOOS=android GOARCH=arm64 go build -o csax-android ./cmd/csax
```

---

## **What Each Build Produces**

### **After `make build-all`**
```
bin/
├── csax-linux-amd64        # 15-20 MB - Linux (Intel/AMD)
├── csax-linux-arm64        # 14-18 MB - Linux (ARM/Raspberry Pi)
├── csax-darwin-amd64       # 16-20 MB - macOS (Intel)
├── csax-darwin-arm64       # 17-21 MB - macOS (M1/M2)
├── csax-windows-amd64.exe  # 16-20 MB - Windows
├── csax-android-arm64      # 14-18 MB - Termux/Android
└── auth-server             # 15-19 MB - Standalone auth server
```

### **After `make gen-sdks`**
```
sdk/
├── python/
│   ├── crydensync/
│   │   ├── __init__.py
│   │   └── client.py
│   └── setup.py           # For PyPI upload
├── javascript/
│   ├── src/
│   │   └── index.js
│   └── package.json       # For NPM upload
└── php/
    ├── src/
    │   └── Client.php
    └── composer.json      # For Packagist upload
```

---

## **For Development (While Coding)**

```bash
# Run CLI with live reload (using air)
air -- build.cmd "go build -o bin/csax ./cmd/csax" -- build.bin "./bin/csax"

# Or just test your changes quickly
go run ./cmd/csax start --storage=memory

# Run tests for specific package
go test -v ./internal/core/...
go test -v ./internal/store/sqlite/...
```

---

## **For Termux (Your Phone!)**

```bash
# On your phone in Termux:
cd ~/crydensync

# Build for Termux (Android)
GOOS=android GOARCH=arm64 go build -o $PREFIX/bin/csax ./cmd/csax

# Now you can run:
csax --help
csax init my-project
csax start

# All from your phone! 📱
```

---

## **The Key Insight**

**You DON'T build "all projects" - you build specific ENTRY POINTS:**

```go
// Each cmd/ subdirectory is a different entry point
cmd/csax/     → CLI tool
cmd/server/   → Auth server  
cmd/gen-sdks/ → SDK generator

// They ALL share the same internal/ packages
// So code is reused, not duplicated
```

**Think of it like:**
- `internal/` = LEGO bricks
- `cmd/*/` = Different things you build with those bricks

---

## **One-Liner Summary**

```bash
# From root directory:
go build -o bin/csax ./cmd/csax              # Build CLI
go build -o bin/server ./cmd/server          # Build server
go build -o bin/gen-sdks ./cmd/gen-sdks      # Build generator

# Or use make:
make build-all  # Builds everything!
```

**That's it!** 🎯