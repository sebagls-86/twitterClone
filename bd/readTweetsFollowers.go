package bd

import (
	"context"
	"time"

	"github.com/sebagls-86/twitterClone/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweetsFollowers(ID string, page int) ([]*models.TweetsFollowes, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterClone")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)

	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})

	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, conditions)
	var result []*models.TweetsFollowes
	err := cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true

}
