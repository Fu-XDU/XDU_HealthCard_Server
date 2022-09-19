package model

type Summary struct {
	HealthCard *HealthCard `json:"health_card"`
	ThreeCheck *ThreeCheck `json:"three_check"`
}
