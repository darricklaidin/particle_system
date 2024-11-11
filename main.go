package main

import (
	"fmt"
	"particle_system/particles"
	"time"
)

func main() {
	coffee := particles.NewCoffee(5, 3)
	// coffee := particles.NewCoffee(100, 3)
	coffee.Start()

	timer := time.NewTicker(100 * time.Millisecond)
	for {
		<-timer.C
		fmt.Print("\033[H\033[2J")
		coffee.Update()
		fmt.Println(coffee.Display())
	}
}
