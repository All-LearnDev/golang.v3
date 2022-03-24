package routers

import (
	authorController "projects/controller/author"
	bookController "projects/controller/book"
	"projects/controller/developer"
	userController "projects/controller/user"
	"projects/utils"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func InitializeApiMapping(rest *echo.Echo) {
	// Enable cors in Echo:
	rest.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "http://localhost:1323/"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Book Group:
	bookGroup := rest.Group("/books")
	bookGroup.GET("/initialize", bookController.InitializeData)
	bookGroup.GET("/list", bookController.ListBook)
	bookGroup.POST("/validate", bookController.ValidateBook)
	bookGroup.POST("/add", bookController.AddBook)
	bookGroup.POST("/update", bookController.UpdateBook)
	bookGroup.GET("/findbyid/:id", bookController.FindById)
	bookGroup.GET("/list/:page/:pageSize", bookController.Paging)

	// User group :
	rest.POST("/users/add", userController.AddUser)
	rest.GET("/users/list", userController.ListUser)
	rest.GET("/users/findbyid/:id", userController.FindByUserId)
	///rest.GET("/users/download/excel", userController.WriteExcelFile)

	// Working with JWT and user management:
	authGroup := rest.Group("/author")
	authGroup.GET("/login", authorController.Login)
	authGroup.GET("/logout", authorController.Logout)
	authGroup.GET("/register", authorController.Register)

	// Working with JWT :
	tokenGroup := rest.Group("/token")
	tokenGroup.GET("/valid/:access_token", authorController.ValidToken)
	tokenGroup.GET("/expire/:access_token", authorController.ExpireToken)
	tokenGroup.GET("/renew/:refreshToken", authorController.RenewToken)

	// Test - Authentication with token:
	// Restricted Admin group - Test security :
	adminGroup := rest.Group("/admin")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &utils.JwtCustomClaims{},
		SigningKey: []byte("konmeo12397"),
	}
	adminGroup.Use(middleware.JWTWithConfig(config))
	adminGroup.GET("/list/book", bookController.ListBook)

	// Test many to many in GoRM:
	developerGroup := rest.Group("/developer")
	developerGroup.POST("/add", developer.AddNewDeveloper)
	developerGroup.GET("/getbyid/:id", developer.GetDeveloperById)
	developerGroup.GET("/list", developer.ListDevelopers)

}
