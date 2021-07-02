package main

import (
	"fmt"
	"math/rand"
	"time"
)

func caballo1(can chan string) {
	time.Sleep(time.Second * 1)
	can <- "Caballo 1"
}

func caballo2(can chan string) {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	can <- "Caballo 2"
}

func caballo3(can chan string) {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	can <- "Caballo 3"
}

func main() {
	can1 := make(chan string)
	can2 := make(chan string)
	can3 := make(chan string)

	go caballo1(can1)
	go caballo2(can2)
	go caballo3(can3)

	select {
	case ganador := <-can1:
		fmt.Println("Ganador: ", ganador)
	case ganador := <-can2:
		fmt.Println("Ganador: ", ganador)
	case ganador := <-can3:
		fmt.Println("Ganador: ", ganador)
	}
}
