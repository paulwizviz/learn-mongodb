package mongoutil

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDocument(ctx context.Context, client *mongo.Client, dbname, collname string, data interface{}) error {

	result, err := client.Database(dbname).Collection(collname).InsertOne(ctx, data)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
