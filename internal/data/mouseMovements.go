package data

type Metadata struct {
	ScreenX int `json:"screenX" validate:"required,number"`
	ScreenY int `json:"screenY" validate:"required,number"`
}

type Movement struct {
	X  int `json:"x" validate:"required,number"`
	Y  int `json:"y" validate:"required,number"`
	TS int `json:"ts" validate:"required,number"`
}

type MouseMovements struct {
	Collector *Collector  `json:"collector" validate:"dive"`
	Movements *[]Movement `json:"movements" validate:"min=1,dive"`
	Metadata  *Metadata   `json:"metadata" validate:"dive"`
}
