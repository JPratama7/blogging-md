package blog

import (
	"blogging-md/model"
	"blogging-md/repository"
	"context"
	gsql "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

type postRepoImpl struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) repository.BlogPostRepository {
	return &postRepoImpl{
		db: db,
	}
}

func (p postRepoImpl) Create(c context.Context, post model.BlogPost) (res model.BlogPost, err error) {
	res = post
	res.ID = xid.New()

	query := gsql.NewStruct(&res).InsertInto("blog_posts", &res)
	query.SetFlavor(gsql.MySQL)

	q, args := query.Build()
	_, err = p.db.ExecContext(c, q, args...)
	return
}

func (p postRepoImpl) Fetch(c context.Context, page, size int) (res []model.BlogPost, err error) {
	query := gsql.NewSelectBuilder()
	query.Select("id", "title", "content", "author_id", "created_at", "updated_at")
	query.From("blog_posts")
	query.OrderBy("created_at DESC")
	query.Limit(size)
	query.Offset(page * size)
	query.SetFlavor(gsql.MySQL)

	q, args := query.Build()

	err = p.db.SelectContext(c, &res, q, args...)

	return
}

func (p postRepoImpl) FindByID(c context.Context, id xid.ID) (res model.BlogPost, err error) {
	query := gsql.NewSelectBuilder()
	query.Select("id", "title", "content", "author_id", "created_at", "updated_at")
	query.From("blog_posts")
	query.Where(query.Equal("id", id))
	query.SetFlavor(gsql.MySQL)

	q, args := query.Build()

	err = p.db.GetContext(c, &res, q, args...)
	return
}

func (p postRepoImpl) Update(c context.Context, post model.BlogPost) (res model.BlogPost, err error) {
	query := gsql.
		NewUpdateBuilder().
		Update("blog_posts")
	query.SetFlavor(gsql.MySQL)

	query.Set(
		query.Assign("title", post.Title),
		query.Assign("content", post.Content),
		query.Assign("author_id", post.AuthorID),
		query.Assign("updated_at", post.UpdatedAt),
	)

	query.Where(query.Equal("id", post.ID))

	q, args := query.Build()
	res = post
	_, err = p.db.ExecContext(c, q, args...)
	return
}

func (p postRepoImpl) Delete(c context.Context, id xid.ID) (res model.BlogPost, err error) {
	query := gsql.
		NewDeleteBuilder().
		DeleteFrom("blog_posts")
	query.
		Where(query.Equal("id", id)).
		SetFlavor(gsql.MySQL)

	q, args := query.Build()

	_, err = p.db.ExecContext(c, q, args...)
	return
}
