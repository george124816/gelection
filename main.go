package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/george124816/gelection/cmd/http"
	"github.com/george124816/gelection/cmd/migrate"
	otel "github.com/george124816/gelection/internal"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	err := otel.StartLogs()
	err = otel.StartMetrics()

	err = migrate.Migrate()
	if err != nil {
		slog.Error("Failed to migrate: ", err)
	}

	http.Start()
}

func publishMessage() {
	topic := "vote"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		slog.Error("failed to dial leader: ", err)

	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	if err != nil {
		slog.Error("failed to write messages:", err)

	}

	if err := conn.Close(); err != nil {
		slog.Error("failed to close writer: ", err)

	}

}

func consumeMessage() {
	topic := "vote"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	if err != nil {
		slog.Error("failed to dial leader: ", err)

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
		slog.Error("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		slog.Error("failed to close connection", err)
	}
}
