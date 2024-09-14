package user

import (
	"blogging-md/model"
	"blogging-md/repository"
	"context"
	gsql "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

type userRepoImpl struct {
	db *sqlx.DB
}

func (u *userRepoImpl) Create(c context.Context, user model.User) (res model.User, err error) {
	res = user
	res.ID = xid.New()

	query := gsql.NewStruct(&res).InsertInto("users", &res)
	query.SetFlavor(gsql.MySQL)

	q, args := query.Build()

	_, err = u.db.Exec(q, args...)
	return
}

func (u *userRepoImpl) FindByEmail(c context.Context, email string) (res model.User, err error) {
	query := gsql.NewSelectBuilder()
	query = query.Select("id", "name", "email", "password_hash", "created_at", "updated_at")
	query = query.From("users")
	query.SetFlavor(gsql.MySQL)
	query = query.Where(query.Equal("email", email))

	q, args := query.Build()

	err = u.db.GetContext(c, &res, q, args...)
	return
}

func (u *userRepoImpl) FindByID(c context.Context, id string) (res model.User, err error) {
	query := gsql.NewSelectBuilder()
	query.Select("id", "name", "email", "password_hash", "created_at", "updated_at")
	query.From("users")
	query.Where(query.Equal("id", id))
	query.SetFlavor(gsql.MySQL)

	q, args := query.Build()

	err = u.db.GetContext(c, &res, q, args...)
	return
}

func (u *userRepoImpl) FindByEmailAndPassword(c context.Context, email, passwordHashed string) (res model.User, err error) {
	query := gsql.NewSelectBuilder()
	query.Select("id", "name", "email", "password_hash", "created_at", "updated_at")
	query.From("users")
	query.Where(query.Equal("email", email))
	query.And(query.Equal("password_hash", passwordHashed))
	query.SetFlavor(gsql.MySQL)

	q, args := query.Build()

	err = u.db.GetContext(c, &res, q, args...)
	return
}

func New(db *sqlx.DB) repository.UserRepository {
	return &userRepoImpl{db}
}
