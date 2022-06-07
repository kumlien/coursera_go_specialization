package main

import (
	"fmt"
	"sync"
)

func main() {
	CSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopStick)
	}

	channel := make(chan string)

	host := Host{0, new(sync.Mutex)}
	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{CSticks[i],
			CSticks[(i+1)%5], &host, i + 1, channel}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

	for i := 0; i < 5; i++ {
		msg := <-channel
		fmt.Println(msg)
	}
}

func (p Philosopher) eat() {
	for i := 0; i < 3; i++ {
		p.host.waitForPermission()
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("starting to eat ", p.number)
		fmt.Println("finishing eating ", p.number)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		p.host.doneEating()
	}
	p.channel <- "Philosopher " + fmt.Sprint(p.number) + " is done eating"
}

func (h Host) waitForPermission() {
	for {
		h.mutex.Lock()
		if h.counter == 2 {
			h.mutex.Unlock()
			continue
		}
		break
	}
	h.counter++
	h.mutex.Unlock()
}

func (h Host) doneEating() {
	h.mutex.Lock()
	h.counter--
	h.mutex.Unlock()
}

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	leftCS, rightCS *ChopStick
	host            *Host
	number          int
	channel         chan string
}

type Host struct {
	counter int
	mutex   *sync.Mutex
}
