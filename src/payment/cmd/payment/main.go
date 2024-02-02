package main

import (
	"encoding/json"
	"log/slog"

	entity "github.com/dtonavitor/imersao-full-cycle/src/payment/internal/entity"
	"github.com/dtonavitor/imersao-full-cycle/src/payment/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	"context"
)

func main() {
	ctx := context.Background()
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "orders")
	for msg := range msgs {
		var orderRequest entity.OrderRequest
		err := json.Unmarshal(msg.Body, &orderRequest)
		if err != nil {
			slog.Error(err.Error())
			break
		}
		response, err := orderRequest.Proccess()
		if err != nil {
			slog.Error(err.Error())
			break
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			slog.Error(err.Error())
			break
		}

		err = rabbitmq.Publish(ctx, ch, string(responseJSON), "amq.direct")
		if err != nil {
			slog.Error(err.Error())
			break
		}
		msg.Ack(false)
		slog.Info("Order processed")
	}

}
