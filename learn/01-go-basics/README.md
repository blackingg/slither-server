# Module 01 — Go Basics

In this module, you will learn the absolute fundamentals of Go syntax and how it maps to what you already know in JavaScript and TypeScript.

---

## Concept Explanation

### 1. Variables and Static Typing
Go is a **statically typed** language, meaning every variable has a type known at compile time, and that type cannot change.
* In JS, you can declare `let score = 100` and later assign `score = "dead"`.
* In Go, doing this will result in a compiler error. Once a variable is an integer, it is always an integer.

### 2. Shorthand Variable Declaration `:=`
Inside functions, Go provides the `:=` operator. This declares a new variable and infers its type based on the value on the right.
```go
score := 100 // Go infers score is an int
```
*Note:* You can only use `:=` inside functions. Outside functions (at package level), you must use the `var` keyword.

### 3. Basic Types
Go has strict numeric types. Unlike JS where all numbers are `Number` (64-bit floats under the hood), Go distinguishes between:
* `int`: Whole numbers (size depends on platform, usually 64-bit).
* `float64`: Double-precision floating-point numbers (used for decimal coordinates, speed calculations, etc.).
* `string`: Text values, immutable sequence of bytes.
* `bool`: `true` or `false`.

Go **does not** automatically convert types. You cannot add an `int` to a `float64` without casting: `float64(someInt) + someFloat`.

### 4. Functions
Go functions are declared using the `func` keyword. You must specify the types of all parameters and the return type.
```go
func AddScores(a int, b int) int {
    return a + b
}
```

### 5. Control Flow: `if` / `else`
Similar to JS, but parentheses around the condition are **omitted**. Curley braces `{}` are **mandatory**, even for single-line statements.
```go
if speed > maxSpeed {
    speed = maxSpeed
} else {
    // do something else
}
```

### 6. The `for` Loop and `range`
Go has **only one** looping keyword: `for`.
* **Standard loop**: `for i := 0; i < 10; i++ { ... }`
* **While-like loop**: `for condition { ... }`
* **Infinite loop**: `for { ... }`
* **Iterating slices (arrays)**: Go uses the `range` keyword. `range` returns both the index and the value. If you don't need the index, you use the blank identifier `_` to discard it.
```go
for index, val := range slice {
    // use index and val
}

for _, val := range slice {
    // discard index, use val
}
```

### 7. String Formatting with `fmt.Sprintf`
In JavaScript, you format strings using template literals: `` `Player ${name} has ${points} pts` ``.
In Go, you use `fmt.Sprintf` from the standard library with formatting verbs:
* `%s` for strings
* `%d` for decimal integers
* `%f` for floating-point numbers (e.g. `%.2f` for two decimal places)
* `%v` for default value formatting of any type

```go
msg := fmt.Sprintf("Player %s has %d pts", name, points)
```

---

## JS → Go Mapping

| JS / TS | Go |
| :--- | :--- |
| `const name = "Player1";` | `name := "Player1"` *(constants exist in Go via `const`, but `:=` is standard for local vars)* |
| `let score = 100;` | `var score int = 100` *or* `score := 100` |
| `function greet(name) { return "Hi " + name; }` | `func greet(name string) string { return "Hi " + name }` |
| `for (let i = 0; i < 5; i++) { ... }` | `for i := 0; i < 5; i++ { ... }` |
| `for (const x of arr) { ... }` | `for _, x := range arr { ... }` |
| `` `Rank ${rank}: ${name}` `` | `fmt.Sprintf("Rank %d: %s", rank, name)` |

---

## Why This Exists in Go

* **No implicit type coercion:** Prevents bizarre JavaScript issues like `"" + 1` returning `"1"`, or `[] + {}` returning `[object Object]`. Everything is explicit and clean.
* **Unified Loop Syntax (`for`):** Makes the language simpler to read and write. You don't have to choose between `for`, `while`, `do-while`, or `.forEach()`.
* **The Blank Identifier (`_`):** Go forbids unused variables. If a function or loop returns a value you don't need, you **must** assign it to `_`. This forces you to be explicit about what data you are discarding.

---

## Where This Appears in Slither

1. **Clamping vectors:** Ensuring client-sent coordinates don't move the snake faster than the maximum speed threshold.
2. **Score processing:** Iterating through all particles the snake ate in a tick, summing their value to increase the player's score.
3. **Leaderboard formatting:** Compiling a formatted string for each slot in the top 10 rankings to send over WebSockets to the client.

---

## How to Run Tests

Navigate to this module's directory and run the tests:

```powershell
cd learn/01-go-basics
go test -v
```

Alternatively, from the root of the repository:
```powershell
go test -v ./learn/01-go-basics/...
```
