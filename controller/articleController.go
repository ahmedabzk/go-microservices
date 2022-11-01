package controller

import (
	"net/http"

	"github.com/ahmedabzk/go-microservices/models"
	"github.com/gin-gonic/gin"
)

func GetAllArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		articles := models.ReturnAllArticles()

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title":"articles",
				"payload": articles,
			},
		)
	}
}

func GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleId := c.Param("article_id")

		article, err := models.GetArticleById(articleId)

		if err != nil{
			c.HTML(
				http.StatusBadRequest, 
				"article.html",
				gin.H{"payload": err},
			)
		}
		c.HTML(
			http.StatusOK, 
			"article.html",
			gin.H{
				"article": article.Title,
				"payload": article,
			},
		)
	}
}
