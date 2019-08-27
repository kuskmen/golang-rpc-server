package main

import (
	yamq "github.com/kuskmen/yamq/pkg/server"
)

func main() {
	yamq.NewServer(":1234")
}
