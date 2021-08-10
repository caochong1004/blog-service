package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/longjoy/blog-service/docs"
	"github.com/longjoy/blog-service/global"
	"github.com/longjoy/blog-service/internal/middleware"
	"github.com/longjoy/blog-service/internal/routers/api"
	"github.com/longjoy/blog-service/internal/routers/api/v1"
	"github.com/longjoy/blog-service/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key: "/auth",
		FillInterval: time.Second,
		Capacity: 10,
		Quantum: 10,
	})
func NewRouter() *gin.Engine  {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}else{
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	//获得请求超时时间配置
	r.Use(middleware.ContextTimeout(global.ServerSetting.RequestTimeout*time.Second))
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
