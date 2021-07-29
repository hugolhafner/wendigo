package pulsar

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type publish struct{}

var Publish = publish{}

func (p *publish) Fingerprint(ctx context.Context, uniqueHash uint64, body *[]byte) error {
	if FingerprintProducer == nil {
		return errors.New("fingerprint producer not initialised")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	msg := &pulsar.ProducerMessage{
		Payload: *body,
		Key:     fmt.Sprint(uniqueHash),
	}

	if _, err := FingerprintProducer.Send(ctx, msg); err != nil {
		return err
	}

	return nil
}

func (*publish) MouseMovements(ctx context.Context) error {
	if MouseMovementProducer == nil {
		return errors.New("mouse movement producer not initialised")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	msg := &pulsar.ProducerMessage{
		Payload: []byte("Hello There."),
	}

	if _, err := MouseMovementProducer.Send(ctx, msg); err != nil {
		return err
	}

	return nil
}
