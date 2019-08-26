package main

import (
	yamq "github.com/kuskmen/yamq/internal/server"
)

func main() {
	yamq.NewServer(":1234")
}
