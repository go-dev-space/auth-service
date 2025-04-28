package interfaces

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewAuthRouter is a constructor funktion for creating
// a new bounded context subrouter
func NewAuthRouter(hc HealthcheckHandler, rh RegistrationHandler, accessHeader func(http.Handler) http.Handler) http.Handler {
	r := chi.NewRouter()
	// healtcheck route with access header middlware check
	r.With(accessHeader).Get("/healthcheck", hc.Handle)
	// auth user route
	r.Route("/user", func(r chi.Router) {
		// create route
		r.Post("/create", rh.Handle)
	})

	return r
}
