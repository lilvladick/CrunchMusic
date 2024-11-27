package postgres

import (
	"time"
)

type Author struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password,omitempty"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type News struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	NewsContent string    `json:"news_content"`
	AuthorID    int       `json:"author_id"`
	CategoryID  int       `json:"category_id"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewsComment struct {
	ID             int       `json:"id"`
	NewsID         int       `json:"news_id"`
	AuthorID       int       `json:"author_id"`
	CommentContent string    `json:"comment_content"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
