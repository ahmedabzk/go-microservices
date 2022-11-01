package routes

import "github.com/gin-gonic/gin"

func ArticlesRouter(incommingRoutes *gin.Engine){
	incommingRoutes.GET("/", controller.GetAllArticles())
	incommingRoutes.GET("articles/:article_id", controller.GetArticle())
}
