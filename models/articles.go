package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var ArticleList = []Article{
	{ID: 1, Title: "article one", Content: "here is the content"},
	{ID: 2, Title: "article two", Content: "content for article two"},
	{ID: 3, Title: "article three", Content: "content for article three"},
	{ID: 4, Title: "article four", Content: "content for article four"},
	{ID: 5, Title: "article five", Content: "content for article five"},
	{ID: 6, Title: "article six", Content: "content for article six"},
	{ID: 7, Title: "article seven", Content: "content for article seven"},
	{ID: 8, Title: "article eight", Content: "content for article eight"},
}

func ReturnAllArticles() []Article {
	return ArticleList
}

func GetArticleById(id int) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}
