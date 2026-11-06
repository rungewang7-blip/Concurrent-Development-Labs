# Dining Philosophers

Five philosophers think and eat using five shared forks. A philosopher needs
both neighbouring forks to eat. At most four philosophers may try to pick up
forks at the same time, preventing a circular wait. Each eats three times,
so the program can finish.

## Run

From this directory, run `go run .`. Each philosopher should print three
`is eating` messages. The order may vary.

The program uses the Go standard library and requires Go 1.23 or newer.

## Source and license

Based on Dr Joseph Kehoe's Dining Philosophers Go template (21 October 2024).
Modified by Runge Wang (C00304516). The supplied Go template states "GPL
Licence" but does not specify a version or include the full license. The
applicable version needs confirmation from the lecturer; the repository
contains [GPLv3 text](../Barrier/COPYING.GPLv3) for reference.
