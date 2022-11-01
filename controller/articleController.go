package controller

import (
	"net/http"
	"strconv"

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
				"title":   "articles",
				"payload": articles,
			},
		)
	}
}

func GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if articleId, err := strconv.Atoi(c.Param("article_id")); err == nil {
			// check if article exsits
			if article, err := models.GetArticleById(articleId); err == nil {
				c.HTML(
					http.StatusOK,
					"article.html",
					gin.H{
						"article": article.Title,
						"payload": article,
					},
				)
			} else {
				c.AbortWithError(http.StatusNotFound, err)
			}
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}

	}
}
