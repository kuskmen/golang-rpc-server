package main

import (
	yamqserver "github.com/kuskmen/yamq/pkg/yamqserver"
)

func main() {
	yamqserver.NewServer(":1234")
}
