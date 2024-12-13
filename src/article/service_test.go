package article

import (
	"database/sql"
	"errors"
	"fmt"
	"realworld/app/test"
	"testing"
)

// Test driven development
// We first write tests and then code
// Scenarios => function, res

// When writing a db connection function
// 1) Driver by native go for mocking
// 2) Memory Aid
// CREATE

func TestCreateSuccess(t *testing.T) {
	//create new db connection
	//new article service
	//call create
	db := test.DbConnection(t)
	var article = &Article{
		AuthorId: 1,
		Title:    "Hey there!",
		Body:     "Hello World!",
	}

	test.SeedDbUgly(db, t)

	articleService := NewArticleService(db)
	newArticle, err := articleService.Create(article)

	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}
	if newArticle.AuthorId != article.AuthorId {
		t.Errorf("Expected to be equal: %v != %v", newArticle.AuthorId, article.AuthorId)
	}
	if newArticle.Title != article.Title {
		t.Errorf("Expected to be equal: %v != %v", newArticle.Title, article.Title)
	}
	if newArticle.Body != article.Body {
		t.Errorf("Expected to be equal: %v != %v", newArticle.Body, article.Body)
	}
}

func TestCreateError(t *testing.T) {

	db := test.DbConnection(t)

	test.SeedDbUgly(db, t)

	var article = &Article{
		AuthorId: 1,
		Title:    "Hey there!",
		Body:     "Hello World!",
	}
	var articleDummy = &Article{
		AuthorId: 2,
		Title:    "Hey there!!",
		Body:     "Hello World!!",
	}

	articleService := NewArticleService(db)
	newArticle, err := articleService.Create(article)

	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}
	if newArticle.AuthorId == articleDummy.AuthorId {
		t.Errorf("Expected to be equal: %v != %v", newArticle.AuthorId, article.AuthorId)
	}
	if newArticle.Title == articleDummy.Title {
		t.Errorf("Expected to be equal: %v != %v", newArticle.Title, article.Title)
	}
	if newArticle.Body == articleDummy.Body {
		t.Errorf("Expected to be equal: %v != %v", newArticle.Body, article.Body)
	}

}

// UPDATE
func TestUpdateSuccess(t *testing.T) {
	//create new db connection
	//new article service
	//call create
	db := test.DbConnection(t)
	var article = Article{
		Id:       1,
		AuthorId: 1,
		Title:    "Hey there!!",
		Body:     "Hello World!",
	}

	test.SeedDbUgly(db, t)

	articleService := NewArticleService(db)
	newArticle, err := articleService.Update(article)

	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}
	if newArticle.AuthorId != article.AuthorId {
		t.Errorf("Expected to be equal: %v != %v", newArticle.AuthorId, article.AuthorId)
	}
	if newArticle.Title != article.Title {
		t.Errorf("Expected to be equal: %v != %v", newArticle.Title, article.Title)
	}
	if newArticle.Body != article.Body {
		t.Errorf("Expected to be equal: %v != %v", newArticle.Body, article.Body)
	}
}

func TestUpdateError(t *testing.T) {

	db := test.DbConnection(t)

	test.SeedDbUgly(db, t)

	var article = Article{
		Id:       1,
		AuthorId: 1,
		Title:    "Hey there@!",
		Body:     "Hello World!",
	}
	var articleDummy = Article{
		AuthorId: 2,
		Title:    "Hey there!",
		Body:     "Hello World@!",
	}

	articleService := NewArticleService(db)
	newArticle, err := articleService.Update(article)

	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}
	if newArticle.AuthorId == articleDummy.AuthorId {
		t.Errorf("Expected to be equal: %v != %v", newArticle.AuthorId, article.AuthorId)
	}
	if newArticle.Title == articleDummy.Title {
		t.Errorf("Expected to be equal: %v != %v", newArticle.Title, article.Title)
	}
	if newArticle.Body == articleDummy.Body {
		t.Errorf("Expected to be equal: %v != %v", newArticle.Body, article.Body)
	}

}

func TestDeleteSuccess(t *testing.T) {

	db := test.DbConnection(t)

	test.SeedDbUgly(db, t)

	var article = Article{
		AuthorId: 1,
		Title:    "Hey there@!",
		Body:     "Hello World!",
	}

	articleService := NewArticleService(db)
	newArticle, err := articleService.Create(&article)
	fmt.Println("Created: ", newArticle.Id)
	articleId, err := articleService.Delete(newArticle.Id)

	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}

	_, err := articleService.FindById(articleId)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("No errors just empty, thus correct behaviour.")
	}

}
