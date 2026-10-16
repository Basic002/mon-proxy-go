package tcp

import (
	"fmt"
	"net"
	"os"
)

func StartServer(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Serveur en écoute sur le port", port, "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur de connexion:", err)
			continue
		}

		fmt.Println("Nouveau client connecté :", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	requete := string(buffer[:n])
	fmt.Println("\n=== Nouvelle Requête HTTP ===")
	fmt.Println(requete)
	fmt.Println("=============================")

	body := "<html><body><h1>Salut depuis mon serveur Go !</h1></body></html>"
	reponse := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		fmt.Sprintf("Content-Length: %d\r\n", len(body)) +
		"Connection: close\r\n" +
		"\r\n" +
		body

	conn.Write([]byte(reponse))
}
