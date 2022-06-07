package main

import "time"

/*
Race conditions.
Race condition is a state where the outcome of some source code is non-deterministic and
typically depends on the low level of interleaving done by the os.

In this example two functions writes the same variable 'raceOn'.
Both functions are invoked using the 'go' keyword which makes it possible for the
Go runtime to schedule them concurrently. Since there is no control mechanism in
place controlling which function actually writes the value first, the result
differs between different invocations.
*/

var raceOn int = 0

func main() {
	go hello1()
	go hello2()
	time.Sleep(100 * time.Millisecond)
	println("Hello from main, value of raceOn is", raceOn)
}

func hello1() {
	raceOn = 1
}

func hello2() {
	raceOn = 2
}
