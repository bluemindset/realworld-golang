package article

import (
	"database/sql"
	"fmt"
)

type ArticleService struct {
	db *sql.DB
}

func NewArticleService(db *sql.DB) *ArticleService {
	return &ArticleService{db}
}

func (service *ArticleService) Create(article *Article) (*Article, error) {
	var createdArticle Article

	fmt.Println(article.Title, article.Body, article.AuthorId)
	if err := service.db.QueryRow(`INSERT INTO articles (author_id, title, body) 
									VALUES ($1, $2, $3)
									RETURNING *;`, article.AuthorId, article.Title, article.Body).Scan(&createdArticle.Id,
		&createdArticle.Title,
		&createdArticle.Body,
		&createdArticle.AuthorId); err != nil {
		return nil, fmt.Errorf("failed to create %v: %v", article, err)
	}
	return &createdArticle, nil
	// INSERT -> RETURN TO ROW -> COPY IT IN MY TYPE
}

func (service *ArticleService) Update(article *Article) (*Article, error) {
	var updatedArticle *Article
	if err := service.db.QueryRow(`UPDATE articles
									SET articlename = ?,
									title = ?,
									body = ?,
									WHERE author_id = ? 
									RETURNING id, title, body, author_id;`, article.Title, article.Body, article.AuthorId, article.Id).Scan(updatedArticle); err != nil {
		return nil, fmt.Errorf("failed to update %v: %v", article, err)
	}
	return updatedArticle, nil
}

func (service *ArticleService) Delete(articleid int) (int, error) {
	if err := service.db.QueryRow(`DELETE FROM articles
									WHERE id = ?`, articleid); err != nil {
		return -1, fmt.Errorf("failed to create %d: %v", articleid, err)
	}
	return articleid, nil
}
