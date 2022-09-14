package main

import (
	"time"
    "strconv"
    "math/rand"
	"fmt"
	"encoding/json"
	"github.com/streadway/amqp"
)

func main() {
	// conecta ao servidor de mensageria
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {fmt.Printf(err.Error())}
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	if err != nil {fmt.Printf(err.Error())}
	defer ch.Close()

	// declara as filas
	requestQueue, err := ch.QueueDeclare(
		"request", false, false, false, false, nil)
	if err != nil {fmt.Printf(err.Error())}

	replyQueue, err := ch.QueueDeclare(
		"response", false, false, false, false, nil)
	if err != nil {fmt.Printf(err.Error())}

	// cria consumidor
	msgsFromServer, err := ch.Consume(replyQueue.Name, "", true, false,
		false, false, nil)
	if err != nil {fmt.Printf(err.Error())}

	start := time.Now()
	for i := 0; i< 10; i++{

		t1 := time.Now()

		// prepara request
		rand.Seed(time.Now().UnixNano())
		var random = rand.Intn(20)
		
		msgRequestBytes := []byte(strconv.Itoa(random))
		if err != nil {fmt.Printf(err.Error())}

		// publica request
		err = ch.Publish("", requestQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})
		if err != nil {fmt.Printf(err.Error())}

		// Receive message
		x := <- msgsFromServer

		var feedback [20]int
		json.Unmarshal(x.Body, &feedback)
		fmt.Println(feedback)
		t2 := time.Now()
		
		y := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Println(y)
	}
	elapsed := time.Since(start)
	fmt.Printf("Tempo: %s \n", elapsed)
}
