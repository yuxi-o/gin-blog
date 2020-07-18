package routers

import (
	_ "io"
	_ "os"

	"github.com/gin-gonic/gin"

	"blog/middleware/jwt"
	"blog/pkg/setting"
	"blog/routers/api"
	v1 "blog/routers/api/v1"

	_ "blog/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	//	gin.DisableConsoleColor()
	//	f, _ :=os.Create("gin.log")
	//	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)

		apiv1.POST("/tags", v1.AddTag)

		apiv1.PUT("/tags/:id", v1.EditTag)

		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)

		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.POST("/articles", v1.AddArticle)

		apiv1.PUT("/articles/:id", v1.EditArticle)

		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
