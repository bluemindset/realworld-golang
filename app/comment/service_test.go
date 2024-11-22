package comment

import (
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

func TestCreateSuccess(t *testing.T) {
	//create new db connection
	//new comment service
	//call create
	db := test.DbConnection(t)
	var comment = &Comment{
		AuthorId:  1,
		ArticleId: 1,
		Body:      "Hello World!",
	}

	commentService := NewCommentService(db)
	newComment, err := commentService.Create(comment)
	fmt.Println("helloooooooooooooooooooooooooooooo")
	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}
	if newComment.AuthorId != comment.AuthorId {
		t.Errorf("Expected to be equal: %v != %v", newComment.AuthorId, comment.AuthorId)
	}
	if newComment.ArticleId != comment.ArticleId {
		t.Errorf("Expected to be equal: %v != %v", newComment.ArticleId, comment.ArticleId)
	}
	if newComment.Body != comment.Body {
		t.Errorf("Expected to be equal: %v != %v", newComment.Body, comment.Body)
	}
}

func TestCreateError(t *testing.T) {

	db := test.DbConnection(t)

	var comment = &Comment{
		AuthorId:  1,
		ArticleId: 1,
		Body:      "Hello World",
	}
	var commentDummy = &Comment{
		AuthorId:  2,
		ArticleId: 2,
		Body:      "12341",
	}

	articleService := NewCommentService(db)
	newComment, err := articleService.Create(comment)

	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}
	if newComment.AuthorId == commentDummy.AuthorId {
		t.Errorf("Expected to be equal: %v != %v", newComment.AuthorId, comment.AuthorId)
	}
	if newComment.ArticleId == commentDummy.ArticleId {
		t.Errorf("Expected to be equal: %v != %v", newComment.ArticleId, comment.ArticleId)
	}
	if newComment.Body == commentDummy.Body {
		t.Errorf("Expected to be equal: %v != %v", newComment.Body, comment.Body)
	}

}
