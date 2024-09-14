package router

import (
	"blogging-md/controller"
	"github.com/go-chi/chi/v5"
)

func UserRoute(route chi.Router) {
	route.Group(func(r chi.Router) {
		r.Post("/register", controller.Register)
		r.Post("/login", controller.Login)
	})

}
