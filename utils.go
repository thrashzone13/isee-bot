package main

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func LogIfError(err error, params ...string) {
	if err != nil {
		if len(params) > 0 {
			log.Fatalln(params[0])
			return
		}
		log.Fatalln(err)
	}
}

func UnMarshalDBResult(data interface{}, model interface{}) {
	b, err := bson.Marshal(data)
	LogIfError(err)

	err = bson.Unmarshal(b, model)
	LogIfError(err)
}
