package pulsar

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"time"
)

type publish struct {}
var Publish = publish{}

func (p *publish) Fingerprint(ctx context.Context) error {
	if FingerprintProducer == nil {
		return errors.New("fingerprint producer not initialised")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	msg := &pulsar.ProducerMessage{
		Payload: []byte(fmt.Sprint("Hello")),
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
		Payload: []byte(fmt.Sprint("Hello")),
	}

	if _, err := MouseMovementProducer.Send(ctx, msg); err != nil {
		return err
	}

	return nil
}