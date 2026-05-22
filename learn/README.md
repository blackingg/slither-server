# Slither Go Learning System

Welcome to the Go learning system for building the **Slither** multiplayer backend. Instead of working through dry syntax examples, you'll learn Go by solving problems directly modeled after the architecture of our game.

## Module Curriculum

| Module | Status | Topic & Key Concepts |
| :--- | :--- | :--- |
| **[01-go-basics](file:///c:/Users/Black/Documents/Dev/slither/slither-server/learn/01-go-basics/README.md)** | 🟢 **Active** | Variables, `:=`, primitive types, loops, `range`, `fmt.Sprintf` |
| **02-structs-and-methods** | 🔒 *Locked* | Structs, value/pointer receivers, methods, constructors |
| **03-slices-and-maps** | 🔒 *Locked* | Slices (append, capacity), maps (crud, safe check), iteration |
| **04-pointers** | 🔒 *Locked* | Pointers, value copy vs reference, nil safety, pointer receivers |
| **05-errors** | 🔒 *Locked* | Error type, multiple return values, `fmt.Errorf`, sentinels |
| **06-goroutines-and-channels** | 🔒 *Locked* | Concurrency, channels, `select`, game tick loop |
| **07-mutexes-and-shared-state** | 🔒 *Locked* | Data races, `sync.Mutex`, `sync.RWMutex`, `defer` locking |
| **08-interfaces** | 🔒 *Locked* | Interfaces, implicit implementation, type assertion, mock store |
| **09-json-and-http** | 🔒 *Locked* | Struct tags, `json.Marshal/Unmarshal`, standard library HTTP |
| **10-websockets** | 🔒 *Locked* | Gorilla WebSockets, Hub pattern, read/write pumps, disconnects |

## How to Progress

1. Open the active module's directory (e.g., `learn/01-go-basics/`).
2. Read the `README.md` explanation and the JavaScript-to-Go mappings.
3. Open `exercise.go` and implement the stubbed functions marked with `TODO`.
4. Run the unit tests to verify your implementation.
5. Once your tests pass, check the reference solution in `solution/exercise.go` to compare your approach with idiomatic Go.
6. Tell the AI agent when you're ready to start the next module!
