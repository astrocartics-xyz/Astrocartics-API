package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("https://api.astrocartics.xyz/swagger/doc.json"), //The url pointing to API definition
	))

<<<<<<< Updated upstream
	r.Route("api/v1", func(r chi.Router) {
		// Redirect from /v1 to swagger UI. We are on subdomain.
=======
	r.Route("/v1", func(r chi.Router) {
		// Redirect from /v1 to swagger UI
>>>>>>> Stashed changes
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
		r.Get("/regions/{regionID}/systems", GetSystemsByRegionIDHandler)
		r.Get("/constellations/{constellationID}/systems", GetSystemsByConstellationIDHandler)

		r.Get("/stargates", GetStargatesHandler)
		r.Get("/systems/{systemID}/stargates", GetStargateBySystemIDHandler)
		r.Get("/constellations/{constellationID}/stargates", GetStargateByConstellationIDHandler)
		r.Get("/regions/{regionID}/stargates", GetStargateByRegionIDHandler)

		r.Get("/planets", GetPlanetsHandler)
		r.Get("/planets/{planetID}", GetPlanetByIDHandler)
		r.Get("/systems/{systemID}/planets", GetPlanetsBySystemIDHandler)

		r.Get("/stations", GetStationsHandler)
		r.Get("/stations/{stationID}", GetStationByIDHandler)
		r.Get("/systems/{systemID}/stations", GetStationsBySystemIDHandler)

		r.Get("/regions/{regionID}/heatmap", GetSystemHeatmapByRegionHandler)
		r.Get("/systems/{systemID}/kills/summary", GetKillsBySystemIDHandler)
		r.Get("/constellations/{constellationID}/kills/summary", GetKillsByConstellationIDHandler)
		r.Get("/regions/{regionID}/kills/summary", GetKillsByRegionIDHandler)
		r.Get("/systems/{systemID}/killmails", GetRecentKillmailsBySystemIDHandler)

		r.Get("/rankings/regions/top", GetTopRegionsHandler)
		r.Get("/rankings/constellations/top", GetTopConstellationsHandler)
		r.Get("/rankings/systems/top", GetTopSystemsHandler)

		r.Get("/reports/spectral-class-counts", GetSpectralClassCountsHandler)
	})
<<<<<<< Updated upstream
} 
=======
}
>>>>>>> Stashed changes
