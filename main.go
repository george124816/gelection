package main

import (
	"context"
	"fmt"
	"log"
	"time"

	// "github.com/george124816/gelection/internal/db"
	"github.com/george124816/gelection/cmd/http"
	"github.com/george124816/gelection/cmd/migrate"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	err := migrate.Migrate()
	if err != nil {
		log.Fatalln("Failed to migrate: ", err)
	}

	http.Start()
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
