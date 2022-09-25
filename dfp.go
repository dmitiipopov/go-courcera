/*
Implement the dining philosopher’s problem with the following constraints/modifications.
1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5.
7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
on a line by itself, where <number> is the number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"time"

	"sync"
)

var eatWgroup sync.WaitGroup

type chopstick struct{ sync.Mutex }

type philosopher struct {
	id                  int
	leftChopstick, rightChopstick *chopstick
}


func (p philosopher) eat() {
	defer eatWgroup.Done()
	for j := 0; j < 3; j++ {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		say("staring to eat", p.id)
		time.Sleep(time.Second)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		say("finishing eating", p.id)
		time.Sleep(time.Second)
	}

}

func say(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}

func main() {
	count := 5
	chopsticks := make([]*chopstick, count)
	for i := 0; i < count; i++ {
		chopsticks[i] = new(chopstick)
	}
	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, leftChopstick: chopsticks[i], rightChopstick: chopsticks[(i+1)%count]}
		eatWgroup.Add(1)
		go philosophers[i].eat()
	}
	eatWgroup.Wait()

}