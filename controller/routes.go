package controller

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
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