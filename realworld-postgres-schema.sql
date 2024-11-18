-- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    bio TEXT,
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- User Follows Table
CREATE TABLE user_follows (
    follower_id INTEGER REFERENCES users(id),
    followed_id INTEGER REFERENCES users(id),
    PRIMARY KEY (follower_id, followed_id),
    CHECK (follower_id != followed_id)
);

-- Articles Table
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    body TEXT NOT NULL,
    author_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Article Tags Table
CREATE TABLE article_tags (
    article_id INTEGER REFERENCES articles(id),
    tag_name VARCHAR(50),
    PRIMARY KEY (article_id, tag_name)
);

-- Favorites Table
CREATE TABLE article_favorites (
    user_id INTEGER REFERENCES users(id),
    article_id INTEGER REFERENCES articles(id),
    PRIMARY KEY (user_id, article_id)
);

-- Comments Table
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    body TEXT NOT NULL,
    article_id INTEGER REFERENCES articles(id),
    author_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_articles_author ON articles(author_id);
CREATE INDEX idx_comments_article ON comments(article_id);
CREATE INDEX idx_comments_author ON comments(author_id);
CREATE INDEX idx_article_tags_article ON article_tags(article_id);
CREATE INDEX idx_user_follows_follower ON user_follows(follower_id);
CREATE INDEX idx_user_follows_followed ON user_follows(followed_id);
CREATE INDEX idx_article_favorites_user ON article_favorites(user_id);
CREATE INDEX idx_article_favorites_article ON article_favorites(article_id);
