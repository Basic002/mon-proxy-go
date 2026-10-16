package main

import (
	"proxy/internal/tcp"
)

func main() {
	tcp.StartServer(":8080")
}