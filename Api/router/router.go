package router

import (
	"log"
	"net/http"

	"api/constants"
	"api/infrastructure/handler"
	"api/infrastructure/middleware"
	"api/infrastructure/persistance"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetRoutes ルータのセットアップ
func SetRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	authMiddleware, err := middleware.SetAuth()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// Repository
	userRepo := persistance.NewUserRepository(db)
	projectRepo := persistance.NewProjectRepository(db)
	// Handler
	userHandle := handler.NewUserHandler(userRepo)
	projectHandle := handler.NewProjectHandler(projectRepo)

	// CORS対策
	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: middleware.Methods(),
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders:     middleware.Headers(),
		AllowCredentials: true,
		AllowOrigins: []string{
			// localhost
			constants.API_DEFAULT,
			//constants.API_ENDPOINT,
			//constants.API_ENDPOINT_LOGIN,
			//constants.API_ENDPOINT_SNS,
			//constants.API_ENDPOINT_INFOMATION,
			//constants.API_ENDPOINT_CREATOR,
		},
	}))

	v1 := router.Group(constants.API + constants.API_VERSION)
	{
		// AuthLogin
		v1.POST("/login", authMiddleware.LoginHandler)
		v1.POST("/logout", authMiddleware.LogoutHandler)

		// auth
		authGroup := v1.Group("/auth")
		authGroup.Use(authMiddleware.MiddlewareFunc())
		{
			authGroup.GET("/refresh_token", authMiddleware.RefreshHandler)
		}

		// UserAPI
		userGroup := v1.Group("/user")
		userGroup.POST("", userHandle.CreateUser) // 作成

		userGroup.Use(authMiddleware.MiddlewareFunc())
		{
			userGroup.DELETE(":id", userHandle.DeleteUser) // 削除
			userGroup.PUT(":id", userHandle.UpdateUser)    // 更新
			userGroup.GET("", userHandle.GetUsers)         // 全取得
			userGroup.GET(":id", userHandle.GetUser)       // 取得
		}

		me := v1.Group("/me")
		me.Use(authMiddleware.MiddlewareFunc())
		{
			me.GET("", userHandle.GetMe)
		}

		// ProjectAPI
		projectGroup := v1.Group("/project")
		projectGroup.Use(authMiddleware.MiddlewareFunc())
		{
			projectGroup.POST("", projectHandle.CreateProject)      // 作成
			projectGroup.DELETE(":id", projectHandle.DeleteProject) // 削除
			projectGroup.PUT(":id", projectHandle.UpdateProject)    // 更新
			projectGroup.GET("", projectHandle.GetProjects)         // 全取得
			projectGroup.GET(":id", projectHandle.GetProject)       // 取得
		}

		// TEST
		v1.GET("/", func(c *gin.Context) {
			// 200番だったら呼び出し
			c.JSON(http.StatusOK, gin.H{
				"message": "成功",
			})
		})
		// NoRoute
		router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			log.Printf("NoRoute claims: %#v\n", claims)
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		})
	}
	return router
}
