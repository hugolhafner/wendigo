package data

type Collector struct {
	Version string `json:"version" validate:"required"`
	Domain  string `json:"domain" validate:"required,hostname"`
}

type Fingerprint struct {
	Collector *Collector `json:"collector" validate:"dive"`
	UserAgent string     `json:"userAgent" validate:"required"`
}
