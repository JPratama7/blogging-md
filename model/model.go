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
	ID        xid.ID    `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	AuthorID  xid.ID    `db:"author_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Comment struct {
	ID        xid.ID    `db:"id"`
	PostID    xid.ID    `db:"post_id"`
	Author    string    `db:"author_name"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
