package model

type Effect struct {
	TriggerEventID string `json:"trig"`
	FuncScript     string `json:"script"`
}
