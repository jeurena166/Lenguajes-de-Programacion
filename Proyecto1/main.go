package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

const (
	songStorageDir = "song_storage"
	serverPort     = "9988" // Change this port if needed
)

func main() {
	// Create the directory for storing songs if it doesn't exist
	if err := os.MkdirAll(songStorageDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating song storage directory: %v\n", err)
		return
	}

	// Start the server
	listener, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server is listening on port %s\n", serverPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client connected from %s\n", clientAddr)

	// Send a list of song names to the client
	songList, err := getSongs()
	if err != nil {
		fmt.Printf("Error sending song list to %s: %v\n", clientAddr, err)
		return
	}
	_, err = conn.Write([]byte(strings.Join(songList, "\n")))
	if err != nil {
		fmt.Printf("Error sending song list to %s: %v\n", clientAddr, err)
		return
	}

	// Read client requests
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Client %s disconnected: %v\n", clientAddr, err)
			return
		}

		request := string(buffer[:n])
		fmt.Printf("Received request from %s: %s\n", clientAddr, request)

		// Handle client requests
		songFileName := filepath.Join(songStorageDir, request)
		songFile, err := os.Open(songFileName)
		if err != nil {
			fmt.Printf("Error opening requested song %s for client %s: %v\n", request, clientAddr, err)
			continue
		}
		defer songFile.Close()

		_, err = io.Copy(conn, songFile)
		if err != nil {
			fmt.Printf("Error sending requested song %s to client %s: %v\n", request, clientAddr, err)
			return
		}
	}
}

func getSongs() ([]string, error) {
	files, err := filepath.Glob(filepath.Join(songStorageDir, "*.mp3"))
	if err != nil {
		return nil, err
	}

	songList := make([]string, len(files))
	for i, file := range files {
		songList[i] = filepath.Base(file)
	}
	return songList, nil
}
