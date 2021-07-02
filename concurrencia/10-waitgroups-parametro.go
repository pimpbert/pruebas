package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go imprime(&wg)
	wg.Wait()
	fmt.Println("Olas vienen olas van")
}

func imprime(wg *sync.WaitGroup) {
	fmt.Println("soy una gorutine")
	wg.Done()
}
