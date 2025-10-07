package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/george124816/gelection/cmd/http"
	"github.com/george124816/gelection/cmd/migrate"
	otel "github.com/george124816/gelection/internal"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	sign := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	loggerProvider, err := otel.StartLogs()
	if err != nil {
		slog.Error(err.Error())
	}
	metricProvider, err := otel.StartMetrics()
	if err != nil {
		slog.Error(err.Error())
	}

	err = migrate.Migrate()
	if err != nil {
		slog.Error(err.Error())
	}

	server, err := http.Start()
	if err != nil {
		slog.Error(err.Error())
	}

	ctx, cancelContext := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancelContext()

	receivedSignal := <-sign
	slog.Warn(receivedSignal.String())

	server.Shutdown(ctx)
	loggerProvider.Shutdown(ctx)
	metricProvider.Shutdown(ctx)

}

func publishMessage() {
	topic := "vote"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		slog.Error("failed to dial leader: ", "error", err)

	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	if err != nil {
		slog.Error("failed to write messages:", "error", err)

	}

	if err := conn.Close(); err != nil {
		slog.Error("failed to close writer: ", "error", err)

	}

}

func consumeMessage() {
	topic := "vote"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	if err != nil {
		slog.Error("failed to dial leader: ", "error", err)

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
		slog.Error("failed to close batch:", "error", err)
	}

	if err := conn.Close(); err != nil {
		slog.Error("failed to close connection", "error", err)
	}
}
