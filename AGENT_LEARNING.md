# AGENT_LEARNING.md — Slither / Go Learning System

## Read This First (Agent Instructions)

This file is the single source of truth for how to teach this developer Go through building Slither, a multiplayer snake game. When this file is dropped into a repo, you — the agent — are responsible for creating and running the learning system described here.

**Do not create all modules upfront.** Create each module only when the developer asks to start it or move to the next one. Modules are generated on demand, one at a time.

---

## Who You're Teaching

- Knows JavaScript and TypeScript well. Has built React apps. Understands async/await, fetch, WebSockets from the client side.
- Has **zero backend experience**. Has never written a compiled language. Doesn't know what a pointer is. Has never thought about concurrency.
- Goal: learn Go properly while building something real and fun — not through toy examples, but through problems that are directly connected to the game.
- Learns by doing, not by reading. Gets bored by theory without immediate application.

**Tone when teaching:** direct, no fluff, no hand-holding but no condescension. Explain the *why* before the *how*. Always map new Go concepts to the JS equivalent they already know. If there's no JS equivalent, say so and explain why Go needs it.

---

## The Project Being Built

**Slither** — a massively multiplayer snake game.

- Shared 4000×4000 world, all players in the same space
- No accounts — device fingerprinting (FingerprintJS) identifies players
- Snake grows by eating particles
- Colliding head-first into another snake's body kills you
- On death: snake becomes static glowing particles **in place** — tracing the exact shape of your body. No explosion, no velocity. The particles just appear where your segments were, glow in place, and wait to be eaten.
- Particles are multicolored (random from a glow palette), not your snake's color
- Particle value: `snake.Points / (segments × particlesPerSegment)` — total value is conserved
- Snake length: `BASE_LENGTH + floor(points / POINTS_PER_SEGMENT)` — derived from points every tick
- Particle wobble/pulse is client-side only, seeded from particle ID so all clients render the same thing

**Stack:** Go backend · React + Vite PWA frontend · WebSockets · Redis · Fly.io

---

## How to Create a Module

When the developer says "start module N" or "next module" or "create module [topic]", do the following:

### 1. Create the directory

```
learn/
└── NN-module-name/
    ├── README.md
    ├── exercise.go
    ├── exercise_test.go
    └── solution/
        └── exercise.go
```

`NN` is zero-padded (01, 02, 03...). Always create all four files at once.

### 2. README.md must contain

- **Concept explanation** — plain English, no jargon without definition
- **JS → Go mapping** — a side-by-side comparison showing what this looks like in JS/TS vs Go. Always include this. It's the fastest way to make new syntax feel familiar.
- **Why this exists in Go** — what problem does this language feature solve. Don't skip this. The developer needs to understand the reasoning, not just the syntax.
- **Where this appears in Slither** — one or two concrete examples of exactly where this concept is used in the game code. Make it tangible.
- **How to run** — exact commands to run tests from this module's directory

### 3. exercise.go must

- Declare `package` matching the directory name (e.g. `package structs`)
- Contain 4–6 functions with `TODO` comments for the developer to implement
- Start simple (first 1–2 functions are straightforward) and build toward something game-relevant (last 1–2 functions should feel like real Slither logic)
- Use `// TODO:` comments that explain *what* to write, not *how* — give enough direction that they're not lost, not so much that you're doing it for them
- Never have the correct implementation already written — stubs only (`return 0`, `return ""`, `return nil`, empty body)
- Include a comment at the top explaining what the file is and how to run tests

### 4. exercise_test.go must

