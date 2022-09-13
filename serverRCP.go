package main

import (
	"fmt"
	"net"
	"os"
	"distributed_project/vendas"
	"net/rpc"
)

func server(){

	vendas := new(vendas.VendasMOM)

	// cria um novo server RPC
	server := rpc.NewServer()
	err := server.RegisterName("Sales", vendas)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} 
	// cria um listenet TCP
	ln, err := net.Listen("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer func(ln net.Listener) {
		var err = ln.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ln)

	// aguarda por invocações
	fmt.Println("Server is ready ...")
	server.Accept(ln)
}

func main() {

	go server()

	_, _ = fmt.Scanln()
}

