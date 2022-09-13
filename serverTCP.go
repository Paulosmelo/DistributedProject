package main

import (
	"fmt"
	"net"
	"os"
	"distributed_project/vendas"
)

func main() {
	HelloServerTCP()
}

func HelloServerTCP() {
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {fmt.Printf(err.Error())}

	ln, err := net.ListenTCP("tcp", r)
	if err != nil {fmt.Printf(err.Error())}

	fmt.Println("Server listening on:", r)

	for{
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(0)
		}
		go HandleTCPConnection(conn)
	}
}


func HandleTCPConnection(conn net.Conn){
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {fmt.Printf("Error Reading message.\n")}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf(err.Error())			
			os.Exit(0)
		}
	}(conn)
	
	r := vendas.Vendas{}.GetVendas(string(buffer[:mLen]))	
	fmt.Println("Vendas at day", string(buffer[:mLen]), "was", r)
}

