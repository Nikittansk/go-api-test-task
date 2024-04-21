package controllers

import (
	"net/http"
	"time"

	"github.com/Nikittansk/go-api-test-task/database"
	"github.com/Nikittansk/go-api-test-task/models"
	"github.com/gin-gonic/gin"
)

func GetComments(ctx *gin.Context) {
	var comments []models.Comment

	rows, err := database.DB.Query("SELECT id, text, user_id FROM comments")
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

func GetComment(ctx *gin.Context) {
	var comment models.Comment

	id := ctx.Param("id")

	if err := database.DB.QueryRow("SELECT id, text, user_id FROM comments WHERE id=$1", id).Scan(&comment.ID, &comment.Text, &comment.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func UpdateComment(ctx *gin.Context) {
	var comment models.Comment

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	updatedTime := time.Now().Format("2006-01-02 15:04:05.999999")

	_, err := database.DB.Exec("UPDATE comments SET text=$1, updated_at=$2 WHERE id=$3", comment.Text, updatedTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := database.DB.Exec("DELETE FROM comments WHERE id=$1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
