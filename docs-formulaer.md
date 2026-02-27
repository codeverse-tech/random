┌─────────────────────────────────────────────────────────┐
│                   DOCS STRUCTURE                         │
├─────────────────────────────────────────────────────────┤
│  🏠 HOME                                                 │
│     - One-liner: "Auth that you own, in any language"   │
│     - Copy-paste install: `curl -fsSL https://get.csax.io | bash`
│     - See it in action (30-second demo)                  │
├─────────────────────────────────────────────────────────┤
│  🚀 QUICKSTART (Most important!)                         │
│     - 5-minute tutorial with copy-paste commands         │
│     - Works immediately (no account needed)              │
│     - "Hello World" of auth                              │
├─────────────────────────────────────────────────────────┤
│  📦 INSTALLATION                                          │
│     - Download for your OS (big buttons)                  │
│     - Homebrew, apt, curl one-liners                      │
│     - Docker pull command                                 │
├─────────────────────────────────────────────────────────┤
│  🔧 GUIDES (Common tasks)                                 │
│     - Add auth to your React app                          │
│     - Protect Express routes                              │
│     - Use with Python/Flask                               │
│     - Deploy to production                                │
├─────────────────────────────────────────────────────────┤
│  ⚙️ UNDER THE HOOD                                         │
│     - How tokens work (simple explanation)                │
│     - Data flow diagram                                   │
│     - Why no lock-in?                                     │
│     - Security model                                      │
├─────────────────────────────────────────────────────────┤
│  📚 REFERENCE                                              │
│     - CLI commands                                        │
│     - Configuration options                               │
│     - API endpoints                                       │
├─────────────────────────────────────────────────────────┤
│  💡 PHILOSOPHY                                             │
│     - Why we built this                                   │
│     - Who it's for                                        │
│     - Problems it solves                                  │
│     - Our principles                                      │
└─────────────────────────────────────────────────────────┘


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CrydenSync - Auth infrastructure you control</title>
    <style>
        /* Modern CSS Reset */
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        @import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap');
        @import url('https://fonts.googleapis.com/css2?family=Fira+Code:wght@400;500&display=swap');

        body {
            font-family: 'Inter', -apple-system, sans-serif;
            background: #0B1120;
            color: #E2E8F0;
            line-height: 1.6;
        }

        /* Custom scrollbar */
        ::-webkit-scrollbar {
            width: 8px;
            height: 8px;
        }

        ::-webkit-scrollbar-track {
            background: #1E293B;
        }

        ::-webkit-scrollbar-thumb {
            background: #475569;
            border-radius: 4px;
        }

        ::-webkit-scrollbar-thumb:hover {
            background: #64748B;
        }

        /* Navigation */
        .navbar {
            position: sticky;
            top: 0;
            background: rgba(11, 18, 32, 0.9);
            backdrop-filter: blur(12px);
            border-bottom: 1px solid #1E293B;
            padding: 1rem 2rem;
            z-index: 50;
        }

        .nav-container {
            max-width: 1400px;
            margin: 0 auto;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .nav-left {
            display: flex;
            align-items: center;
            gap: 3rem;
        }

        .logo {
            font-size: 1.5rem;
            font-weight: 700;
            background: linear-gradient(135deg, #818CF8, #C084FC);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            letter-spacing: -0.5px;
        }

        .nav-links {
            display: flex;
            gap: 2rem;
        }

        .nav-links a {
            color: #94A3B8;
            text-decoration: none;
            font-weight: 500;
            font-size: 0.95rem;
            transition: color 0.2s;
        }

        .nav-links a:hover {
            color: #F8FAFC;
        }

        .nav-right {
            display: flex;
            align-items: center;
            gap: 1.5rem;
        }

        .nav-right a {
            color: #94A3B8;
            text-decoration: none;
        }

        .github-btn {
            background: #1E293B;
            padding: 0.5rem 1rem;
            border-radius: 8px;
            display: flex;
            align-items: center;
            gap: 0.5rem;
            color: #F8FAFC;
            border: 1px solid #334155;
        }

        /* Hero Section */
        .hero {
            background: linear-gradient(180deg, #0B1120 0%, #0F172A 100%);
            padding: 4rem 2rem;
            position: relative;
            overflow: hidden;
        }

        .hero::before {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background: radial-gradient(circle at 50% 50%, rgba(129, 140, 248, 0.1) 0%, transparent 50%);
        }

        .hero-container {
            max-width: 1200px;
            margin: 0 auto;
            position: relative;
            z-index: 1;
        }

        .hero-badge {
            display: inline-block;
            background: rgba(129, 140, 248, 0.1);
            border: 1px solid rgba(129, 140, 248, 0.2);
            border-radius: 100px;
            padding: 0.5rem 1rem;
            margin-bottom: 2rem;
            color: #818CF8;
            font-weight: 500;
            font-size: 0.9rem;
        }

        .hero h1 {
            font-size: 4rem;
            font-weight: 800;
            line-height: 1.1;
            margin-bottom: 1.5rem;
            background: linear-gradient(135deg, #F8FAFC, #CBD5E1);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            max-width: 800px;
        }

        .hero-subtitle {
            font-size: 1.25rem;
            color: #94A3B8;
            max-width: 600px;
            margin-bottom: 3rem;
        }

        .hero-cta {
            display: flex;
            gap: 1rem;
            margin-bottom: 4rem;
        }

        .btn-primary {
            background: linear-gradient(135deg, #818CF8, #C084FC);
            color: #0B1120;
            padding: 1rem 2rem;
            border-radius: 12px;
            text-decoration: none;
            font-weight: 600;
            transition: transform 0.2s;
        }

        .btn-primary:hover {
            transform: translateY(-2px);
        }

        .btn-secondary {
            background: #1E293B;
            color: #F8FAFC;
            padding: 1rem 2rem;
            border-radius: 12px;
            text-decoration: none;
            font-weight: 600;
            border: 1px solid #334155;
            transition: transform 0.2s;
        }

        .btn-secondary:hover {
            transform: translateY(-2px);
        }

        .hero-stats {
            display: flex;
            gap: 4rem;
        }

        .stat {
            text-align: center;
        }

        .stat-number {
            font-size: 2rem;
            font-weight: 700;
            color: #F8FAFC;
            display: block;
        }

        .stat-label {
            color: #64748B;
            font-size: 0.9rem;
        }

        /* Install Section */
        .install-section {
            padding: 4rem 2rem;
            background: #0F172A;
            border-top: 1px solid #1E293B;
            border-bottom: 1px solid #1E293B;
        }

        .install-container {
            max-width: 1200px;
            margin: 0 auto;
        }

        .install-grid {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 4rem;
            align-items: start;
        }

        .install-left h2 {
            font-size: 2rem;
            font-weight: 700;
            margin-bottom: 1rem;
            color: #F8FAFC;
        }

        .install-left p {
            color: #94A3B8;
            margin-bottom: 2rem;
            font-size: 1.1rem;
        }

        .install-methods {
            display: flex;
            flex-direction: column;
            gap: 1rem;
        }

        .install-method {
            background: #1E293B;
            border-radius: 12px;
            padding: 1.5rem;
            border: 1px solid #334155;
        }

        .install-method h3 {
            color: #F8FAFC;
            margin-bottom: 1rem;
            font-size: 1.1rem;
        }

        .code-block {
            background: #0B1120;
            border-radius: 8px;
            padding: 1rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
            border: 1px solid #334155;
        }

        .code-block code {
            font-family: 'Fira Code', monospace;
            color: #CBD5E1;
        }

        .copy-btn {
            background: #334155;
            border: none;
            color: #F8FAFC;
            padding: 0.5rem 1rem;
            border-radius: 6px;
            cursor: pointer;
            font-size: 0.9rem;
            transition: background 0.2s;
        }

        .copy-btn:hover {
            background: #475569;
        }

        .os-grid {
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 1rem;
            margin-top: 1rem;
        }

        .os-btn {
            background: #1E293B;
            border: 1px solid #334155;
            border-radius: 8px;
            padding: 1rem;
            text-align: center;
            color: #F8FAFC;
            text-decoration: none;
            transition: all 0.2s;
        }

        .os-btn:hover {
            background: #334155;
            transform: translateY(-2px);
        }

        /* Quickstart Section */
        .quickstart {
            padding: 4rem 2rem;
            background: #0B1120;
        }

        .quickstart-container {
            max-width: 1200px;
            margin: 0 auto;
        }

        .section-header {
            text-align: center;
            margin-bottom: 4rem;
        }

        .section-header h2 {
            font-size: 2.5rem;
            font-weight: 700;
            color: #F8FAFC;
            margin-bottom: 1rem;
        }

        .section-header p {
            color: #94A3B8;
            font-size: 1.1rem;
            max-width: 600px;
            margin: 0 auto;
        }

        .steps-grid {
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 2rem;
            margin-bottom: 4rem;
        }

        .step-card {
            background: #0F172A;
            border: 1px solid #1E293B;
            border-radius: 16px;
            padding: 2rem;
            position: relative;
        }

        .step-number {
            width: 40px;
            height: 40px;
            background: linear-gradient(135deg, #818CF8, #C084FC);
            border-radius: 10px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-weight: 700;
            font-size: 1.25rem;
            margin-bottom: 1.5rem;
            color: #0B1120;
        }

        .step-card h3 {
            color: #F8FAFC;
            margin-bottom: 1rem;
            font-size: 1.25rem;
        }

        .step-card p {
            color: #94A3B8;
            margin-bottom: 1.5rem;
            font-size: 0.95rem;
        }

        .language-tabs {
            margin-top: 2rem;
        }

        .tab-buttons {
            display: flex;
            gap: 0.5rem;
            margin-bottom: 1rem;
            border-bottom: 1px solid #1E293B;
            padding-bottom: 0.5rem;
        }

        .tab-btn {
            background: none;
            border: none;
            color: #64748B;
            padding: 0.5rem 1rem;
            cursor: pointer;
            font-weight: 500;
            transition: color 0.2s;
        }

        .tab-btn.active {
            color: #818CF8;
            border-bottom: 2px solid #818CF8;
        }

        .tab-content {
            background: #1E293B;
            border-radius: 12px;
            padding: 1.5rem;
            border: 1px solid #334155;
        }

        .tab-content pre {
            font-family: 'Fira Code', monospace;
            color: #CBD5E1;
            white-space: pre-wrap;
        }

        /* Features Grid */
        .features {
            padding: 4rem 2rem;
            background: #0F172A;
        }

        .features-grid {
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 2rem;
            margin-top: 3rem;
        }

        .feature-card {
            background: #0B1120;
            border: 1px solid #1E293B;
            border-radius: 16px;
            padding: 2rem;
            transition: transform 0.2s;
        }

        .feature-card:hover {
            transform: translateY(-4px);
        }

        .feature-icon {
            font-size: 2rem;
            margin-bottom: 1rem;
        }

        .feature-card h3 {
            color: #F8FAFC;
            margin-bottom: 0.5rem;
        }

        .feature-card p {
            color: #94A3B8;
            font-size: 0.95rem;
        }

        /* Architecture Section */
        .architecture {
            padding: 4rem 2rem;
            background: #0B1120;
        }

        .diagram-container {
            background: #0F172A;
            border: 1px solid #1E293B;
            border-radius: 16px;
            padding: 3rem;
            margin: 2rem 0;
        }

        .diagram {
            display: flex;
            justify-content: space-around;
            align-items: center;
            flex-wrap: wrap;
            gap: 2rem;
        }

        .diagram-item {
            text-align: center;
        }

        .diagram-box {
            width: 200px;
            height: 120px;
            background: linear-gradient(135deg, #818CF8, #C084FC);
            border-radius: 12px;
            display: flex;
            align-items: center;
            justify-content: center;
            color: #0B1120;
            font-weight: 600;
            margin-bottom: 1rem;
        }

        .diagram-box.secondary {
            background: linear-gradient(135deg, #34D399, #60A5FA);
        }

        .diagram-box.tertiary {
            background: linear-gradient(135deg, #F87171, #FBBF24);
        }

        .diagram-arrow {
            color: #64748B;
            font-size: 2rem;
        }

        .explanation-grid {
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 2rem;
            margin: 3rem 0;
        }

        .explanation-card {
            background: #1E293B;
            border: 1px solid #334155;
            border-radius: 12px;
            padding: 1.5rem;
        }

        .explanation-card h4 {
            color: #F8FAFC;
            margin-bottom: 1rem;
            font-size: 1.1rem;
        }

        .explanation-card p {
            color: #94A3B8;
            margin-bottom: 1rem;
        }

        .explanation-card ul {
            list-style: none;
        }

        .explanation-card li {
            color: #CBD5E1;
            margin-bottom: 0.5rem;
            padding-left: 1.5rem;
            position: relative;
        }

        .explanation-card li::before {
            content: '✓';
            color: #34D399;
            position: absolute;
            left: 0;
        }

        /* Philosophy Section */
        .philosophy {
            padding: 4rem 2rem;
            background: #0F172A;
        }

        .philosophy-grid {
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 2rem;
            margin-bottom: 3rem;
        }

        .philosophy-card {
            background: #0B1120;
            border: 1px solid #1E293B;
            border-radius: 16px;
            padding: 2rem;
        }

        .philosophy-card h3 {
            color: #F8FAFC;
            margin-bottom: 1.5rem;
            font-size: 1.25rem;
        }

        .philosophy-list {
            list-style: none;
        }

        .philosophy-list li {
            margin-bottom: 1rem;
            padding-left: 2rem;
            position: relative;
            color: #94A3B8;
        }

        .philosophy-list.problem li::before {
            content: '✗';
            color: #F87171;
            position: absolute;
            left: 0;
            font-weight: 700;
        }

        .philosophy-list.solution li::before {
            content: '✓';
            color: #34D399;
            position: absolute;
            left: 0;
            font-weight: 700;
        }

        .principles {
            background: linear-gradient(135deg, rgba(129, 140, 248, 0.1), rgba(192, 132, 252, 0.1));
            border: 1px solid #1E293B;
            border-radius: 16px;
            padding: 3rem;
        }

        .principles h3 {
            color: #F8FAFC;
            margin-bottom: 2rem;
            font-size: 1.5rem;
            text-align: center;
        }

        .principles-grid {
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 1.5rem;
        }

        .principle-item {
            text-align: center;
        }

        .principle-icon {
            font-size: 2rem;
            margin-bottom: 0.5rem;
        }

        .principle-item span {
            color: #CBD5E1;
            font-weight: 500;
        }

        /* Footer */
        .footer {
            background: #0B1120;
            border-top: 1px solid #1E293B;
            padding: 4rem 2rem 2rem;
        }

        .footer-container {
            max-width: 1200px;
            margin: 0 auto;
            display: grid;
            grid-template-columns: repeat(4, 1fr);
            gap: 2rem;
        }

        .footer-section h4 {
            color: #F8FAFC;
            margin-bottom: 1rem;
            font-size: 1rem;
        }

        .footer-section a {
            display: block;
            color: #64748B;
            text-decoration: none;
            margin-bottom: 0.5rem;
            transition: color 0.2s;
        }

        .footer-section a:hover {
            color: #F8FAFC;
        }

        .footer-bottom {
            max-width: 1200px;
            margin: 3rem auto 0;
            padding-top: 2rem;
            border-top: 1px solid #1E293B;
            text-align: center;
            color: #475569;
        }

        /* Responsive */
        @media (max-width: 1024px) {
            .hero h1 {
                font-size: 3rem;
            }

            .steps-grid,
            .features-grid,
            .philosophy-grid,
            .principles-grid,
            .footer-container {
                grid-template-columns: repeat(2, 1fr);
            }

            .install-grid {
                grid-template-columns: 1fr;
            }
        }

        @media (max-width: 768px) {
            .hero h1 {
                font-size: 2.5rem;
            }

            .steps-grid,
            .features-grid,
            .philosophy-grid,
            .principles-grid,
            .footer-container {
                grid-template-columns: 1fr;
            }

            .nav-links {
                display: none;
            }
        }
    </style>
</head>
<body>
    <!-- Navigation -->
    <nav class="navbar">
        <div class="nav-container">
            <div class="nav-left">
                <div class="logo">crydensync</div>
                <div class="nav-links">
                    <a href="#docs">Docs</a>
                    <a href="#guides">Guides</a>
                    <a href="#architecture">Architecture</a>
                    <a href="#philosophy">Philosophy</a>
                </div>
            </div>
            <div class="nav-right">
                <a href="https://github.com/crydensync">GitHub</a>
                <a href="https://discord.gg/crydensync">Discord</a>
                <a href="#" class="github-btn">
                    <span>⭐</span> 5.2k
                </a>
            </div>
        </div>
    </nav>

    <!-- Hero Section -->
    <section class="hero">
        <div class="hero-container">
            <div class="hero-badge">
                🚀 v1.0.0 just released
            </div>
            <h1>Auth infrastructure<br>you control</h1>
            <p class="hero-subtitle">
                One auth engine for all languages. Self-hosted. Open source. 
                Your users stay yours.
            </p>
            <div class="hero-cta">
                <a href="#install" class="btn-primary">Get Started →</a>
                <a href="https://github.com/crydensync" class="btn-secondary">GitHub</a>
            </div>
            <div class="hero-stats">
                <div class="stat">
                    <span class="stat-number">100%</span>
                    <span class="stat-label">Open Source</span>
                </div>
                <div class="stat">
                    <span class="stat-number">12+</span>
                    <span class="stat-label">Languages</span>
                </div>
                <div class="stat">
                    <span class="stat-number">0</span>
                    <span class="stat-label">Cloud Lock-in</span>
                </div>
                <div class="stat">
                    <span class="stat-number">∞</span>
                    <span class="stat-label">Offline Dev</span>
                </div>
            </div>
        </div>
    </section>

    <!-- Install Section -->
    <section class="install-section" id="install">
        <div class="install-container">
            <div class="install-grid">
                <div class="install-left">
                    <h2>Install in 10 seconds</h2>
                    <p>One CLI. Works everywhere. No account needed.</p>
                    <div class="install-methods">
                        <div class="install-method">
                            <h3>Quick install (macOS/Linux)</h3>
                            <div class="code-block">
                                <code>curl -fsSL https://get.crydensync.io | bash</code>
                                <button class="copy-btn" onclick="copyToClipboard('curl -fsSL https://get.crydensync.io | bash')">Copy</button>
                            </div>
                        </div>
                        <div class="install-method">
                            <h3>Homebrew</h3>
                            <div class="code-block">
                                <code>brew install crydensync/tap/csax</code>
                                <button class="copy-btn" onclick="copyToClipboard('brew install crydensync/tap/csax')">Copy</button>
                            </div>
                        </div>
                        <div class="install-method">
                            <h3>Docker</h3>
                            <div class="code-block">
                                <code>docker run -p 50051:50051 crydensync/server</code>
                                <button class="copy-btn" onclick="copyToClipboard('docker run -p 50051:50051 crydensync/server')">Copy</button>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="install-right">
                    <h2>Download for your OS</h2>
                    <p>Pre-built binaries for all platforms</p>
                    <div class="os-grid">
                        <a href="#" class="os-btn">🍎 macOS Intel</a>
                        <a href="#" class="os-btn">🍎 macOS M1/M2</a>
                        <a href="#" class="os-btn">🐧 Linux amd64</a>
                        <a href="#" class="os-btn">🐧 Linux arm64</a>
                        <a href="#" class="os-btn">🪟 Windows</a>
                        <a href="#" class="os-btn">📱 Termux</a>
                        <a href="#" class="os-btn">🖥️ FreeBSD</a>
                        <a href="#" class="os-btn">📦 Source</a>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!-- Quickstart Section -->
    <section class="quickstart">
        <div class="quickstart-container">
            <div class="section-header">
                <h2>Get started in 5 minutes</h2>
                <p>Three steps to add auth to any app</p>
            </div>

            <div class="steps-grid">
                <div class="step-card">
                    <div class="step-number">1</div>
                    <h3>Install CLI</h3>
                    <p>Download and install the CrydenSync CLI tool</p>
                    <div class="code-block" style="background: #0B1120;">
                        <code>csax --version</code>
                    </div>
                </div>

                <div class="step-card">
                    <div class="step-number">2</div>
                    <h3>Start auth server</h3>
                    <p>Initialize and run your local auth server</p>
                    <div class="code-block" style="background: #0B1120;">
                        <code>csax init my-app</code>
                    </div>
                    <div class="code-block" style="background: #0B1120; margin-top: 0.5rem;">
                        <code>csax dev start</code>
                    </div>
                </div>

                <div class="step-card">
                    <div class="step-number">3</div>
                    <h3>Use in your app</h3>
                    <p>Connect using the SDK for your language</p>
                    
                    <div class="language-tabs">
                        <div class="tab-buttons">
                            <button class="tab-btn active" onclick="switchTab('python')">Python</button>
                            <button class="tab-btn" onclick="switchTab('js')">JavaScript</button>
                            <button class="tab-btn" onclick="switchTab('go')">Go</button>
                            <button class="tab-btn" onclick="switchTab('php')">PHP</button>
                        </div>
                        
                        <div class="tab-content active" id="python-tab">
                            <pre>from crydensync import CrydenSync

auth = CrydenSync()  # localhost:50051
tokens = auth.login("user@example.com", "pass")
print(f"Welcome {tokens.user.email}")</pre>
                        </div>
                        
                        <div class="tab-content" id="js-tab" style="display: none;">
                            <pre>import { CrydenSync } from 'crydensync';

const auth = new CrydenSync();
const tokens = await auth.login('user@example.com', 'pass');
console.log(`Welcome ${tokens.user.email}`);</pre>
                        </div>
                        
                        <div class="tab-content" id="go-tab" style="display: none;">
                            <pre>import "github.com/crydensync/core"

auth := core.New()
tokens, _ := auth.Login(ctx, "user@example.com", "pass")
fmt.Printf("Welcome %s", tokens.User.Email)</pre>
                        </div>
                        
                        <div class="tab-content" id="php-tab" style="display: none;">
                            <pre>$auth = new CrydenSync();
$tokens = $auth->login('user@example.com', 'pass');
echo "Welcome " . $tokens->user->email;</pre>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!-- Features Grid -->
    <section class="features">
        <div class="quickstart-container">
            <div class="section-header">
                <h2>Why developers choose CrydenSync</h2>
                <p>Built for developers who want control without complexity</p>
            </div>

            <div class="features-grid">
                <div class="feature-card">
                    <div class="feature-icon">🔒</div>
                    <h3>Your data stays yours</h3>
                    <p>Auth server runs on your infrastructure. User data never leaves your control. No cloud, no tracking, no lock-in.</p>
                </div>
                
                <div class="feature-card">
                    <div class="feature-icon">🌍</div>
                    <h3>Polyglot by design</h3>
                    <p>gRPC generates SDKs for Python, JS, Go, PHP, Rust, Java, and more. Same API, any language.</p>
                </div>
                
                <div class="feature-card">
                    <div class="feature-icon">📱</div>
                    <h3>Offline first</h3>
                    <p>Develop on a plane, train, or beach. Full auth flow works without internet. Swap to cloud DB in production.</p>
                </div>
                
                <div class="feature-card">
                    <div class="feature-icon">⚡</div>
                    <h3>Go superpowers</h3>
                    <p>Direct import for Go developers means zero network calls. Just `import "github.com/crydensync/core"` and you're done.</p>
                </div>
                
                <div class="feature-card">
                    <div class="feature-icon">🔄</div>
                    <h3>Pluggable architecture</h3>
                    <p>Swap SQLite for PostgreSQL. Bcrypt for Argon2. Memory for Redis. Every component is an interface.</p>
                </div>
                
                <div class="feature-card">
                    <div class="feature-icon">📊</div>
                    <h3>Built-in observability</h3>
                    <p>Audit logs, rate limiting, metrics. Know what's happening in your auth system.</p>
                </div>
            </div>
        </div>
    </section>

    <!-- Architecture Section -->
    <section class="architecture" id="architecture">
        <div class="quickstart-container">
            <div class="section-header">
                <h2>How it works</h2>
                <p>Simple architecture, powerful possibilities</p>
            </div>

            <div class="diagram-container">
                <div class="diagram">
                    <div class="diagram-item">
                        <div class="diagram-box">Your App</div>
                        <span>Python/JS/Go/...</span>
                    </div>
                    <div class="diagram-arrow">→</div>
                    <div class="diagram-item">
                        <div class="diagram-box secondary">Auth Server</div>
                        <span>gRPC API</span>
                    </div>
                    <div class="diagram-arrow">→</div>
                    <div class="diagram-item">
                        <div class="diagram-box tertiary">Your DB</div>
                        <span>Postgres/SQLite</span>
                    </div>
                </div>
            </div>

            <div class="explanation-grid">
                <div class="explanation-card">
                    <h4>🔐 JWT-based authentication</h4>
                    <p>Stateless by design. Tokens contain user claims, validated by signature. No session storage needed.</p>
                    <ul>
                        <li>Access tokens: 15min expiry</li>
                        <li>Refresh tokens: 7 days, revocable</li>
                        <li>No database lookup for validation</li>
                    </ul>
                </div>
                
                <div class="explanation-card">
                    <h4>⚡ gRPC for polyglot support</h4>
                    <p>Define once, generate clients for 12+ languages automatically.</p>
                    <ul>
                        <li>Strongly typed contracts</li>
                        <li>Bi-directional streaming</li>
                        <li>Automatic client generation</li>
                    </ul>
                </div>
                
                <div class="explanation-card">
                    <h4>📦 Pluggable storage</h4>
                    <p>Every component is an interface. Swap implementations without changing code.</p>
                    <ul>
                        <li>SQLite for development</li>
                        <li>PostgreSQL for production</li>
                        <li>Redis for rate limiting</li>
                    </ul>
                </div>
                
                <div class="explanation-card">
                    <h4>🔍 Audit logging</h4>
                    <p>Every auth event is logged. Know who did what and when.</p>
                    <ul>
                        <li>Login attempts (success/failure)</li>
                        <li>Password changes</li>
                        <li>Token refreshes</li>
                    </ul>
                </div>
            </div>

            <div class="diagram-container">
                <h4 style="color: #F8FAFC; margin-bottom: 1rem;">Login flow</h4>
                <ol style="color: #94A3B8; line-height: 2;">
                    <li>1. App calls <code style="background: #1E293B; padding: 0.2rem 0.5rem; border-radius: 4px;">auth.login(email, password)</code></li>
                    <li>2. Auth server checks rate limits, validates password against your DB</li>
                    <li>3. Auth server creates JWT with user claims, returns to app</li>
                    <li>4. App stores token, sends with future requests</li>
                    <li>5. Auth server validates token signature (no DB lookup!)</li>
                </ol>
                <div style="background: linear-gradient(135deg, rgba(129, 140, 248, 0.1), rgba(192, 132, 252, 0.1)); padding: 1rem; border-radius: 8px; margin-top: 1rem; border: 1px solid #334155;">
                    <p style="color: #CBD5E1; margin: 0;"><strong>🎯 Key insight:</strong> Your database stores users. Your auth server validates tokens. Your data never leaves your infrastructure.</p>
                </div>
            </div>
        </div>
    </section>

    <!-- Philosophy Section -->
    <section class="philosophy" id="philosophy">
        <div class="quickstart-container">
            <div class="section-header">
                <h2>Our philosophy</h2>
                <p>Why we built CrydenSync and who it's for</p>
            </div>

            <div class="philosophy-grid">
                <div class="philosophy-card">
                    <h3>😫 The problem with auth today</h3>
                    <ul class="philosophy-list problem">
                        <li>Every project rewrites the same auth code</li>
                        <li>Cloud auth locks you in (your users become their users)</li>
                        <li>Testing auth offline is impossible</li>
                        <li>Documentation is complex and scattered</li>
                        <li>Switching languages means rebuilding auth</li>
                        <li>Self-hosted solutions are hard to configure</li>
                    </ul>
                </div>
                
                <div class="philosophy-card">
                    <h3>✨ Our solution</h3>
                    <ul class="philosophy-list solution">
                        <li>One auth engine for all languages</li>
                        <li>You own your data - always</li>
                        <li>Works offline, deploys anywhere</li>
                        <li>Simple, friendly docs with examples</li>
                        <li>Same API across all languages</li>
                        <li>Pluggable architecture, sensible defaults</li>
                    </ul>
                </div>
            </div>

            <div class="principles">
                <h3>⚡ Our core principles</h3>
                <div class="principles-grid">
                    <div class="principle-item">
                        <div class="principle-icon">🔓</div>
                        <span>Open source</span>
                    </div>
                    <div class="principle-item">
                        <div class="principle-icon">🏠</div>
                        <span>Self-hosted first</span>
                    </div>
                    <div class="principle-item">
                        <div class="principle-icon">📱</div>
                        <span>Offline by design</span>
                    </div>
                    <div class="principle-item">
                        <div class="principle-icon">🌍</div>
                        <span>Polyglot from day one</span>
                    </div>
                    <div class="principle-item">
                        <div class="principle-icon">🔒</div>
                        <span>Your users are yours</span>
                    </div>
                    <div class="principle-item">
                        <div class="principle-icon">🎯</div>
                        <span>Simple by default</span>
                    </div>
                    <div class="principle-item">
                        <div class="principle-icon">⚡</div>
                        <span>Performance matters</span>
                    </div>
                    <div class="principle-item">
                        <div class="principle-icon">📊</div>
                        <span>Observable by design</span>
                    </div>
                </div>
            </div>

            <div style="margin-top: 3rem; text-align: center;">
                <h4 style="color: #F8FAFC; margin-bottom: 1rem;">👥 Who we built this for</h4>
                <div style="display: grid; grid-template-columns: repeat(4, 1fr); gap: 1rem;">
                    <div style="background: #1E293B; padding: 1rem; border-radius: 8px;">
                        <strong style="color: #F8FAFC;">Indie Devs</strong>
                        <p style="color: #94A3B8; font-size: 0.9rem; margin-top: 0.5rem;">Ship auth in hours, not weeks</p>
                    </div>
                    <div style="background: #1E293B; padding: 1rem; border-radius: 8px;">
                        <strong style="color: #F8FAFC;">Startups</strong>
                        <p style="color: #94A3B8; font-size: 0.9rem; margin-top: 0.5rem;">Move fast without lock-in</p>
                    </div>
                    <div style="background: #1E293B; padding: 1rem; border-radius: 8px;">
                        <strong style="color: #F8FAFC;">Enterprises</strong>
                        <p style="color: #94A3B8; font-size: 0.9rem; margin-top: 0.5rem;">Self-hosted, secure, auditable</p>
                    </div>
                    <div style="background: #1E293B; padding: 1rem; border-radius: 8px;">
                        <strong style="color: #F8FAFC;">Open Source</strong>
                        <p style="color: #94A3B8; font-size: 0.9rem; margin-top: 0.5rem;">Add auth without SaaS dependencies</p>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!-- Footer -->
    <footer class="footer">
        <div class="footer-container">
            <div class="footer-section">
                <h4>Product</h4>
                <a href="#features">Features</a>
                <a href="#install">Installation</a>
                <a href="#docs">Documentation</a>
                <a href="#guides">Guides</a>
                <a href="#pricing">Pricing (free)</a>
            </div>
            
            <div class="footer-section">
                <h4>Developers</h4>
                <a href="https://github.com/crydensync">GitHub</a>
                <a href="/api">API Reference</a>
                <a href="/sdks">SDKs</a>
                <a href="/cli">CLI Tool</a>
                <a href="/examples">Examples</a>
            </div>
            
            <div class="footer-section">
                <h4>Community</h4>
                <a href="https://discord.gg/crydensync">Discord</a>
                <a href="https://twitter.com/crydensync">Twitter</a>
                <a href="/blog">Blog</a>
                <a href="/contributing">Contributing</a>
                <a href="/github">GitHub Discussions</a>
            </div>
            
            <div class="footer-section">
                <h4>Legal</h4>
                <a href="/license">MIT License</a>
                <a href="/privacy">Privacy</a>
                <a href="/security">Security</a>
                <a href="/terms">Terms</a>
            </div>
        </div>
        
        <div class="footer-bottom">
            <p>© 2024 CrydenSync. Open source. Self-hosted. Your users are yours.</p>
            <p style="margin-top: 0.5rem; font-size: 0.9rem;">Built with ❤️ for developers who value freedom</p>
        </div>
    </footer>

    <script>
        function copyToClipboard(text) {
            navigator.clipboard.writeText(text);
            
            // Show a temporary tooltip
            const btn = event.target;
            const originalText = btn.textContent;
            btn.textContent = 'Copied!';
            setTimeout(() => {
                btn.textContent = originalText;
            }, 2000);
        }

        function switchTab(lang) {
            // Update tab buttons
            document.querySelectorAll('.tab-btn').forEach(btn => {
                btn.classList.remove('active');
            });
            event.target.classList.add('active');
            
            // Update tab content
            document.querySelectorAll('.tab-content').forEach(content => {
                content.style.display = 'none';
            });
            document.getElementById(`${lang}-tab`).style.display = 'block';
        }

        // Smooth scroll for anchor links
        document.querySelectorAll('a[href^="#"]').forEach(anchor => {
            anchor.addEventListener('click', function (e) {
                e.preventDefault();
                const target = document.querySelector(this.getAttribute('href'));
                if (target) {
                    target.scrollIntoView({ behavior: 'smooth' });
                }
            });
        });

        // Add copy buttons to all code blocks
        document.querySelectorAll('.code-block').forEach(block => {
            const code = block.querySelector('code');
            if (code && !block.querySelector('.copy-btn')) {
                const btn = document.createElement('button');
                btn.className = 'copy-btn';
                btn.textContent = 'Copy';
                btn.onclick = () => copyToClipboard(code.textContent);
                block.appendChild(btn);
            }
        });
    </script>
</body>
</html>