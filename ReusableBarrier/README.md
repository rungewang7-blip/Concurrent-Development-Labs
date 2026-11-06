# Reusable Barrier

This Go program uses an atomic counter and two channel turnstiles to make a
barrier that can be used in a loop. Five goroutines meet at the barrier in
each of three rounds. In each round, all Part A messages appear before any
Part B message for that round.

Run from this directory:

```sh
go run main.go
```

Run the multi-round check with `go test`.

The two-turnstile idea follows Dr Joseph Kehoe's reusable barrier lecture.
The Go lab template also suggests replacing the mutex with an atomic counter.

Author: Runge Wang (C00304516)

## License

This Lab 4 implementation is distributed under the GNU General Public License
version 3 or later. See the [GPLv3 license text](../Barrier/COPYING.GPLv3).
Dr Joseph Kehoe is the author of the referenced lab template and lecture; the
implementation in this directory is a new implementation.
