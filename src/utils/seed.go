package utils

import (
	"database/sql"
	"fmt"
)

func SeedDbUgly(db *sql.DB) {

	// Insert a user
	var userID int
	err := db.QueryRow(`
        INSERT INTO users (username, email, password_hash)
        VALUES ($1, $2, $3)
        RETURNING id;
    `, "testuser", "testuser@example.com", "hashed_password").Scan(&userID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted user with ID:", userID)

	// Insert an article
	var articleID int
	err = db.QueryRow(`
        INSERT INTO articles (author_id, title, body)
        VALUES ($1, $2, $3)
        RETURNING id;
    `, userID, "Sample Article Title", "This is the body of the sample article.").Scan(&articleID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted article with ID:", articleID)

	// Insert a comment
	var commentID int
	err = db.QueryRow(`
        INSERT INTO comments (article_id, author_id, body)
        VALUES ($1, $2, $3)
        RETURNING id;
    `, articleID, userID, "This is a sample comment on the article.").Scan(&commentID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted comment with ID:", commentID)
}
