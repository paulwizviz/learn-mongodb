package mongoutil

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllDocument(ctx context.Context, client *mongo.Client, dbname, collname string) ([]bson.M, error) {
	coll := client.Database(dbname).Collection(collname)
	cur, err := coll.Find(ctx, bson.D{})
	var result []bson.M
	for cur.Next(ctx) {
		var found bson.M
		err := cur.Decode(&found)
		if err != nil {
			log.Println(err)
		}
		result = append(result, found)
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}
