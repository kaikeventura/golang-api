package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/golang-api/database"
	"github.com/kaikeventura/golang-api/model"
	"strconv"
)

func ShowBook(context *gin.Context) {
	strId := context.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Id has to be Integer",
		})

		return
	}

	db := database.GetDatabase()

	var book model.Book
	err = db.First(&book, id).Error
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot find book: " + err.Error(),
		})

		return
	}

	context.JSON(200, book)
}

func CreateBook(context *gin.Context) {
	db := database.GetDatabase()

	var book model.Book

	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})

		return
	}

	context.JSON(200, book)
}

func ShowBooks(context *gin.Context) {
	db := database.GetDatabase()

	var books []model.Book
	err := db.Find(&books).Error
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot list books: " + err.Error(),
		})

		return
	}

	context.JSON(200, books)
}

func UpdateBook(context *gin.Context) {
	db := database.GetDatabase()

	var book model.Book

	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot update book: " + err.Error(),
		})

		return
	}

	context.JSON(200, book)
}

func DeleteBook(context *gin.Context) {
	strId := context.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Id has to be Integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&model.Book{}, id).Error
	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot delete book: " + err.Error(),
		})

		return
	}

	context.Status(204)
}
