package main

import "fmt"
import "sync"
import "time"

type Chop struct{ sync.Mutex }

type Phil struct {
	leftHand, rightHand *Chop
	id, counter         int
}

// The host that checks the length of the c channel 
func host(c chan *Phil, mu *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c

			time.Sleep(15 * time.Millisecond)
		}
	}
 mu.Done()
}

//Eats depending on c channel length
func (p Phil) Eat(c chan *Phil, mu *sync.WaitGroup) {

	for j := 0; j < 3; j++ {
		c <- &p
		if p.counter < 3 {
			p.leftHand.Lock()
			p.rightHand.Lock()

			fmt.Println("Starting to eat", p.id)

			fmt.Println("Finishing eating", p.id)

			p.leftHand.Unlock()
			p.rightHand.Unlock()
		}
	}
	mu.Done()
}

func main() {
	
	//Create c channel
	c := make(chan *Phil, 2)
	
	//Create chopsticks
	Chops := make([]*Chop, 5)
	for i := 0; i < 5; i++ {
		Chops[i] = new(Chop)
	}
	
	//Create philosophers
	Philos := make([]*Phil, 5)
	for i := 0; i < 5; i++ {
		Philos[i] = &Phil{Chops[i], Chops[(i+1)%5], i + 1, 0}
	}

	// WaitGroup with all its go routines
	var mu sync.WaitGroup

	mu.Add(5)

	go host(c, &mu)
	go Philos[0].Eat(c, &mu)
	go Philos[1].Eat(c, &mu)
	go Philos[2].Eat(c, &mu)
	go Philos[3].Eat(c, &mu)
	go Philos[4].Eat(c, &mu)

	mu.Wait()

}
