package main

import (
	"fmt"
	"time"
)

func main() {
	customerOrders := []string{"cappuccino", "flat white", "americano", "americano"}

	brewChan := make(chan string, 1)
	distributeChan := make(chan string, 2)

	go func(orders []string) {
		for _, order := range orders {
			fmt.Println("Received new order: ", order)
			brewChan <- order
		}

		close(brewChan)
	}(customerOrders)

	go func() {
		for order := range brewChan {
			fmt.Println("Barista is preparing: ", order)

			if order == "flat white" {
				time.Sleep(3 * time.Second)
			}

			distributeChan <- order
		}

		close(distributeChan)
	}()

	for coffee := range distributeChan {
		fmt.Println("Happy customer with a ", coffee)
	}
}
