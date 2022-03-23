package core

type Camera struct {
	Domain     string `json:"domain"`
	LastTimeUp int64  `json:"last-time-up"` //unix time
	Up         bool   `json:"up"`
	Image      string `json:"image"` // base64
}
