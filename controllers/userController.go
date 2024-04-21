package controllers

import (
	"net/http"
	"time"

	"github.com/Nikittansk/go-api-test-task/database"
	"github.com/Nikittansk/go-api-test-task/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	var users []models.User

	rows, err := database.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		users = append(users, user)
	}
	defer rows.Close()

	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	var user models.User

	id := ctx.Param("id")

	if err := database.DB.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func InsertUser(ctx *gin.Context) {
	var (
		user models.User
		id   int
	)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := database.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&id); err != nil {
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

func UpdateUser(ctx *gin.Context) {
	var user models.User

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	updatedTime := time.Now().Format("2006-01-02 15:04:05.999999")

	_, err := database.DB.Exec("UPDATE users SET name=$1, email=$2, updated_at=$3 WHERE id=$4", user.Name, user.Email, updatedTime, id)
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

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := database.DB.Exec("DELETE FROM users WHERE id=$1", id)
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
