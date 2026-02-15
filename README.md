# Go Sentinel: Concurrent Performance & Connectivity Monitor

**Go Sentinel** is a high-performance utility designed to monitor network services. Developed as part of a deep dive into the Go ecosystem, it focuses on building non-blocking, memory-efficient systems suitable for large-scale informatics projects.

## 📖 Project Overview

This project implements a **Concurrent Web Scraper Engine**. By moving away from traditional class-based inheritance and adopting Go’s composition model, the engine achieves high throughput with minimal overhead.

## ✨ Key Features

* **Concurrency-First Architecture:** Utilizes Goroutines and Channels to scale network requests without blocking the main execution thread.
* **Defensive Design:** Implements Go's "Verbose Guardrails," ensuring every `nil` error is audited to prevent memory leaks and security vulnerabilities.
* **High-Speed Compilation:** Leverages the Go feedback loop for rapid iterative design, bridging the gap between scripting ease and compiled performance.
* **Standardized Layout:** Adheres to a minimalist directory structure for better package management and reusability.

## 🛠️ Tech Stack

* **Language:** Go (Golang) 1.23+
* **Standard Library:** `net/http` (Networking), `sync` (Concurrency), `time` (Benchmarking).
* **Environment:** Strathmore University Informatics & Computer Science.

## 📂 Final Directory Map

```text
AI-Learning-Go-programming/ 
├── bin/                 # Compiled standalone executables
├── cmd/                 
│   └── app.go           # Entry point: Concurrent Site Checker
├── documentation/       # Research logs & Notion backups
├── package/             # Reusable internal logic
├── go.mod               # Project identity & dependency management
└── README.md            # Project documentation
