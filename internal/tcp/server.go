package tcp

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func StartServer(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Erreur lors du démarrage :", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Proxy en écoute sur le port", port, "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	buffer := make([]byte, 4096)
	n, err := clientConn.Read(buffer)
	if err != nil {
		return
	}

	requete := string(buffer[:n])
	lignes := strings.Split(requete, "\n")
	premiereLigne := strings.TrimSpace(lignes[0])
	fmt.Println("->", premiereLigne)

	parties := strings.Split(premiereLigne, " ")
	if len(parties) < 3 {
		return
	}
	methode := parties[0]
	cible := parties[1]

	if methode != "CONNECT" {
		clientConn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\nProxy HTTP classique non supporte, utilisez le HTTPS."))
		return
	}

	serverConn, err := net.Dial("tcp", cible)
	if err != nil {
		fmt.Println("Erreur de connexion à la cible :", err)
		return
	}
	defer serverConn.Close()

	clientConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

	go io.Copy(clientConn, serverConn)
	io.Copy(serverConn, clientConn)
}