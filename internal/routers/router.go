package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog-service/internal/middleware"
	"github.com/longjoy/blog-service/internal/routers/api"
	"github.com/longjoy/blog-service/internal/routers/api/v1"
	_ "github.com/longjoy/blog-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine  {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/auth",api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	upload := NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags",tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id",tag.Update)
		apiv1.PATCH("/tags/:id/state",tag.Update)
		apiv1.GET("/tags",tag.List)
		apiv1.GET("/tags/:id",tag.Get)

		apiv1.POST("/articles",article.Create)
		apiv1.DELETE("/articles/:id",article.Delete)
		apiv1.PUT("/articles/:id",article.Update)
		apiv1.PATCH("/articles/:id/state",article.Update)
		apiv1.GET("/articles/:id",article.Get)
		apiv1.GET("/articles",article.List)

	}
	return r
}
