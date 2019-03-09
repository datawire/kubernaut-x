package broker

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
)

func (b *Broker) createClaim() http.HandlerFunc {
	return nil
}

func (b *Broker) deleteClaim() http.HandlerFunc {
	return nil
}

func (b *Broker) getClaim() http.HandlerFunc {
	return nil
}

func (b *Broker) listClaims() http.HandlerFunc {
	return nil
}

func (b *Broker) claimCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var claim *ClaimJSON
		var err error

		if claimName := chi.URLParam(r, "name"); claimName != "" {
			claim = &ClaimJSON{}
		} else {

		}

		if err != nil {
			if r.Method == "DELETE" {
				// OK: delete is idempotent.
			}
		}

		//if articleID := chi.URLParam(r, "articleID"); articleID != "" {
		//	article, err = dbGetArticle(articleID)
		//} else if articleSlug := chi.URLParam(r, "articleSlug"); articleSlug != "" {
		//	article, err = dbGetArticleBySlug(articleSlug)
		//} else {
		//	render.Render(w, r, ErrNotFound)
		//	return
		//}
		//if err != nil {
		//	render.Render(w, r, ErrNotFound)
		//	return
		//}

		ctx := context.WithValue(r.Context(), "claim", claim)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
