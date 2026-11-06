# Producer–Consumer

A producer creates 20 numbered events and places them in a buffered channel.
Two consumers take events from the shared buffer. Once production is finished,
the channel closes and the program waits for both consumers to finish.

## Run and test

From this directory:

```sh
go run .
go test -timeout 20s
go test -race -timeout 20s
```

The output contains each event number from 0 to 19 once; the order can vary.
Go 1.23 or newer is required. The program uses only the standard library.

## Files and source

- `main.go`: event, producer, consumer, and shared buffer.
- `main_test.go`: checks that every event is consumed exactly once.
- `go.mod`: Go module and version.

The producer, consumer, event and buffer concepts follow the C++ lab skeleton
provided by Joseph Kehoe. This Go implementation was written separately.
The supplied `Event.h` and `SafeBuffer.h` carry GPLv3-or-later notices; no
source code or notices from those files were copied into this implementation.

## Outstanding requirements

The supplied skeleton uses a buffer size of 20 and 100 threads, but it does
not state whether these numbers are mandatory. This example uses a buffer of
5 and two consumers; confirm the required values and submission details with
the lecturer if needed.

Author: Runge Wang (C00304516)
