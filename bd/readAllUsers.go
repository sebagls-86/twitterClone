package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/sebagls-86/twitterClone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(ID string, page int64, search string, searchType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitterClone")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, _ = CheckRelation(r)

		if searchType == "new" && !found {
			include = true
		}
		if searchType == "follow" && found {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		if include {
			s.Password = ""
			s.Bio = ""
			s.Location = ""
			s.Banner = ""
			s.WebSite = ""
			s.Email = ""

			results = append(results, &s)
		}

	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)
	return results, true

}
