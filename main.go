package main

import (
	"aidanwoods.dev/go-paseto"
	"blogging-md/repository/blog"
	"blogging-md/repository/comment"
	"blogging-md/repository/user"
	"blogging-md/router"
	"blogging-md/util"
	"github.com/JPratama7/util/token"
	"github.com/JPratama7/util/token/option"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
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

	if err = db.Ping(); err != nil {
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
	commentsRepo := comment.New(db)

	route.Use(middleware.AllowContentType("application/json"))
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)
	route.Use(middleware.Timeout(5 * time.Second))

	route.Use(util.InsertToRequest("userDb", userRepo))
	route.Use(util.InsertToRequest[token.Token]("tokenizer", tokenizer))
	route.Use(util.InsertToRequest("postDb", blogPostRepo))
	route.Use(util.InsertToRequest("commentsDb", commentsRepo))

	router.UserRoute(route)
	router.PostRoute(route)

	log.Println("Server started on port " + os.Getenv("PORT"))
	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"), route))
}
