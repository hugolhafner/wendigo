package env

import (
	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

var (
	PORT                  = genv.Key("PORT").Default(3000).Int()
	KafkaServers          = genv.Key("KAFKA_SERVERS").String()
	KafkaSecurityProtocol = genv.Key("KAFKA_SECURITY_PROTOCOL").Default("SASL_SSL").String()
	KafkaMechanism        = genv.Key("KAFKA_MECHANISM").Default("PLAIN").String()
	KafkaUsername         = genv.Key("KAFKA_USERNAME").String()
	KafkaPassword         = genv.Key("KAFKA_PASSWORD").String()
	MouseMovementTopic    = genv.Key("MOUSE_MOVEMENT_TOPIC").Default("mouse-movements").String()
	FingerprintTopic      = genv.Key("FINGERPRINT_TOPIC").Default("fingerprints").String()
)
