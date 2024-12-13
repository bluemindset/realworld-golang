package comment

import (
	"database/sql"
	"fmt"
)

type CommentService struct {
	db *sql.DB
}

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{db}
}

func (service *CommentService) Create(comment *Comment) (*Comment, error) {
	var createdComment Comment

	fmt.Println(comment.ArticleId, comment.Body, comment.AuthorId)
	if err := service.db.QueryRow(`INSERT INTO comments (article_id, body, author_id) 
									VALUES ($1, $2, $3)
									RETURNING *;`, comment.ArticleId, comment.Body, comment.AuthorId).Scan(&createdComment.Id,
		&createdComment.ArticleId,
		&createdComment.Body,
		&createdComment.AuthorId); err != nil {
		return nil, fmt.Errorf("failed to create %v: %v", comment, err)
	}
	return &createdComment, nil
	// INSERT -> RETURN TO ROW -> COPY IT IN MY TYPE
}

func (service *CommentService) Update(comment *Comment) (*Comment, error) {
	var updatedComment *Comment
	if err := service.db.QueryRow(`UPDATE comments
									SET article_id  = $1,
									body = $2,
									author_id = $3,
									WHERE comment_id = $4 
									RETURNING id, article_id, body, author_id;`, comment.Id, comment.ArticleId, comment.Body, comment.AuthorId).Scan(updatedComment); err != nil {
		return nil, fmt.Errorf("failed to update %v: %v", comment, err)
	}
	return updatedComment, nil
}

func (service *CommentService) Delete(commentid int) (int, error) {
	if err := service.db.QueryRow(`DELETE FROM comments
									WHERE id = $1`, commentid); err != nil {
		return -1, fmt.Errorf("failed to create %d: %v", commentid, err)
	}
	return commentid, nil
}
