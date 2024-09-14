package comment

import (
	"blogging-md/model"
	"blogging-md/repository"
	"context"
	gsql "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

type commentRepoImpl struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) repository.CommentRepository {
	return &commentRepoImpl{
		db: db,
	}
}

func (cr commentRepoImpl) Create(c context.Context, comment model.Comment) (res model.Comment, err error) {
	res = comment
	res.ID = xid.New()

	query := gsql.NewStruct(&res).InsertInto("comments")
	query.SetFlavor(gsql.MySQL)

	q, args := query.Build()

	_, err = cr.db.ExecContext(c, q, args...)
	return
}

func (cr commentRepoImpl) FindByIDPost(c context.Context, idPost xid.ID) (res []model.Comment, err error) {
	query := gsql.NewSelectBuilder()
	query.Select("id", "content", "author_name", "created_at", "updated_at")
	query.From("comments")
	query.Where(query.Equal("id_post", idPost))
	query.SetFlavor(gsql.MySQL)

	q, args := query.Build()

	err = cr.db.SelectContext(c, &res, q, args...)

	return
}
