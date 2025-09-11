package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/george124816/gelection/internal/app/models/election"
	"github.com/george124816/gelection/internal/db"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	// postgres example
	db.Connect()
	db.Schema()
	helloWorldDatabase()

	// kafka example
	publishMessage()
	consumeMessage()
}

func helloWorldDatabase() {
	var err error

	_ = election.Delete(3)

	err = election.Create("Humble election")
	if err != nil {
		log.Fatalln("Failed to create a election:", err)
	}

	var name string
	name, err = election.Read(4)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name)
}

func publishMessage() {
	topic := "vote"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader: ", err)

	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)

	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer: ", err)

	}

	fmt.Println("oi")
}

func consumeMessage() {
	topic := "vote"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	if err != nil {
		log.Fatal("failed to dial leader: ", err)

	}

	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6)

	b := make([]byte, 10e3)
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}

		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection", err)
	}
}
