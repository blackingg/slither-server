# Slither Server — Go Backend Learning Project

I created this repository as a hands-on way to learn Go. My preferred learning style has always been getting my hands dirty—building, breaking, and rebuilding things to understand how they work under the hood. 

The goal of this project is to learn Go properly by building the backend server for a **massively multiplayer Slither.io clone** from scratch.

> [!TIP]
> Looking for the frontend client? You can find it here: [slither-client](https://github.com/blackingg/slither-client).

---

## 🚀 The Project: Slither Server

This is the backend server for a real-time multiplayer snake game where players compete in a shared 4000×4000 grid.

* **Stack:** Go (Backend) · WebSockets · Redis · Fly.io
* **Features:** Head-to-body collisions, custom particle generation on player death, and state broadcasting.

---

## 📚 Learn Go With Me!

If you want to use this repository to learn Go too, there is a **10-module curriculum** built directly into the codebase. You can complete the exercises locally while unit tests verify your progress.

### How to use the learning system:
1. **Explore the Curriculum:** Start in the [/learn](learn/README.md) directory.
2. **Pick a Module:** Read the concepts and JS-to-Go syntax mappings in the module's README (e.g., [Module 01: Go Basics](learn/01-go-basics/README.md)).
3. **Solve Exercises:** Implement the stubbed-out functions in `exercise.go`.
4. **Run Tests:** Verify your solution:
   ```bash
   go test -v ./learn/01-go-basics/...
   ```
5. **Compare Solutions:** Once your tests pass, check the reference code in `solution/exercise.go` to compare your approach with idiomatic Go.
