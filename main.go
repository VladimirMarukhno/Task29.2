package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	natChan := make(chan int, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		go func() {
			tmp := rand.Intn(100)
			natChan <- tmp
		}()

		select {
		case nat := <-natChan:
			fmt.Println("ввод:", nat)
			fmt.Println("квадрат:", nat*nat)
		case sig := <-sigChan:
			fmt.Println("Сигнал стоп, завершаю работу!", sig)
			return
		}
	}
}
