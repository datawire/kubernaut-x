package broker

import (
	"fmt"
	"github.com/datawire/kubernaut/pkg/claimregistry"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

type Broker struct {
	router      *chi.Mux
	adminRouter *chi.Mux
	claims      *claimregistry.Registry
}

func NewBroker(claims claimregistry.Registry) *Broker {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	b := &Broker{
		claims: &claims,
		router: r,
	}

	b.configureRoutes()

	return b
}

func (b *Broker) Run(port int, adminPort int) error {
	fmt.Printf("Starting Broker... :%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), b.router)
}
