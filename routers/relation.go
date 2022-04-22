package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func Relation(ctx *gin.Context) {

	ID := ctx.Query("id")

	if len(ID) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"We need the ID parameter": err.Error()})
		return
	}

	var t models.Relation

	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.InserRelation(t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Something went wrong in the relation": err.Error()})
		return
	}
	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"Couldn't insert relation": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Relation created"})

}
