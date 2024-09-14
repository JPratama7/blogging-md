package model

import (
	"github.com/rs/xid"
	"time"
)

type User struct {
	ID        xid.ID    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type BlogPost struct {
	ID        xid.ID    `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	AuthorID  xid.ID    `db:"author_id" json:"author_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Comment struct {
	ID         xid.ID    `db:"id" json:"id"`
	PostID     xid.ID    `db:"post_id" json:"post_id"`
	AuthorName string    `db:"author_name" json:"author_id"`
	Content    string    `db:"content" json:"content"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
