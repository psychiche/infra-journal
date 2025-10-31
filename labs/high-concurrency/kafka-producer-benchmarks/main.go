package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "bench-topic",
		BatchSize:    100,
		BatchTimeout: 10 * time.Millisecond,
		RequiredAcks: kafka.RequireAll,
	}
	defer writer.Close()

	ctx := context.Background()
	start := time.Now()
	const total = 1000
	for i := 0; i < total; i++ {
		msg := kafka.Message{Key: []byte(fmt.Sprintf("k-%d", i)), Value: []byte("hello"), Time: time.Now()}
		if err := writer.WriteMessages(ctx, msg); err != nil {
			panic(err)
		}
	}
	dur := time.Since(start)
	fmt.Printf("sent=%d dur=%s rps=%.2f\n", total, dur, float64(total)/dur.Seconds())
}
