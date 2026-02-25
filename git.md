**GIT LEARNING**
    build: Changes affecting build system or dependencies (npm, webpack, Gradle). Example: build: update webpack version.
    chore: Misc tasks that don’t change src or tests (like updating configs). Example: chore: clean up old scripts.
    ci: Continuous Integration config changes. Example: ci: update GitHub Actions workflow.
    docs: Documentation only. Example: docs: add API usage examples.
    style: Formatting, linting, whitespace—no functional code change. Example: style: fix indentation.
    refactor: Code restructuring without changing behavior. Example: refactor: extract utility function.
    perf: Performance improvements. Example: perf: optimize image loading.
    test: Adding or fixing tests. Example: test: add unit tests for auth module.
    feat: A new feature. Example: feat: add login with Google.
    fix: Bug fix. Example: fix: correct null pointer in API.

<type>(<scope>): <short description>

Types:
- feat:   new feature
- fix:    bug fix
- docs:   documentation only
- style:  formatting / no logic change
- refactor: code restructure / no behavior change
- perf:   performance improvement
- test:   adding / fixing tests
- build:  build / dependency changes
- ci:     CI config changes
- chore:  misc / maintenance

Scopes (optional):
- auth, api, db, ui, checkout, payment, docs, build, ci, deps

Examples:
feat(auth): add Google login
fix(api): handle null response
docs(readme): update setup instructions
style(ui): fix button spacing
refactor(auth): extract login helper
perf(cache): optimize query caching
test(auth): add token refresh test
build(package): upgrade webpack
ci(github): add workflow for tests
chore: remove unused scripts

Alright 😎 let’s unlock **Git wizard mode**.
Here are the most useful `git log` combinations you'll actually use in real projects:

# 🧙 1️⃣ Beautiful Graph View (Most Useful)
```bash
git log --oneline --graph --all --decorate
```
What it does:

* `--oneline` → compact commits
* `--graph` → shows branch structure
* `--all` → shows all branches
* `--decorate` → shows branch/tag names

Example output:

```
* 8f3a1c2 (HEAD -> main) Add palindrome fix
* 4c91d20 Initial commit
```

This is 🔥 for understanding branching.

# 🧙 2️⃣ See Only the Last 5 Commits
```bash
git log --oneline -5
```

Super useful when you only care about recent changes.

# 🧙 3️⃣ See What Changed in Each Commit
```bash
git log -p
```
Shows the actual code differences (diffs).
⚠️ This can be long.
Better version:

```bash
git log -p -3
```

(only last 3 commits)

# 🧙 4️⃣ See Who Changed What (Blame Mode)
```bash
git blame palin.py
```

Shows:

* Who wrote each line
* When it was written
* Which commit

Incredibly useful in teams.

# 🧙 5️⃣ Search Commit Messages
Find commits containing a word:

```bash
git log --oneline --grep="fix"
```

Find commits by author:

```bash
git log --oneline --author="Your Name"
```

# 🧙 6️⃣ See Files Changed Per Commit
```bash
git log --stat
```

Shows:

```
palin.py | 10 +++++-----
```

Nice summary view.

# 🧙 7️⃣ Clean Professional View (My Favorite)
Add this as an alias:

```bash
git config --global alias.lg "log --oneline --graph --decorate --all"
```

Then just run:

```bash
git lg
```

And you get the magic view instantly ✨

# 🧠 Bonus: Count Commits Properly
After you make your first commit:

```bash
git rev-list --count HEAD
```
