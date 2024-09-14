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

func PostRoute(route chi.Router) {
	route.Route("/posts", func(r chi.Router) {
		r.Get("/", controller.Fetch)
		r.Post("/", controller.Create)
		r.Get("/{id}", controller.GetPost)
		r.Put("/{id}", controller.Update)
		r.Delete("/{id}", controller.Delete)
		r.Get("/{id}/comments", controller.GetComment)
		r.Post("/{id}/comments", controller.CreateComment)
	})
}
