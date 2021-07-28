package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/hugolhafner/wendigo/internal/env"
	"log"
	"time"
)

var Client pulsar.Client
var FingerprintProducer pulsar.Producer
var MouseMovementProducer pulsar.Producer

func Init() {
	token := pulsar.NewAuthenticationToken(env.PulsarToken)

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var err error
	Client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL:                     env.PulsarUri,
		Authentication:          token,
		TLSTrustCertsFilePath:   "", // TODO: This
	})

	if err != nil {
		log.Fatalf("Unable to initialise pulsar client: %v", err)
	}

	FingerprintProducer, err = Client.CreateProducer(pulsar.ProducerOptions{
		Topic: env.FingerprintTopic,
	})

	if err != nil {
		log.Fatalf("Unable to create fingerprint producer: %v", err)
	}

	MouseMovementProducer, err = Client.CreateProducer(pulsar.ProducerOptions{
		Topic: env.MouseMovementTopic,
	})

	if err != nil {
		log.Fatalf("Unable to create mouse movement producer: %v", err)
	}
}