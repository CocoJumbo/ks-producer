package models

type SimpleMessage struct {
	ID   string `json:"id"`
	Data string `json:"data"`
	Key  string `json:"key"`
}
