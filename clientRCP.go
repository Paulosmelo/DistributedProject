package main

import (
	//"strconv"
	"fmt"
	"log"
	"net/rpc"
	"time"
	//"math/rand"
)

func client() {
	var reply string
	times := [1] time.Duration{}
	var SAMPLE_SIZE = 1

	// conecta ao servidor
	client, err := rpc.Dial("tcp", "localhost:1313")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// loop
	for i := 0; i < SAMPLE_SIZE; i++ {

		//rand.Seed(time.Now().UnixNano())
		//var random = strconv.Itoa(rand.Intn(20))
		// prepara request & start time
		t1 := time.Now()
	
		// invoca operação remota
		client.Call("Sales.GetVendasRPC", "5", &reply)
		// stop time
		times[i] = time.Now().Sub(t1)
		//fmt.Println("Vendas for day", random, "is", reply)

	}
	totalTime := time.Duration(0)
	for i := range times {
		totalTime += times[i]
	}
	fmt.Printf("Total Duration: %v [%v]", totalTime, SAMPLE_SIZE)
}

func main() {

	go client()

	fmt.Scanln()
}
