# Learning Plan: Concurrent Port Scanner in Go

## Phase 1: Scan a Single Port

**Goal**: Write a program that checks if ONE port is open on a host.

**Concepts you'll learn**:
- Go program structure (`package main`, `func main()`)
- Importing packages
- Variables and types
- TCP connections with `net.DialTimeout`
- Basic error handling

**Milestone**: Running the program and seeing "Port 80 is open" or "Port 80 is closed".

---

## Phase 2: Scan Multiple Ports Sequentially

**Goal**: Extend the program to scan a range of ports, one after another.

**Concepts you'll learn**:
- `for` loops (Go's only loop construct)
- Slices (Go's dynamic arrays)
- Collecting and sorting results
- `fmt.Sprintf` for string formatting

**Milestone**: Scanning ports 1-100 and listing all open ones (slowly, one by one).

---

## Phase 3: Add Concurrency with Goroutines

**Goal**: Scan multiple ports at the same time using goroutines.

**Concepts you'll learn**:
- Goroutines (`go` keyword)
- Why unbounded concurrency is dangerous
- `sync.WaitGroup` for waiting on goroutines to finish
- Channels for sending results between goroutines
- The difference between buffered and unbuffered channels

**Milestone**: Scanning ports much faster, but noticing it can overwhelm the system.

---

## Phase 4: Implement the Worker Pool Pattern

**Goal**: Control how many goroutines run at once using a worker pool.

**Concepts you'll learn**:
- The worker pool pattern (a fixed number of goroutines pulling from a shared queue)
- Channel direction types (`chan<-` send-only, `<-chan` receive-only)
- Closing channels to signal "no more work"
- Coordinating producers, workers, and consumers

**Milestone**: Scanning ports fast AND controlled, with a configurable number of workers.

---

## Phase 5: CLI Flags and Polish

**Goal**: Make the scanner a proper command-line tool.

**Concepts you'll learn**:
- `flag` package for parsing command-line arguments
- Pointers (flags return pointers)
- `time.Duration` type
- Input validation
- Measuring execution time with `time.Now()` and `time.Since()`
- Structs (for grouping port + open/closed into a `Result` type)

**Milestone**: A polished CLI tool matching the example in the readme.

---

## How to Use This Plan

- Work through each phase in order
- Run your code at each milestone to see it working
- Don't move on until you understand the current phase
- Experiment! Change values, break things, see what happens
