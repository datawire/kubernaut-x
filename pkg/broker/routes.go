package broker

import (
	"github.com/go-chi/chi"
)

func (b *Broker) configureRoutes() {
	b.router.Get("/healthz", b.healthz())

	b.router.Route("/claims", func(r chi.Router) {
		r.Get("/", b.listClaims())
		r.Route("/{name}", func(r chi.Router) {
			r.Use(claimCtx)
			r.Get("/", b.getClaim())
			r.Post("/", b.createClaim())
			r.Delete("/", b.deleteClaim())
		})
	})
}

func (b *Broker) configureAdminRoutes() {

}
