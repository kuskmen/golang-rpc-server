package main

import (
	yamq "github.com/kuskmen/yamq/pkg/yamqclient"
)

func main() {
	client := yamq.NewClient("127.0.0.1:1234")

	client.Query("test")
}
