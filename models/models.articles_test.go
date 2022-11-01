package models

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := ReturnAllArticles()

	// check that the list of articles returned is same as the length of the variable holding the list
	if len(alist) != len(ArticleList){
		t.Fail()
	}

	// check that each member is identicle

	for i, v := range alist{
		if v.ID != ArticleList[i].ID || v.Title != ArticleList[i].Title || v.Content != ArticleList[i].Content{
			t.Fail()
			break
		}
	}
}
