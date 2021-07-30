package kafka

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hugolhafner/wendigo/internal/env"
	"github.com/sirupsen/logrus"
)

var Producer *kafka.Producer

func Init() {
	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": env.KafkaServers,
		"security.protocol": env.KafkaSecurityProtocol,
		"sasl.mechanisms":   env.KafkaMechanism,
		"sasl.username":     env.KafkaUsername,
		"sasl.password":     env.KafkaPassword,
	})

	if err != nil {
		log.Fatalf("Unable to initialise kafka producer: %v", err)
	}
}

func Shutdown(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		<-ctx.Done()
		cancel()
	}()

	// Wait 15s
	if lost := Producer.Flush(1000 * 15); lost > 0 {
		logrus.Errorf("Error flushing producer on shutdown, lost %d message(s)", lost)
	}

	wg.Done()
}
