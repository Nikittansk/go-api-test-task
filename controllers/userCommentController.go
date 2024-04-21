package controllers

import (
	"net/http"

	"github.com/Nikittansk/go-api-test-task/database"
	"github.com/Nikittansk/go-api-test-task/models"
	"github.com/gin-gonic/gin"
)

func GetUserComments(ctx *gin.Context) {
	var comments []models.Comment

	id := ctx.Param("id")

	rows, err := database.DB.Query("SELECT id, text, user_id FROM comments WHERE user_id=$1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.Text, &comment.UserId); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		comments = append(comments, comment)
	}
	defer rows.Close()

	ctx.JSON(http.StatusOK, comments)
}

func InsertUserComment(ctx *gin.Context) {
	var (
		comment    models.Comment
		idResponse int
	)

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := database.DB.QueryRow("INSERT INTO comments (text, user_id) VALUES ($1, $2) RETURNING id", comment.Text, id).Scan(&idResponse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": idResponse,
	})
}
