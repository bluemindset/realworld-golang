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

func (service *ArticleService) Update(article Article) (Article, error) {
	var updatedArticle Article
	if err := service.db.QueryRow(`UPDATE articles
									SET title = $1,
									body = $2
									WHERE id = $3 
									RETURNING  title, body, author_id,id;`, article.Title, article.Body, article.Id).Scan(&updatedArticle.Title, &updatedArticle.Body,
		&updatedArticle.AuthorId, &updatedArticle.Id); err != nil {
		return Article{}, fmt.Errorf("failed to update %v: %v", article, err)
	}
	return updatedArticle, nil
}

func (service *ArticleService) Delete(articleid int) (int, error) {
	if err := service.db.QueryRow(`DELETE FROM articles
									WHERE id = $1
									RETURNING id;`, articleid).Scan(&articleid); err != nil {
		return -1, fmt.Errorf("failed to delete %d: %v", articleid, err)
	}
	return articleid, nil
}

func (service *ArticleService) Find() ([]Article, error) {
	var allArticles []Article
	rows, err := service.db.Query(`SELECT title, body, author_id FROM ARTICLES;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var article Article

		if err := rows.Scan(&article.Title, &article.Body, &article.AuthorId); err != nil {
			return []Article{}, err
		}
		allArticles = append(allArticles, article)
	}
	if err = rows.Err(); err != nil {
		return []Article{}, err
	}
	return allArticles, nil
}

func (service *ArticleService) FindById(articleId int) (Article, error) {
	var article Article
	if err := service.db.QueryRow(`SELECT  title, body, author_id FROM ARTICLES WHERE id=$1;`, articleId).Scan(&article.Title, &article.Body, &article.AuthorId); err != nil {
		return Article{}, err
	}
	return article, nil
}
