# go-fibonacci-journey

A monorepository serving as a structured learning path for Go, starting with a fundamental problem (Fibonacci) and progressively introducing core Go features, problem-solving techniques, and architectural patterns.

## Learning Path Overview

This repository is designed to guide you through various aspects of Go programming, using the classic Fibonacci sequence calculation as a running example. Each phase builds upon the previous one, introducing new Go concepts and addressing increasing complexity.

*   **Phase 1: Foundations - Basic Iterative Fibonacci (`cmd/01_basic_iterative`)**
    *   **Goal:** Introduce basic Go syntax, function definition, loop structures, and standard output. Provide a quick win with a runnable program.
    *   **Go Concepts:** Go program structure (`package main`, `func main`), variable declaration, `for` loops, slices, basic function calls, `fmt` package, `go run`, `go build`.

*   **Phase 2: Robust Input & Error Handling (`cmd/02_input_flags`)**
    *   **Goal:** Make the program more user-friendly and robust by accepting command-line arguments and handling potential errors.
    *   **Go Concepts:** `flag` package for CLI arguments, `strconv` package for type conversion, Go's idiomatic error handling (`error` interface, `nil` checks), `log` package for error reporting.

*   **Phase 3: Scaling with Arbitrary Precision (`cmd/03_big_int`)**
    *   **Goal:** Address the `int` overflow limitation for larger Fibonacci numbers and introduce `math/big`.
    *   **Go Concepts:** Understanding data type limitations, `math/big` package for large number arithmetic, Go's package import mechanism, potential performance implications of arbitrary precision.

*   **Phase 4: Introducing Concurrency (Go's Strength) (`cmd/04_concurrent_generator`, `cmd/04_concurrent_calculator`)**
    *   **Goal:** Leverage Go's unique strengths by exploring concurrency for potentially parallelizing Fibonacci-related tasks.
    *   **Go Concepts:** `goroutines`, `channels`, `select` statement, `sync` package (e.g., `WaitGroup`), concurrent programming patterns.

*   **Phase 5: Exposing as a Service (`cmd/05_http_server`)**
    *   **Goal:** Wrap the Fibonacci calculation in a simple HTTP server to expose it as an API.
    *   **Go Concepts:** `net/http` package basics, handling HTTP requests, JSON encoding/decoding.

## Getting Started

To get started with this learning journey:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/aastom/go-fibonacci-journey.git
    cd go-fibonacci-journey
    ```

2.  **Run Phase 1 (Basic Iterative Fibonacci):**
    ```bash
    go run ./cmd/01_basic_iterative
    ```

3.  **Explore other phases:** Navigate to the respective `cmd/` subdirectories and follow their `main.go` comments for specific instructions.

## Contributions

This repository is primarily for guided learning. While external contributions are not actively sought for the core learning path, feedback and suggestions are welcome via issues.
