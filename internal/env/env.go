package env

import (
	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

var (
	PORT      = genv.Key("PORT").Default(3000).Int()
	PulsarToken = genv.Key("PULSAR_TOKEN").String()
	PulsarUri = genv.Key("PULSAR_URI").String()
	MouseMovementTopic = genv.Key("MOUSE_MOVEMENT_TOPIC").Default("persistent://testinger/default/mouse-movements").String()
	FingerprintTopic = genv.Key("FINGERPRINT_TOPIC").Default("persistent://testinger/default/fingerprints").String()
)
