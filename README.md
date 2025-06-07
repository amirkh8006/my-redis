# MyRedis (A Simple Redis-like In-Memory Database in Go)

This is a lightweight, Redis-like in-memory key-value store written in Go. It supports basic `Set`, `Get`, `Persist`, and `Load` operations using a map with concurrency safety.

## Features

- In-memory key-value store
- Thread-safe with `sync.RWMutex`
- Persistence using Go's `gob` encoding
- Simple file-based data loading and saving


## Usage

### Run the project

Make sure you have [Go installed](https://go.dev/doc/install).

```bash
go run ./cmd/main.go
