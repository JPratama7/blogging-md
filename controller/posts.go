package controller

import (
	"blogging-md/model"
	"blogging-md/model/requests"
	"blogging-md/repository"
	"blogging-md/util"
	"github.com/JPratama7/util/token"
	"github.com/rs/xid"
	"net/http"
	"strconv"
	"time"
)

func Fetch(w http.ResponseWriter, r *http.Request) {
	bearer, err := util.ExtractBearer(r)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	tokenizer, err := util.ExtractFromRequest[token.Token](r, "tokenizer")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	userDb, err := util.ExtractFromRequest[repository.UserRepository](r, "userDb")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	postDb, err := util.ExtractFromRequest[repository.BlogPostRepository](r, "postDb")
	if err != nil {
		util.NewError().
			SetStatus("Failed To Extract Post Repository").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	dec, err := tokenizer.Decrypt(bearer)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	idUser := xid.NilID()
	err = dec.Get("id", &idUser)
	if err != nil || idUser.IsNil() {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	_, err = userDb.FindByID(r.Context(), idUser)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	page := util.ExtractQuery(r, "page", "1")
	size := util.ExtractQuery(r, "page_size", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Parse Page").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Parse Page Size").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	posts, err := postDb.Fetch(r.Context(), pageInt, sizeInt)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Fetch Posts").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	util.NewResponse[[]model.BlogPost]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Fetched Posts").
		SetData(posts).
		Write(w)
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	bearer, err := util.ExtractBearer(r)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	tokenizer, err := util.ExtractFromRequest[token.Token](r, "tokenizer")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	userDb, err := util.ExtractFromRequest[repository.UserRepository](r, "userDb")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	postDb, err := util.ExtractFromRequest[repository.BlogPostRepository](r, "postDb")
	if err != nil {
		util.NewError().
			SetStatus("Failed To Extract Post Repository").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	dec, err := tokenizer.Decrypt(bearer)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	idUser := xid.NilID()
	err = dec.Get("id", &idUser)
	if err != nil || idUser.IsNil() {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	_, err = userDb.FindByID(r.Context(), idUser)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	req := new(requests.Post)

	if err = util.Decode(r, req); err != nil {
		util.NewError().
			SetStatus("Failed To Decode Request").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	if req.Title == "" || req.Content == "" {
		util.NewError().
			SetStatus("Title And Content Are Required").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	post, err := postDb.Create(r.Context(), model.BlogPost{
		Title:     req.Title,
		Content:   req.Content,
		AuthorID:  idUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		util.NewError().
			SetStatus("Failed To Create Post").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	util.NewResponse[model.BlogPost]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Created Post").
		SetData(post).
		Write(w)

	return
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	bearer, err := util.ExtractBearer(r)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	tokenizer, err := util.ExtractFromRequest[token.Token](r, "tokenizer")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	userDb, err := util.ExtractFromRequest[repository.UserRepository](r, "userDb")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	postDb, err := util.ExtractFromRequest[repository.BlogPostRepository](r, "postDb")
	if err != nil {
		util.NewError().
			SetStatus("Failed To Extract Post Repository").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	dec, err := tokenizer.Decrypt(bearer)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	idUser := xid.NilID()
	err = dec.Get("id", &idUser)
	if err != nil || idUser.IsNil() {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	_, err = userDb.FindByID(r.Context(), idUser)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	id := util.ExtractURL(r, "id", "")
	idPost, err := xid.FromString(id)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Parse ID").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	post, err := postDb.FindByID(r.Context(), idPost)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Find Post").
			SetCode(http.StatusNotFound).
			Write(w)
		return
	}

	util.NewResponse[model.BlogPost]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Fetched Post").
		SetData(post).
		Write(w)
	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	bearer, err := util.ExtractBearer(r)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	tokenizer, err := util.ExtractFromRequest[token.Token](r, "tokenizer")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	userDb, err := util.ExtractFromRequest[repository.UserRepository](r, "userDb")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	postDb, err := util.ExtractFromRequest[repository.BlogPostRepository](r, "postDb")
	if err != nil {
		util.NewError().
			SetStatus("Failed To Extract Post Repository").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	dec, err := tokenizer.Decrypt(bearer)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	idUser := xid.NilID()
	err = dec.Get("id", &idUser)
	if err != nil || idUser.IsNil() {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	_, err = userDb.FindByID(r.Context(), idUser)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	id := util.ExtractURL(r, "id", "")
	idPost, err := xid.FromString(id)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Parse ID").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	post, err := postDb.FindByID(r.Context(), idPost)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Find Post").
			SetCode(http.StatusNotFound).
			Write(w)
		return
	}

	if post.AuthorID != idUser {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	req := new(requests.Post)

	if err = util.Decode(r, req); err != nil {
		util.NewError().
			SetStatus("Failed To Decode Request").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	post.Title = req.Title
	post.Content = req.Content
	post.UpdatedAt = time.Now()

	_, err = postDb.Update(r.Context(), post)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Update Post").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	util.NewResponse[model.BlogPost]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Updated Post").
		SetData(post).
		Write(w)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	bearer, err := util.ExtractBearer(r)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	tokenizer, err := util.ExtractFromRequest[token.Token](r, "tokenizer")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	userDb, err := util.ExtractFromRequest[repository.UserRepository](r, "userDb")
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	postDb, err := util.ExtractFromRequest[repository.BlogPostRepository](r, "postDb")
	if err != nil {
		util.NewError().
			SetStatus("Failed To Extract Post Repository").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	dec, err := tokenizer.Decrypt(bearer)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	idUser := xid.NilID()
	err = dec.Get("id", &idUser)
	if err != nil || idUser.IsNil() {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	_, err = userDb.FindByID(r.Context(), idUser)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	id := util.ExtractURL(r, "id", "")
	idPost, err := xid.FromString(id)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Parse ID").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	post, err := postDb.FindByID(r.Context(), idPost)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Find Post").
			SetCode(http.StatusNotFound).
			Write(w)
		return
	}

	if post.AuthorID != idUser {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	_, err = postDb.Delete(r.Context(), idPost)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Delete Post").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	util.NewResponse[model.BlogPost]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Deleted Post").
		SetData(post).
		Write(w)
	return
}
