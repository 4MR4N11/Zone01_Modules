package net_cat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	clients     = make(map[net.Conn]string)
	clientLock  sync.Mutex
	clientCount int
	countLock   sync.Mutex
	history     string
)

func printLogo(c net.Conn) {
	c.Write([]byte("Welcome to TCP-Chat!\n"))
	c.Write([]byte("         _nnnn_\n"))
	c.Write([]byte("        dGGGGMMb\n"))
	c.Write([]byte("       @p~qp~~qMb\n"))
	c.Write([]byte("       M|@||@) M|\n"))
	c.Write([]byte("       @,----.JM|\n"))
	c.Write([]byte("      JS^\\__/  qKL\n"))
	c.Write([]byte("     dZP        qKRb\n"))
	c.Write([]byte("    dZP          qKKb\n"))
	c.Write([]byte("   fZP            SMMb\n"))
	c.Write([]byte("   HZM            MMMM\n"))
	c.Write([]byte("   FqM            MMMM\n"))
	c.Write([]byte(" __| \".        |\\dS\"qML\n"))
	c.Write([]byte(" |    `.       | `' \\Zq\n"))
	c.Write([]byte("_)      \\.___.,|     .'\n"))
	c.Write([]byte("\\____   )MMMMMP|   .'\n"))
	c.Write([]byte("     `-'       `--'\n"))
	c.Write([]byte("[ENTER YOUR NAME]:"))
}

func checkValidName(name string) int {
	if len(name) == 0 || name == "\n" {
		return 1
	}
	for _, c := range name {
		if c == ' ' || c == '\t' || c == '\b' {
			return 1
		}
	}
	return 0
}

func broadcastMessage(message string, sender string) {
	clientLock.Lock()
	history += message
	defer clientLock.Unlock()
	for conn := range clients {
		if clients[conn] == sender {
			continue
		} else {
			now := time.Now()
			layout := "2006-01-02 15:04:05"
			formattedTime := now.Format(layout)
			resp := fmt.Sprintf("[%s][%s]:", formattedTime, clients[conn])
			_, err := conn.Write([]byte(message + "\n" + resp))
			if err != nil {
				log.Printf("Error sending message to client: %v", err)
				conn.Close()
				delete(clients, conn)
			}
		}
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	countLock.Lock()
	writer := bufio.NewWriter(c)
	reader := bufio.NewReader(c)
	if clientCount >= 10 {
		writer.WriteString("Sorry, the room is full. Try later.\n")
		writer.Flush()
		countLock.Unlock()
		c.(*net.TCPConn).SetLinger(0)
		return
	}
	clientCount++
	countLock.Unlock()
	printLogo(c)
	name := ""
	var err error
	for {
		name, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		name = strings.TrimSpace(name)
		if checkValidName(name) == 1 {
			c.Write([]byte("[ENTER YOUR NAME]:"))
		} else {
			break
		}
	}
	clientLock.Lock()
	clients[c] = name
	c.Write([]byte(history + "\n"))
	clientLock.Unlock()
	broadcastMessage(fmt.Sprintf("\n%s has joined our chat...", name), name)
	sendPrompt := func(prompt string) {
		writer.WriteString(prompt)
		writer.Flush()
	}
	for {
		now := time.Now()
		layout := "2006-01-02 15:04:05"
		formattedTime := now.Format(layout)
		resp := fmt.Sprintf("[%s][%s]:", formattedTime, name)
		sendPrompt(resp)
		message, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %v", err)
			}
			break
		}
		message = strings.TrimSpace(message)
		if len(message) != 0 {
			broadcastMessage(fmt.Sprintf("\n%s %s", resp, message), name)
		}
	}
	clientLock.Lock()
	delete(clients, c)
	clientLock.Unlock()
	countLock.Lock()
	clientCount--
	countLock.Unlock()
	broadcastMessage(fmt.Sprintf("\n%s has left our chat...", name), name)
}

func Listen(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Starting server")
	for i := 0; i < 60; i++ {
		fmt.Printf(".")
		time.Sleep(time.Second / 20)
	}
	fmt.Printf(". ;)\n\n")
	time.Sleep(time.Second)
	fmt.Println("Listening on the port: " + port)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if err.Error() == "use of closed network connection" {
				fmt.Println("test")
				break
			}
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
