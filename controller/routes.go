package controller

import "net/http"

func RegisterRoutes() {
	const basePath = "/api/v1"

	http.HandleFunc(basePath+"/regions", GetRegionsHandler)
	http.HandleFunc(basePath+"/regions/id/", GetRegionByIDHandler)
	http.HandleFunc(basePath+"/regions/name/", GetRegionByNameHandler)
	http.HandleFunc(basePath+"/constellations", GetConstellationsHandler)
	http.HandleFunc(basePath+"/constellations/id/", GetConstellationByIDHandler)
	http.HandleFunc(basePath+"/constellations/name/", GetConstellationByNameHandler)
	http.HandleFunc(basePath+"/systems", GetSystemsHandler)
	http.HandleFunc(basePath+"/systems/id/", GetSystemByIDHandler)
	http.HandleFunc(basePath+"/systems/name/", GetSystemByNameHandler)
	http.HandleFunc(basePath+"/stargates", GetStargatesHandler)
	http.HandleFunc(basePath+"/stargates/id/", GetStargateBySystemIDHandler)
	http.HandleFunc(basePath+"/systems/spectral_class_counts", GetSpectralClassCountsHandler)
} 