`index.md` — Landing
Purpose:
* 10-second understanding
* Install snippet
* 15-line example
* Key differentiators

Keep it minimal.

`philosophy.md`
Purpose:
* Why this exists
* Anti lock-in stance
* Offline-first mindset
* Explicit > magic
* Core vs platform

This is emotional + architectural positioning.
`architecture.md`
Purpose:
* Diagrams
* Flow of login
* Interface injection model
* Component boundaries
* Request lifecycle

This is technical authority.

`design-decisions.md`
This is where maturity shows.
Structure it like this: 
Design Decisions
Why No OAuth (Yet)
Explain:
* OAuth is a protocol layer
* Core authentication comes first
* OAuth can be modular
* Avoid early complexity
* Keep core minimal

Position it as intentional minimalism — not missing feature.
Why Interfaces Over Generics
Explain:
* Interfaces define behavior contracts
* Clear extension points
* Better boundary enforcement
* Generics are not a dependency abstraction tool
* Interfaces keep infrastructure swappable

This shows architectural depth.
Why No Built-In Database
Explain:
* Database choice is architectural
* Different apps need different persistence
* Avoid schema coupling
* Avoid migrations inside auth layer
* Storage is a pluggable concern

This reinforces infra-agnostic identity.
Why No Hosted Mode
Explain:
* This is a library, not a platform
* No control plane
* No cloud dependency
* Runs anywhere Go runs
* Fully self-contained

That’s your anti-SaaS clarity.
`getting-started.md`
Goal:
Someone copies this and has working auth in 60 seconds.
Must include:
* Install
* Config
* Minimal example
* Test example

`guide/`
This is operational documentation.
installation.md
Go install, versioning, go.mod details.
quick-start.md
Minimal working setup.
configuration.md
Explain:
* Required config
* Optional config
* Default behavior
middleware.md
Explain:
* net/http
* future framework adapters
* custom middleware usage
testing.md
This is huge for you.
Show:
* In-memory storage
* No-op rate limiter
* Deterministic flow
* CI usage

This is one of your strongest differentiators.
`adapters/`
Document the philosophy of pluggability.
Each page explains:
* Interface definition
* Expected behavior
* Example implementation
* Production considerations

This reinforces your architecture identity.
`api/`
Pure reference.
No philosophy.
No explanation.
Just:
* Function signatures
* Struct fields
* Errors
* Token types

Engineers love clear reference docs.
`examples/`
Each example:
* Full working code
* Minimal explanation
* Copy-paste friendly

Examples sell libraries.