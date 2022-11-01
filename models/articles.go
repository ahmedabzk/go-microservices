package models

type Article struct{
	ID			`json:"id"`
	Title		`json:"title"`
	Content		`json:"content"`
}

var ArticleList = []Article{
	Article{ID: 1, Title: "article one", Content: "here is the content"},
	Article{ID: 2, Title: "article two", Content: "content for article two"},
	Article{ID: 3, Title: "article three", Content: "content for article three"},
	Article{ID: 4, Title: "article four", Content: "content for article four"},
	Article{ID: 5, Title: "article five", Content: "content for article five"},
	Article{ID: 6, Title: "article six", Content: "content for article six"},
	Article{ID: 7, Title: "article seven", Content: "content for article seven"},
	Article{ID: 8, Title: "article eight", Content: "content for article eight"},
}