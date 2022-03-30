package routers

import (
	"projects/controller"
	authorController "projects/controller/author"
	bookController "projects/controller/book"
	"projects/controller/developer"
	"projects/controller/project"
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
	//books/list/paging?size=3&page=0&sort=-name
	bookGroup.GET("/list/paging", bookController.PagingV2)
	bookGroup.DELETE("/delete/:id", bookController.DeleteBookId)

	// Working with JWT and user management:
	authGroup := rest.Group("/author")
	authGroup.GET("/login", authorController.Login)
	authGroup.GET("/logout", authorController.Logout)
	authGroup.POST("/register", authorController.Register)
	authGroup.POST("/register/user/roles", authorController.AddRolesToUser)

	// Add reset password:

	// Working with JWT :
	tokenGroup := rest.Group("/token")
	tokenGroup.GET("/valid/:access_token", authorController.ValidToken)
	tokenGroup.GET("/expire/:access_token", authorController.ExpireToken)
	tokenGroup.GET("/renew/:refreshToken", authorController.RenewToken)

	// Access with JWT + Role Admin
	adminGroup := rest.Group("/admin")
	adminGroup.Use(utils.AdminProcess)
	adminGroup.GET("/list/book", bookController.ListBook)

	// Access with JWT + Role Viewer
	developerGroup := rest.Group("/developer")
	developerGroup.Use(utils.ViewerProcess)
	developerGroup.GET("/list", developer.ListDevelopers)
	developerGroup.POST("/add", developer.AddNewDeveloper)
	developerGroup.GET("/getbyid/:id", developer.GetDeveloperById)
	developerGroup.GET("/delete/:id", developer.DelDeveloperById)
	developerGroup.POST("/update", developer.UpdateDeveloper)

	///Test model mapper:
	developerGroup.POST("/copy", developer.Copy)

	projectGroup := rest.Group("/project")
	projectGroup.GET("/list", project.ListProjects)
	projectGroup.GET("/list/eager", project.ListEagerProjects)
	projectGroup.GET("/lazy/findbyid/:id", project.FindSimpleProjectById)
	projectGroup.GET("/eager/findbyid/:id", project.FindProjectById)

	// Test one to many in GoRM: User - Images
	// User group :
	userGroup := rest.Group("/users")
	userGroup.POST("/add", userController.AddUser)
	userGroup.GET("/list", userController.ListUser)
	userGroup.GET("/list/lazy", userController.ListLazyUser)
	userGroup.GET("/findbyid/:id", userController.FindByUserId)
	// Delete user and all user' images
	userGroup.GET("/delete/:id", userController.DeleteUserById)

	testGroup := rest.Group("/test")

	testGroup.GET("/copy", controller.Test)

	// Continue : Test - one to one

}
