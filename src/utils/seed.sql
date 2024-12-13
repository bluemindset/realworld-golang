-- Insert a user
INSERT INTO users (username, email, password_hash)
VALUES ('testuser', 'testuser@example.com', 'hashed_password')
RETURNING id;

INSERT INTO users (username, email, password_hash)
VALUES ('testuser2', 'testuser@example.com', 'hashed_password')
RETURNING id;
-- Insert an article
INSERT INTO articles (author_id, title, body)
VALUES (user_id, 'Sample Article Title', 'This is the body of the sample article.')
RETURNING id;

-- Insert a comment
INSERT INTO comments (article_id, author_id, body)
VALUES (article_id, user_id, 'This is a sample comment on the article.')
RETURNING id;
