package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Serveur en écoute sur le port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur de connexion:", err)
			continue
		}

		fmt.Println("Nouveau client connecté :", conn.RemoteAddr())

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Client déconnecté :", conn.RemoteAddr())
			return
		}

		message := string(buffer[:n])
		fmt.Print("Reçu du client : ", message)

		reponse := "Serveur dit : " + message
		conn.Write([]byte(reponse))
	}
}
