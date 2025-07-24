package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	r.Route("/api/v1", func(r chi.Router) {
		// Redirect from /api/v1 to swagger UI
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
		})

		r.Get("/regions", GetRegionsHandler)
		r.Get("/regions/{regionID}", GetRegionByIDHandler)

		r.Get("/constellations", GetConstellationsHandler)
		r.Get("/constellations/{constellationID}", GetConstellationByIDHandler)
		r.Get("/regions/{regionID}/constellations", GetConstellationsByRegionIDHandler)

		r.Get("/systems", GetSystemsHandler)
		r.Get("/systems/{systemID}", GetSystemByIDHandler)
		r.Get("/constellations/{constellationID}/systems", GetSystemsByConstellationIDHandler)

		r.Get("/stargates", GetStargatesHandler)
		r.Get("/systems/{systemID}/stargates", GetStargateBySystemIDHandler)

		r.Get("/reports/spectral-class-counts", GetSpectralClassCountsHandler)
	})
} 