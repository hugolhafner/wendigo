package kafka

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hugolhafner/wendigo/internal/env"
)

type publish struct{}

var Publish = publish{}

func (p *publish) Fingerprint(ctx context.Context, uniqueHash uint64, body *[]byte) error {
	if Producer == nil {
		return errors.New("fingerprint producer not initialised")
	}

	_, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &env.FingerprintTopic, Partition: kafka.PartitionAny},
		Value:          *body,
		Key:            []byte(strconv.FormatUint(uniqueHash, 10)),
	}

	if err := Producer.Produce(msg, nil); err != nil {
		return err
	}

	return nil
}

func (*publish) MouseMovements(ctx context.Context, body *[]byte) error {
	if Producer == nil {
		return errors.New("mouse movement producer not initialised")
	}

	_, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &env.MouseMovementTopic, Partition: kafka.PartitionAny},
		Value:          *body,
	}

	if err := Producer.Produce(msg, nil); err != nil {
		return err
	}

	return nil
}
