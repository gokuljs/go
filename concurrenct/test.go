package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	fmt.Println("Processing orders")
	length := len(orders)
	for i := 0; i < length; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Order of %v is %v\n", i+1, orders[i])
	}
}

func updateOrderStatus(orders []*Orders) {
	fmt.Println("Updating order status")
	for _, order := range orders {
		time.Sleep(
			time.Duration(rand.Intn(300) * int(time.Millisecond)),
		)
		status := []string{
			"Processing", "Shipped", "Delivered",
		}[rand.Intn(3)]
		order.Status = status
		fmt.Printf("update status of orderid %v to status of %v\n", order.Id, order.Status)
	}
}

func reportOrderStatuses(orders []*Orders) {
	fmt.Println("reporting order status")
	for _, order := range orders {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("Print order status of id %v = %v\n", order.Id, order.Status)
	}
}

func main() {
	fmt.Println("Concurrency method in go")
	orders := generateOrders(10)
	processOrders(orders)
	updateOrderStatus(orders)
	reportOrderStatuses(orders)
	fmt.Println("All operations completed")
}
