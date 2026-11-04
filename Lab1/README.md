# Lab 1 Semaphores

Three small Go examples based on Dr Joseph Kehoe's semaphore lecture:

- `signalling`: the second task waits for a signal from the first.
- `rendezvous`: A1 and B1 must both finish before A2 or B2 starts.
- `mutual-exclusion`: a one-token channel protects a shared counter.

The buffered channels act as simple semaphores. Run these commands from the
`Lab1` directory:

```sh
go run ./signalling
go run ./rendezvous
go run ./mutual-exclusion
```

The counter should finish at 100000. The order of A1 and B1 can vary, as can
the order of A2 and B2.

The supplied 2023 lecture asks for C/C++, a Makefile and Doxygen. These examples
use Go, following the current preference discussed for the labs; whether the
older submission instructions still apply needs to be checked with the lecturer.

Runge Wang (C00304516)
