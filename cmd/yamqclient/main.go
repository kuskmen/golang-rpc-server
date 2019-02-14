package main

import (
	yamqclient "github.com/kuskmen/yamq/pkg/yamqclient"
)

func main() {
	client := yamqclient.NewClient("127.0.0.1:1234")

	client.Query("test")
}
