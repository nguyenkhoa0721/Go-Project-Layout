package model

type RawTransaction struct {
	ID      string `bson:"_id, omitempty"`
	Payload string `bson:"payload"`
}
