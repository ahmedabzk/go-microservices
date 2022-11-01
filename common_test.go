package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ahmedabzk/go-microservices/models"
	"github.com/gin-gonic/gin"
)

var tempArticleList []models.Article

func TestMain(m *testing.M) {
	// set the gin mode
	gin.SetMode(gin.TestMode)

	// run the other test
	os.Exit(m.Run())
}

// helper function to create a router during test
func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()

	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// helper function to process a request and test it is response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create a response recorder
	w := httptest.NewRecorder()

	// create the service and process the request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}

}

// this function is used to store the main lists for temporary one for test
func SaveList() {
	tempArticleList = models.ArticleList
}

// this function is used to restore the main list from the temporary one
func RestoreList() {
	models.ArticleList = tempArticleList
}
