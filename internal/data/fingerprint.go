package data

type Fingerprint struct {
	Collector *Collector `json:"collector" validate:"dive"`
	UserAgent string     `json:"userAgent" validate:"required"`
}
