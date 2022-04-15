package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/sebagls-86/twitterClone/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitterClone")
	col := db.Collection("relation")

	condition := bson.M{
		"userId":         t.UserID,
		"userRelationId": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil

}
