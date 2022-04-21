package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func EraseRelation(ctx *gin.Context) {

	ID := ctx.Query("id")

	if len(ID) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "We need the id parameter"})
		return
	}

	var t models.Relation

	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.DeleteRelation(t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong erasing the relation"})
		return
	}
	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't delete relation"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Relation erased"})

}
