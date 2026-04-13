package main

import "fmt"

type Orders struct {
	Id     int
	Status string
}

const orderPendingStatus string = "pending"

func generateOrders(count int) []*Orders {
	orderArray := make([]*Orders, count)
	for i := 0; i < count; i++ {
		orderArray[i] = &Orders{
			Id:     i + 1,
			Status: orderPendingStatus,
		}
	}
	return orderArray
}

func processOrders(orders []*Orders) {
	length := len(orders)
	for i := 0; i < length; i++ {
		fmt.Println("Order of %v is %v", i+1, orders[i])
	}
}

func main() {
	fmt.Println("Concurrency method in go")
	orders := generateOrders(10)
	processOrders(orders)
}
