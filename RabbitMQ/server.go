package main

import (
	"encoding/json"
	"fmt"
    "strconv"
	"distributed_project/vendas"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {fmt.Printf(err.Error())}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {fmt.Printf(err.Error())}
	defer ch.Close()

	// declaração de filas
	requestQueue, err := ch.QueueDeclare("request", false, false, false,
		false, nil)
		if err != nil {fmt.Printf(err.Error())}

	replyQueue, err := ch.QueueDeclare("response", false, false, false,
		false, nil)
		if err != nil {fmt.Printf(err.Error())}

	// prepara o recebimento de mensagens do clientserver
	msgsFromClient, err := ch.Consume(requestQueue.Name, "", true, false,
		false, false, nil)
		if err != nil {fmt.Printf(err.Error())}

	fmt.Println("Server is ready...")
	for d := range msgsFromClient {
		
		day, err := strconv.Atoi(string(d.Body))
		if err != nil {fmt.Printf(err.Error())}

		r := vendas.VendasTCP{}.GetVendasTCP(day)

		// prepara resposta
		replyMsgBytes,err := json.Marshal(r)
		if err != nil {fmt.Printf(err.Error())}
		
		// publica resposta
		err = ch.Publish("", replyQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: []byte(replyMsgBytes)})
			if err != nil {fmt.Printf(err.Error())}
	}
}
