# Go Concurrency Essentials

These examples practise basic concurrency in Go:

- `atomic.go`: updates a shared counter using atomic operations.
- `mutex.go`: protects a shared counter using a mutex.
- `signalling.go`: synchronises two goroutines using a channel.
- `semaphore.go`: uses a buffered channel to limit concurrent tasks.
- `sem-ex.go`: uses a semaphore to limit concurrent Collatz calculations.

## Requirements

- Go 1.27.0 or newer, as specified in `go.mod`.
- `golang.org/x/sync`, used by `sem-ex.go` and listed in `go.mod`.

## Running the examples

Open a terminal in this directory and run each example separately:

```sh
go run atomic.go
go run mutex.go
go run signalling.go
go run semaphore.go
go run sem-ex.go
```

Each file has its own `main` function, so the files cannot be built together
as one package. The atomic and mutex examples should both print `10000` as
their final result. The order of output from concurrent tasks can vary.

The examples are based on the materials provided by Dr Joseph Kehoe.
