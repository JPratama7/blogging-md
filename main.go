package main

import (
	"aidanwoods.dev/go-paseto"
	"blogging-md/repository/blog"
	"blogging-md/repository/user"
	"blogging-md/router"
	"blogging-md/util"
	"github.com/JPratama7/util/token"
	"github.com/JPratama7/util/token/option"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	route := chi.NewRouter()

	db, err := sqlx.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(os.Getenv("SECRET_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	tokenizer := token.NewPaseto(
		publicKey,
		privateKey,
		option.WithExpiration(time.Hour),
		option.WithIssuer("blogging-md"),
		option.WithSubject("for-blogging-md"),
	)

	userRepo := user.New(db)
	blogPostRepo := blog.New(db)

	route.Use(middleware.AllowContentType("application/json"))
	route.Use(middleware.RequestID)
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)

	route.Use(util.InsertToRequest("userDb", userRepo))
	route.Use(util.InsertToRequest("tokenizer", tokenizer))
	route.Use(util.InsertToRequest("blogRepo", blogPostRepo))

	router.UserRoute(route)

	log.Fatal(http.ListenAndServe(":8080", route))
}
