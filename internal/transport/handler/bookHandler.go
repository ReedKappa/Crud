package handler

import (
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerBook struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func AddBook(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book handlerBook

		if err := c.BindJSON(&book); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "Неверное тело запроса"})
			return
		}

		id, err := service.AddBook(c.Request.Context(), model.Book(book))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": id})
	}
}

func GetBook(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "Неверно передан id книги"})
			return
		}

		book, err := service.GetBook(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка получения книги"})
			return

		}
		c.JSON(http.StatusOK, handlerBook(book))
	}
}

func GetBookByName(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		if name == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "Не указано название книги"})
			return
		}

		book, err := service.GetBookByName(c.Request.Context(), name)

		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка получения книги по названию"})
			return
		}

		c.JSON(http.StatusOK, handlerBook(book))
	}
}

func GetBooksByAuthor(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		author := c.Param("author")

		if author == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "Не указан автор книги"})
			return
		}

		books, err := service.GetBooksByAuthor(c.Request.Context(), author)

		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка получения книги по автору"})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}

func GetBooksByGenre(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		genre := c.Param("genre")

		if genre == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "Не указан жанр книги"})
			return
		}

		books, err := service.GetBooksByGenre(c.Request.Context(), genre)

		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка получения книги по жанру"})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}