- Test every function in exercise.go
- Use table-driven tests where there are multiple input/output cases (Go's idiomatic test pattern)
- Have descriptive test names that read like sentences: `TestGreet_ReturnsFormattedWelcome`
- Cover edge cases (zero values, empty slices, boundary conditions)
- Print a helpful failure message: `t.Errorf("got %v, want %v", got, want)`
- All tests must FAIL when exercise.go stubs are returned (return 0, "", nil) — if a stub accidentally passes, the test is wrong

### 5. solution/exercise.go must

- Contain the complete, correct, idiomatic implementation
- Include comments explaining non-obvious decisions
- Be the reference the developer checks *after* their tests pass, to compare their approach to idiomatic Go

---

## Module Curriculum

These are the 10 modules. Each builds on the last. Do not skip or reorder them.

### Module 01 — Go Basics
**Concepts:** variables, `:=` shorthand, types, functions, `for` loops, `range`, `fmt.Sprintf`  
**Slither connection:** the foundation of everything — variable declarations and loops appear in every single file  
**Key JS mapping:** `const x = 5` → `x := 5`, `for...of` → `for _, v := range`, template literals → `fmt.Sprintf`  
**Exercise arc:** greet a player → clamp a speed value → sum particle values → count snakes above a score threshold → format a leaderboard entry

### Module 02 — Structs and Methods
**Concepts:** `struct`, attaching methods with receivers, value vs pointer receivers, constructors by convention (`NewXxx`)  
**Slither connection:** `Snake`, `Particle`, `Vec2`, `World` are all structs. `snake.Move()`, `snake.Length()`, `particle.IsExpired()` are methods.  
**Key JS mapping:** `class` with properties → `struct`, `class methods` → methods with receivers, `this` → the receiver variable  
**Exercise arc:** define a `Vec2` struct → add an `Add()` method → define a `Snake` struct → write `Snake.Length()` → write `Snake.IsAlive()` → write a `NewSnake()` constructor

### Module 03 — Slices and Maps
**Concepts:** slices (append, len, cap, slicing with `:`), maps (make, set, get, delete, zero value on missing key), iterating both  
**Slither connection:** snake body is `[]Vec2`. World stores snakes as `map[string]*Snake` and particles as `map[string]*Particle`. Every tick iterates both.  
**Key JS mapping:** `[]` → slice, `{}` used as dictionary → map, `arr.push()` → `append()`, `obj[key]` → `m[key]`  
**Exercise arc:** build a snake body from scratch using append → prepend a new head and trim the tail → store snakes in a map by deviceID → look up a snake safely → delete a snake → count total particle value across a map

### Module 04 — Pointers
**Concepts:** what a pointer is, `*T` vs `T`, `&` operator, when to use pointer receivers, nil pointers, why Go needs explicit pointers when JS doesn't  
**Slither connection:** `World` is always passed as `*World`. Snake methods use pointer receivers so mutations affect the actual snake, not a copy. `*Snake` in the world map means the map holds references, not copies.  
**Key JS mapping:** objects in JS are implicitly by reference. In Go you must be explicit. There is no direct JS equivalent — this needs its own explanation.  
**Exercise arc:** demonstrate value copy vs pointer mutation → write a function that correctly mutates a snake's score → fix a broken function that uses a value receiver when it should use a pointer → safely check for nil before using a pointer → write `World.AddSnake()` that stores a `*Snake`

### Module 05 — Errors
**Concepts:** Go's `error` type, returning `(value, error)`, `errors.New()`, `fmt.Errorf()`, `if err != nil` pattern, sentinel errors, when to panic vs return error  
**Slither connection:** all Redis operations return errors. WebSocket reads return errors. Parsing device IDs can fail. Every I/O boundary uses this pattern.  
**Key JS mapping:** `try/catch` → explicit `if err != nil`, `throw new Error()` → `return errors.New()`, `Promise.reject()` → returning a non-nil error  
**Exercise arc:** write a function that returns an error for invalid input → chain two functions that both return errors → wrap an error with context using `fmt.Errorf` → write a lookup that returns a sentinel "not found" error → handle errors from a simulated Redis-style store

### Module 06 — Goroutines and Channels
**Concepts:** `go` keyword, what a goroutine is, unbuffered vs buffered channels, sending and receiving, `select`, goroutine lifecycle  
**Slither connection:** the game loop runs as a goroutine. The WebSocket hub uses channels for register/unregister/broadcast. Each player connection is its own goroutine.  
**Key JS mapping:** `setTimeout(fn, 0)` / Web Workers → goroutines (much lighter), `EventEmitter` → channels, `Promise.race()` → `select`  
**Exercise arc:** launch a goroutine and receive its result via channel → build a simple ticker that sends values on a channel every N ms → use select to handle two channels → build a mini broadcast system (one sender, multiple receivers) → implement a simplified version of the game loop that ticks and sends state

### Module 07 — Mutexes and Shared State
**Concepts:** data races, `sync.Mutex`, `sync.RWMutex`, `Lock/Unlock`, `RLock/RUnlock`, `defer` for unlocking, why you need this alongside goroutines  
**Slither connection:** `World` has a `sync.RWMutex`. `Tick()` takes a write lock. `Snapshot()` for broadcasting takes a read lock. Without this, concurrent goroutines corrupt the world state.  
**Key JS mapping:** JS is single-threaded so this problem doesn't exist there. There is no direct equivalent — explain why Go's concurrency model requires it.  
**Exercise arc:** demonstrate a data race with a shared counter (run with `-race` flag) → fix it with a mutex → implement a safe counter struct → implement a read-heavy cache that uses RWMutex → write a simplified `World` struct with mutex-protected `AddSnake`, `RemoveSnake`, and `Snapshot`

### Module 08 — Interfaces
**Concepts:** interface declaration, implicit implementation (no `implements` keyword), the empty interface, type assertions, using interfaces for testability  
**Slither connection:** `Store` interface abstracts Redis so it can be swapped with an in-memory store for tests. Message types can share an interface. Allows the game to not care whether storage is Redis or in-memory.  
**Key JS mapping:** TypeScript `interface` → Go `interface` (but Go's is implicit — a type satisfies an interface just by having the right methods, no declaration needed)  
**Exercise arc:** define a `Storer` interface → implement it with an in-memory map → write a function that accepts `Storer` not a concrete type → add a second implementation → write a type assertion → replace a concrete Redis dependency with the interface

### Module 09 — JSON and HTTP
**Concepts:** `encoding/json`, struct tags, `Marshal`/`Unmarshal`, `net/http`, `http.HandleFunc`, `http.ListenAndServe`, reading request body, writing JSON response  
**Slither connection:** world state is serialized to JSON and broadcast to clients every tick. The `/ws` endpoint upgrades HTTP to WebSocket. A `/leaderboard` endpoint returns JSON.  
**Key JS mapping:** `JSON.stringify` → `json.Marshal`, `JSON.parse` → `json.Unmarshal`, `express.get()` → `http.HandleFunc`, struct tags `` `json:"name"` `` → like `@JsonProperty` or TS key renaming  
**Exercise arc:** define a struct with JSON tags → marshal it to JSON → unmarshal JSON back into a struct → handle missing/wrong fields gracefully → write an HTTP handler that returns JSON → write a handler that reads JSON from the request body

### Module 10 — WebSockets
**Concepts:** WebSocket upgrade with `gorilla/websocket`, read/write loops, the Hub pattern, registering and unregistering clients, broadcasting, graceful disconnection  
**Slither connection:** this is the entire multiplayer layer. Every player is a WebSocket connection. The hub manages all of them. The game loop sends snapshots through the hub to every client.  
**Key JS mapping:** `new WebSocket(url)` (client) → `websocket.Upgrader` (server side of the same connection), `ws.onmessage` → the read loop goroutine, `ws.send()` → `conn.WriteMessage()`  
**Exercise arc:** upgrade an HTTP request to WebSocket → read a message and echo it back → add the connection to a hub → broadcast a message to all connected clients → handle disconnection cleanly → wire a ticker to broadcast a "ping" to all clients every second

---

## Module Creation Rules (Repeat for Emphasis)

- **One module at a time.** Wait to be asked before creating the next one.
- **Never write the implementation in exercise.go.** Stubs only. The developer writes the code.
- **Tests must fail on stubs.** If a zero value accidentally passes a test, that test is broken — fix it.
- **solution/exercise.go is complete and idiomatic.** It's the gold standard they compare against after finishing.
- **Keep exercises grounded in Slither.** Every function should feel like it could live in the real game. No abstract "write a linked list" nonsense.
- **If the developer is stuck,** give hints progressively — first a conceptual hint, then a structural hint, then a code snippet if still stuck. Don't just give the answer.
- **If the developer's solution works but isn't idiomatic,** tell them — show them the idiomatic way and explain why Go developers write it that way.

---

## When All 10 Modules Are Done

The developer now has everything they need to build the real game. At that point:

1. Scaffold the real project structure (as described in `SLITHER_PROJECT.md` if present)
2. Reference back to specific modules when writing each piece — "this is the mutex pattern from module 07"
3. The `learn/` directory stays in the repo as a reference

---

## Directory Layout This System Creates

```
slither/                          ← game repo root
├── AGENT_LEARNING.md             ← this file (the brain)
├── SLITHER_PROJECT.md            ← full project spec (if present)
├── learn/
│   ├── README.md                 ← created when first module is requested
│   ├── 01-go-basics/
│   │   ├── README.md
│   │   ├── exercise.go
│   │   ├── exercise_test.go
│   │   └── solution/
│   │       └── exercise.go
│   ├── 02-structs-and-methods/
│   │   └── ...
│   └── ...                       ← created on demand as developer progresses
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── game/
│   ├── ws/
│   └── store/
└── web/                          ← React frontend
```

---

## One Final Note

The developer is building this to learn, not just to ship. Don't optimize for speed. Optimize for understanding. A module that takes two hours and leaves them genuinely understanding pointers is worth more than one that takes ten minutes and leaves them copy-pasting. Make them feel the concepts, not just see them.
