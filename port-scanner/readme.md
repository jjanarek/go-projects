# Go Project: Concurrent Port Scanner (Worker Pool)

## Overview

This project is a command-line TCP port scanner written in Go.

It demonstrates:

- ✅ Worker pool pattern
- ✅ Bounded concurrency
- ✅ TCP networking
- ✅ CLI flag parsing
- ✅ Channel-based coordination
- ✅ Proper goroutine lifecycle management

The scanner:

- Accepts a host
- Scans a configurable port range
- Uses a fixed number of worker goroutines
- Reports open ports
- Prints total scan time

---

## Example Usage

```bash
go run main.go --host scanme.nmap.org --start 1 --end 1024 --workers 50
```

Example output:

```
Scanning scanme.nmap.org (1–1024) with 50 workers...

22 OPEN
80 OPEN
443 OPEN

Scan completed in 3.2s
3 open ports found.
```

---

## How It Works

1. CLI flags define host, port range, worker count, and timeout.
2. A `jobs` channel distributes ports to workers.
3. Each worker:
   - Attempts `net.DialTimeout`
   - Reports result via `results` channel.
4. A `WaitGroup` ensures workers finish cleanly.
5. Open ports are collected, sorted, and printed.

This design prevents spawning thousands of goroutines at once,
which is important for real-world scalability.

---

## What You Learn From This Project

- Why bounded concurrency matters
- How to implement worker pools idiomatically
- How channels coordinate concurrent systems
- How Go handles network I/O
- How to build a practical CLI tool

---

## Possible Enhancements

- Add `--verbose` to show closed ports
- Add progress indicator
- Add context cancellation (`Ctrl+C`)
- Output JSON format
- Benchmark different worker counts
- Write unit tests for validation logic

---

## Build Binary

```bash
go build -o portscan
./portscan --host example.com --start 1 --end 1000 --workers 100
```

---

This project represents a transition from "learning Go syntax"
to "writing production-style Go programs".
