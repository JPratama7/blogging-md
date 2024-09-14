package controller

import (
	"blogging-md/model"
	"blogging-md/model/requests"
	"blogging-md/repository"
	"blogging-md/util"
	"github.com/JPratama7/util/token"
	"github.com/rs/xid"
	"net/http"
	"time"
)

func GetComment(w http.ResponseWriter, r *http.Request) {
	commentDb, err := util.ExtractFromRequest[repository.CommentRepository](r, "commentsDb")
	if err != nil {
		util.NewError().
			SetStatus("Failed To Extract Comment Repository").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	idPost := util.ExtractURL(r, "id")
	if idPost == "" {
		util.NewError().
			SetStatus("Failed To Parse ID").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}
	idPostParsed, err := xid.FromString(idPost)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Parse ID").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	comments, err := commentDb.FindByIDPost(r.Context(), idPostParsed)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Find Comments").
			SetCode(http.StatusNotFound).
			Write(w)
		return
	}

	util.NewResponse[[]model.Comment]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Fetched Comments").
		SetData(comments).
		Write(w)
	return
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
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

	commentDb, err := util.ExtractFromRequest[repository.CommentRepository](r, "commentsDb")
	if err != nil {
		util.NewError().
			SetStatus("Failed To Extract Comment Repository").
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

	userData, err := userDb.FindByID(r.Context(), idUser)
	if err != nil {
		util.NewError().
			SetStatus("Unauthorized").
			SetCode(http.StatusUnauthorized).
			Write(w)
		return
	}

	idPost := util.ExtractURL(r, "id")
	if idPost == "" {
		util.NewError().
			SetStatus("Failed To Parse ID").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}
	idPostParsed, err := xid.FromString(idPost)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Parse ID").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	_, err = postDb.FindByID(r.Context(), idPostParsed)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Find Post").
			SetCode(http.StatusNotFound).
			Write(w)
		return
	}

	req := new(requests.Comment)

	if err = util.Decode(r, req); err != nil {
		util.NewError().
			SetStatus("Failed To Decode Request").
			SetCode(http.StatusBadRequest).
			Write(w)
		return
	}

	comment := model.Comment{
		PostID:     idPostParsed,
		AuthorName: userData.Name,
		Content:    req.Content,
		CreatedAt:  time.Now(),
	}

	_, err = commentDb.Create(r.Context(), comment)
	if err != nil {
		util.NewError().
			SetStatus("Failed To Create Comment").
			SetCode(http.StatusInternalServerError).
			Write(w)
		return
	}

	util.NewResponse[model.Comment]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Created Comment").
		SetData(comment).
		Write(w)
	return

}
