package main

import (
	"fmt"
	"net"
	"time"
   "math/rand"
)


func clientTCP() {
	times := [1000] time.Duration{}
	var SAMPLE_SIZE = 1000

	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {}

	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {}
	
	// loop
	for i := 0; i < SAMPLE_SIZE; i++ {

		rand.Seed(time.Now().UnixNano())
		var random = rand.Intn(20)
		// prepara request & start time
		t1 := time.Now()
			
		_, err = conn.Write([]byte(random))

		times[i] = time.Now().Sub(t1)
	}

	totalTime := time.Duration(0)
	for i := range times {
		totalTime += times[i]
	}
	fmt.Printf("Total Duration: %v [%v]", totalTime, SAMPLE_SIZE)

		// fecha conexõa
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	
}


func main() {
	go clientTCP()
	
	fmt.Scanln()
}
