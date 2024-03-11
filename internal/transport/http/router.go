package http

import (
	"crud/internal/core/interface/service"
	"crud/internal/transport/handler"
	"crud/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(service service.AuthService, postService service.PostService, bookService service.BookService) *gin.Engine {
	router := gin.New()

	router.POST("/register", handler.RegisterUser(service))

	api := router.Group("/api", middleware.AuthMiddleware)
	{
		api.POST("/post", handler.CreatePost(postService))
		api.GET("/post/:id", handler.GetPost(postService))

		api.POST("/book/add", handler.AddBook(bookService))
		api.GET("/book/id/:id", handler.GetBook(bookService))
		api.GET("/book/name/:name", handler.GetBookByName(bookService))
		api.GET("/book/author/:author", handler.GetBooksByAuthor(bookService))
		api.GET("/book/genre/:genre", handler.GetBooksByGenre(bookService))
	}
	return router
}
