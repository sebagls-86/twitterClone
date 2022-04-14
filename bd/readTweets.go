package bd

import (
	"context"
	"log"
	"time"

	"github.com/sebagls-86/twitterClone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, page int64) ([]*models.TweetsSender, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterClone")
	col := db.Collection("tweet")

	var results []*models.TweetsSender

	condition := bson.M{
		"userId": ID,
	}

	opt := options.Find()

	opt.SetLimit(20)
	opt.SetSort(bson.D{{Key: "date", Value: -1}})
	opt.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opt)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.TweetsSender

		err := cursor.Decode(&register)

		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}

	return results, true

}
