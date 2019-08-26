package main

import (
	yamq "github.com/kuskmen/yamq/internal/client"
)

func main() {
	client := yamq.NewClient("127.0.0.1:1234")

	client.Query("test")
}
