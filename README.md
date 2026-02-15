# Go Sentinel: Concurrent Performance & Connectivity Monitor

**Go Sentinel** is a high-performance Go-based utility designed to monitor network connectivity and service availability. By leveraging Go's native concurrency primitives, it demonstrates how to build scalable backend tools that are memory-efficient and type-safe.

## 📖 Project Overview

This project serves as the foundational step in mastering **Golang** for 2026. The primary objective is to move away from the overhead of Python/Node.js and harness the power of **Goroutines** to handle high-frequency network requests simultaneously.

The utility currently operates in a **Concurrent Verification Mode**, where it pings a list of endpoints and reports status updates in real-time through a synchronized communication channel.

## ✨ Key Features

* **Goroutine-Powered Concurrency:** Executes multiple network checks in parallel rather than sequentially, significantly reducing total execution time.
* **Thread-Safe Communication:** Utilizes **Channels** (`chan`) to pass data safely between background workers and the main process, avoiding race conditions.
* **Idiomatic Error Handling:** Implements Go's strict error-checking patterns to ensure the system remains robust even when network endpoints fail.
* **Modular Structure:** Adheres to the standard Go project layout with a clear separation between command entry points and internal logic.
* **Compiled Portability:** Compiles into a single, standalone binary—eliminating the need for a runtime environment on the target machine.

## 🛠️ Tech Stack

* **Language:** Go (Golang) 1.23+
* **Standard Library:** `net/http` (Networking), `sync` (Concurrency), `time` (Performance Benchmarking).
* **Environment:** Developed within the **Strathmore University** informatics ecosystem.
* **Architecture:** Concurrent CSP (Communicating Sequential Processes) model.

## 📂 Directory Structure

```text
go-study/
├── bin/          # Compiled standalone binaries
├── cmd/          # Application entry points (app.go)
├── documentation/# Notion logs and research notes
├── package/      # Reusable internal packages
├── go.mod        # Project identity and dependency tracking
└── README.md     # Project overview
