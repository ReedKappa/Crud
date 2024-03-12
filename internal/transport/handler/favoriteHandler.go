package handler

import (
	"crud/internal/core/interface/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type handlerFavorite struct {
	Login  string `json:"login"`
	BookId int    `json:"bookId"`
}

func AddFavorite(service service.FavoriteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var favorite handlerFavorite

		if err := c.BindJSON(&favorite); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "Неверное тело запроса"})
			return
		}

		err := service.AddFavorite(c.Request.Context(), favorite.Login, favorite.BookId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func GetFavorite(service service.FavoriteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.Param("login")

		if login == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "error"})
			return
		}

		favorite, err := service.GetFavorite(c.Request.Context(), login)

		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка получения книги по автору"})
			return
		}

		c.JSON(http.StatusOK, favorite)
	}
}
