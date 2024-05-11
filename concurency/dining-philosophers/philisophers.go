package main

import (
	"fmt"
	"sync"
	"time"
)

type Philisopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philsophers = []Philisopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3 // how many times a person eats?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	fmt.Println("Dining Philosopher Problem")
	fmt.Println("--------------------------")
	fmt.Println("The table is empty")

	dine()

	fmt.Println("The table is empty")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philsophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philsophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philsophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philsophers); i++ {
		go diningProblem(philsophers[i], wg, forks, seated)
	}
}

func diningProblem(philisopher Philisopher, wg *sync.WaitGroup)
