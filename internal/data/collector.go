package data

type Collector struct {
	Version string `json:"version" validate:"required"`
	Domain  string `json:"domain" validate:"required,hostname"`
}
