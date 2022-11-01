package routes

import (
	"github.com/ahmedabzk/go-microservices/controller"
	"github.com/gin-gonic/gin"
)

func ArticlesRouter(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/", controller.GetAllArticles())
	incommingRoutes.GET("articles/:article_id", controller.GetArticle())
}
