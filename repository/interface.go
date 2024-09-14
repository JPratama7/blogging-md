package repository

import (
	"blogging-md/model"
	"context"
)

type UserRepository interface {
	Create(c context.Context, user model.User) (model.User, error)
	FindByEmail(c context.Context, email string) (model.User, error)
	FindByID(c context.Context, id string) (model.User, error)
	FindByEmailAndPassword(c context.Context, email, passwordHashed string) (model.User, error)
}

type BlogPostRepository interface {
	Create(c context.Context, post model.BlogPost) (model.BlogPost, error)
	Fetch(c context.Context, page, size int) ([]model.BlogPost, error)
	FindByID(c context.Context, id string) (model.BlogPost, error)
	Update(c context.Context, post model.BlogPost) (model.BlogPost, error)
	Delete(c context.Context, id string) (model.BlogPost, error)
}

type CommentRepository interface {
	Create(c context.Context, comment model.Comment) (model.Comment, error)
	FindByIDPost(c context.Context, idPost string) (model.Comment, error)
}
