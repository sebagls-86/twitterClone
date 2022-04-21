package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func ConsultRelation(ctx *gin.Context) {

	ID := ctx.Query("id")

	if len(ID) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "We need the id parameter"})
		return
	}

	var t models.Relation

	t.UserID = IDUser
	t.UserRelationID = ID

	var resp models.ResponseConsultRelation

	status, err := bd.CheckRelation(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	ctx.JSON(http.StatusOK, resp)

}
