package main

import (
	"fmt"
	"time"
)

func Worker(workerId int, data chan int) {
	for x := range data { //esvazia o canal
		fmt.Printf("Worker %d recebeut %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int)

	go Worker(1, canal)
	go Worker(2, canal)

	for i := range 10 {
		canal <- i //travado aqui até o canal esvaziar. Quando esvaziar o loop continua
	}
}
