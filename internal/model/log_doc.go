package model

type LogDoc struct {
	Message string `bson:"message"`
	Owner   string `bson:"owner"`
}
