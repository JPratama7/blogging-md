package controller

import (
	"blogging-md/model"
	"blogging-md/model/requests"
	"blogging-md/model/responses"
	"blogging-md/repository"
	"blogging-md/util"
	"github.com/JPratama7/util/token"
	"log"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	userRepo, err := util.ExtractFromRequest[repository.UserRepository](r, "userDb")
	if err != nil {
		util.NewError().
			SetCode(http.StatusInternalServerError).
			SetStatus("Failed To Extract User Repository").
			Write(w)
		return
	}

	req := new(requests.Register)

	if err = util.Decode(r, req); err != nil {
		util.NewError().
			SetCode(http.StatusBadRequest).
			SetStatus("Failed To Decode Request").
			Write(w)
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		util.NewError().
			SetCode(http.StatusBadRequest).
			SetStatus("Name, Email And Password Are Required").
			Write(w)
		return
	}

	hash, err := util.HashPassword(req.Password)
	if err != nil {
		util.NewError().
			SetCode(http.StatusInternalServerError).
			SetStatus("Failed To Hash Password").
			Write(w)
		return
	}

	_, err = userRepo.FindByEmail(r.Context(), req.Email)
	if err == nil {
		util.NewError().
			SetCode(http.StatusBadRequest).
			SetStatus("Email Already Exists").
			Write(w)
		return
	}

	data, err := userRepo.Create(r.Context(), model.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Println(err)
		util.NewError().
			SetCode(http.StatusInternalServerError).
			SetStatus("Failed To Create User").
			Write(w)
		return
	}

	util.NewResponse[responses.Register]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Registered").
		SetData(responses.Register{
			ID:    data.ID.String(),
			Name:  data.Name,
			Email: data.Email,
		}).Write(w)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	userRepo, err := util.ExtractFromRequest[repository.UserRepository](r, "userDb")
	if err != nil {
		util.NewError().
			SetCode(http.StatusInternalServerError).
			SetStatus("Failed To Extract User Repository").
			Write(w)
		return
	}

	tokenizer, err := util.ExtractFromRequest[token.Token](r, "tokenizer")
	if err != nil {
		util.NewError().
			SetCode(http.StatusInternalServerError).
			SetStatus("Failed To Extract Token").
			Write(w)
		return
	}

	req := new(requests.Login)

	if err = util.Decode(r, req); err != nil {
		util.NewError().
			SetCode(http.StatusBadRequest).
			SetStatus("Failed To Decode Request").
			Write(w)
		return
	}

	if req.Email == "" || req.Password == "" {
		util.NewError().
			SetCode(http.StatusBadRequest).
			SetStatus("Email And Password Are Required").
			Write(w)
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		util.NewError().
			SetCode(http.StatusInternalServerError).
			SetStatus("Failed To Hash Password").
			Write(w)
		return
	}

	user, err := userRepo.FindByEmailAndPassword(r.Context(), req.Email, hashedPassword)
	if err != nil {
		util.NewError().
			SetCode(http.StatusInternalServerError).
			SetStatus("Failed To Find User").
			Write(w)
		return
	}

	tokenized := tokenizer.Encrypt(
		token.WithClaims("id", user.ID.String()),
	)

	util.NewResponse[responses.Login]().
		SetCode(http.StatusOK).
		SetStatus("Successfully Logged In").
		SetData(responses.Login{
			Token: tokenized,
		})
}
