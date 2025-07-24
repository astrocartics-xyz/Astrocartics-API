package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/astrocartics-xyz/Astrocartics-API/service"
	"github.com/go-chi/chi/v5"
)

// respondJSON is a helper to send JSON responses.
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshalling JSON: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// respondError is a helper to send error responses.
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// GetRegionsHandler godoc
// @Summary Get regions
// @Description Get all regions, or search for a region by name
// @Tags regions
// @Accept  json
// @Produce  json
// @Param name query string false "Region name to search for"
// @Success 200 {array} models.Region
// @Router /regions [get]
func GetRegionsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		region, err := service.GetRegionByName(name)
		if err != nil {
			log.Printf("Error fetching region by name '%s': %v", name, err)
			respondError(w, http.StatusInternalServerError, "Failed to retrieve region")
			return
		}
		if region == nil {
			respondError(w, http.StatusNotFound, fmt.Sprintf("Region '%s' not found", name))
			return
		}
		respondJSON(w, http.StatusOK, region)
		return
	}

	regions, err := service.GetAllRegions()
	if err != nil {
		log.Printf("Error fetching regions: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve regions")
		return
	}
	respondJSON(w, http.StatusOK, regions)
}

// GetRegionByIDHandler godoc
// @Summary Get a region by ID
// @Description Get a single region by its unique ID
// @Tags regions
// @Accept  json
// @Produce  json
// @Param regionID path int true "Region ID"
// @Success 200 {object} models.Region
// @Router /regions/{regionID} [get]
func GetRegionByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "regionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid region ID")
		return
	}

	region, err := service.GetRegionByID(id)
	if err != nil {
		log.Printf("Error fetching region %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve region")
		return
	}
	if region == nil {
		respondError(w, http.StatusNotFound, "Region not found")
		return
	}
	respondJSON(w, http.StatusOK, region)
}

// GetConstellationsHandler godoc
// @Summary Get constellations
// @Description Get all constellations, or search for a constellation by name
// @Tags constellations
// @Accept  json
// @Produce  json
// @Param name query string false "Constellation name to search for"
// @Success 200 {array} models.Constellation
// @Router /constellations [get]
func GetConstellationsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		constellation, err := service.GetConstellationByName(name)
		if err != nil {
			log.Printf("Error fetching constellation by name '%s': %v", name, err)
			respondError(w, http.StatusInternalServerError, "Failed to retrieve constellation")
			return
		}
		if constellation == nil {
			respondError(w, http.StatusNotFound, fmt.Sprintf("Constellation '%s' not found", name))
			return
		}
		respondJSON(w, http.StatusOK, constellation)
		return
	}

	constellations, err := service.GetAllConstellations()
	if err != nil {
		log.Printf("Error fetching constellations: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve constellations")
		return
	}
	respondJSON(w, http.StatusOK, constellations)
}

// GetConstellationByIDHandler godoc
// @Summary Get a constellation by ID
// @Description Get a single constellation by its unique ID
// @Tags constellations
// @Accept  json
// @Produce  json
// @Param constellationID path int true "Constellation ID"
// @Success 200 {array} models.Constellation
// @Router /constellations/{constellationID} [get]
func GetConstellationByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "constellationID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid constellation ID")
		return
	}

	constellation, err := service.GetConstellationByIDOrRegionID(id)
	if err != nil {
		log.Printf("Error fetching constellation %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve constellation")
		return
	}
	if constellation == nil {
		respondError(w, http.StatusNotFound, "Constellation not found")
		return
	}
	respondJSON(w, http.StatusOK, constellation)
}

// GetConstellationsByRegionIDHandler godoc
// @Summary Get constellations by region ID
// @Description Get all constellations for a specific region
// @Tags constellations
// @Accept  json
// @Produce  json
// @Param regionID path int true "Region ID"
// @Success 200 {array} models.Constellation
// @Router /regions/{regionID}/constellations [get]
func GetConstellationsByRegionIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "regionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid region ID")
		return
	}

	constellations, err := service.GetConstellationByIDOrRegionID(id)
	if err != nil {
		log.Printf("Error fetching constellations for region %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve constellations")
		return
	}
	if constellations == nil {
		respondError(w, http.StatusNotFound, "No constellations found for this region")
		return
	}
	respondJSON(w, http.StatusOK, constellations)
}

// GetSystemsHandler godoc
// @Summary Get systems
// @Description Get all systems, or search for a system by name
// @Tags systems
// @Accept  json
// @Produce  json
// @Param name query string false "System name to search for"
// @Success 200 {array} models.System
// @Router /systems [get]
func GetSystemsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		system, err := service.GetSystemByName(name)
		if err != nil {
			log.Printf("Error fetching system by name '%s': %v", name, err)
			respondError(w, http.StatusInternalServerError, "Failed to retrieve system")
			return
		}
		if system == nil {
			respondError(w, http.StatusNotFound, fmt.Sprintf("System '%s' not found", name))
			return
		}
		respondJSON(w, http.StatusOK, system)
		return
	}

	systems, err := service.GetAllSystems()
	if err != nil {
		log.Printf("Error fetching systems: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve systems")
		return
	}
	respondJSON(w, http.StatusOK, systems)
}

// GetSystemByIDHandler godoc
// @Summary Get a system by ID
// @Description Get a single system by its unique ID
// @Tags systems
// @Accept  json
// @Produce  json
// @Param systemID path int true "System ID"
// @Success 200 {array} models.System
// @Router /systems/{systemID} [get]
func GetSystemByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "systemID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid system ID")
		return
	}

	system, err := service.GetSystemByIDOrConstellationID(id)
	if err != nil {
		log.Printf("Error fetching system %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve system")
		return
	}
	if system == nil {
		respondError(w, http.StatusNotFound, "System not found")
		return
	}
	respondJSON(w, http.StatusOK, system)
}

// GetSystemsByRegionIDHandler godoc
// @Summary Get systems by region ID
// @Description Get all systems for a specific region
// @Tags systems
// @Accept  json
// @Produce  json
// @Param regionID path int true "Region ID"
// @Success 200 {array} models.System
// @Router /regions/{regionID}/systems [get]
func GetSystemsByRegionIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "regionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid region ID")
		return
	}

	systems, err := service.GetSystemsByRegionID(id)
	if err != nil {
		log.Printf("Error fetching systems for region %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve systems")
		return
	}
	if systems == nil {
		respondError(w, http.StatusNotFound, "No systems found for this region")
		return
	}
	respondJSON(w, http.StatusOK, systems)
}

// GetSystemsByConstellationIDHandler godoc
// @Summary Get systems by constellation ID
// @Description Get all systems for a specific constellation
// @Tags systems
// @Accept  json
// @Produce  json
// @Param constellationID path int true "Constellation ID"
// @Success 200 {array} models.System
// @Router /constellations/{constellationID}/systems [get]
func GetSystemsByConstellationIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "constellationID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid constellation ID")
		return
	}

	systems, err := service.GetSystemByIDOrConstellationID(id)
	if err != nil {
		log.Printf("Error fetching systems for constellation %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve systems")
		return
	}
	if systems == nil {
		respondError(w, http.StatusNotFound, "No systems found for this constellation")
		return
	}
	respondJSON(w, http.StatusOK, systems)
}

// GetStargatesHandler godoc
// @Summary Get stargates
// @Description Get all stargates
// @Tags stargates
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Stargate
// @Router /stargates [get]
func GetStargatesHandler(w http.ResponseWriter, r *http.Request) {
	stargates, err := service.GetAllStargates()
	if err != nil {
		log.Printf("Error fetching stargates: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve stargates")
		return
	}
	respondJSON(w, http.StatusOK, stargates)
}

// GetStargateBySystemIDHandler godoc
// @Summary Get stargates by system ID
// @Description Get all stargates for a specific system
// @Tags stargates
// @Accept  json
// @Produce  json
// @Param systemID path int true "System ID"
// @Success 200 {array} models.Stargate
// @Router /systems/{systemID}/stargates [get]
func GetStargateBySystemIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "systemID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid system ID")
		return
	}

	stargates, err := service.GetStargateBySystemID(id)
	if err != nil {
		log.Printf("Error fetching stargates for system %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve stargates")
		return
	}
	if len(stargates) == 0 {
		respondError(w, http.StatusNotFound, "No stargates found for this system")
		return
	}
	respondJSON(w, http.StatusOK, stargates)
}

// GetSpectralClassCountsHandler godoc
// @Summary Get spectral class counts
// @Description Get a report of system counts by spectral class
// @Tags reports
// @Accept  json
// @Produce  json
// @Success 200 {array} models.SpectralClassCount
// @Router /reports/spectral-class-counts [get]
func GetSpectralClassCountsHandler(w http.ResponseWriter, r *http.Request) {
	counts, err := service.GetSpectralClassCounts()
	if err != nil {
		log.Printf("Error fetching spectral class counts: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve spectral class counts")
		return
	}
	respondJSON(w, http.StatusOK, counts)
}

// GetPlanetsHandler godoc
// @Summary Get planets
// @Description Get all planets, or search for a planet by name
// @Tags planets
// @Accept  json
// @Produce  json
// @Param name query string false "Planet name to search for"
// @Success 200 {array} models.Planet
// @Router /planets [get]
func GetPlanetsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		planet, err := service.GetPlanetByName(name)
		if err != nil {
			log.Printf("Error fetching planet by name '%s': %v", name, err)
			respondError(w, http.StatusInternalServerError, "Failed to retrieve planet")
			return
		}
		if planet == nil {
			respondError(w, http.StatusNotFound, fmt.Sprintf("Planet '%s' not found", name))
			return
		}
		respondJSON(w, http.StatusOK, planet)
		return
	}

	planets, err := service.GetAllPlanets()
	if err != nil {
		log.Printf("Error fetching planets: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve planets")
		return
	}
	respondJSON(w, http.StatusOK, planets)
}

// GetPlanetByIDHandler godoc
// @Summary Get a planet by ID
// @Description Get a single planet by its unique ID
// @Tags planets
// @Accept  json
// @Produce  json
// @Param planetID path int true "Planet ID"
// @Success 200 {object} models.Planet
// @Router /planets/{planetID} [get]
func GetPlanetByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "planetID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid planet ID")
		return
	}

	planet, err := service.GetPlanetByID(id)
	if err != nil {
		log.Printf("Error fetching planet %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve planet")
		return
	}
	if planet == nil {
		respondError(w, http.StatusNotFound, "Planet not found")
		return
	}
	respondJSON(w, http.StatusOK, planet)
}

// GetPlanetsBySystemIDHandler godoc
// @Summary Get planets by system ID
// @Description Get all planets for a specific system
// @Tags planets
// @Accept  json
// @Produce  json
// @Param systemID path int true "System ID"
// @Success 200 {array} models.Planet
// @Router /systems/{systemID}/planets [get]
func GetPlanetsBySystemIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "systemID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid system ID")
		return
	}

	planets, err := service.GetPlanetsBySystemID(id)
	if err != nil {
		log.Printf("Error fetching planets for system %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve planets")
		return
	}
	if planets == nil {
		respondError(w, http.StatusNotFound, "No planets found for this system")
		return
	}
	respondJSON(w, http.StatusOK, planets)
}

// GetStationsHandler godoc
// @Summary Get stations
// @Description Get all stations, or search for a station by name
// @Tags stations
// @Accept  json
// @Produce  json
// @Param name query string false "Station name to search for"
// @Success 200 {array} models.Station
// @Router /stations [get]
func GetStationsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		station, err := service.GetStationByName(name)
		if err != nil {
			log.Printf("Error fetching station by name '%s': %v", name, err)
			respondError(w, http.StatusInternalServerError, "Failed to retrieve station")
			return
		}
		if station == nil {
			respondError(w, http.StatusNotFound, fmt.Sprintf("Station '%s' not found", name))
			return
		}
		respondJSON(w, http.StatusOK, station)
		return
	}

	stations, err := service.GetAllStations()
	if err != nil {
		log.Printf("Error fetching stations: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve stations")
		return
	}
	respondJSON(w, http.StatusOK, stations)
}

// GetStationByIDHandler godoc
// @Summary Get a station by ID
// @Description Get a single station by its unique ID
// @Tags stations
// @Accept  json
// @Produce  json
// @Param stationID path int true "Station ID"
// @Success 200 {object} models.Station
// @Router /stations/{stationID} [get]
func GetStationByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "stationID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid station ID")
		return
	}

	station, err := service.GetStationByID(id)
	if err != nil {
		log.Printf("Error fetching station %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve station")
		return
	}
	if station == nil {
		respondError(w, http.StatusNotFound, "Station not found")
		return
	}
	respondJSON(w, http.StatusOK, station)
}

// GetStationsBySystemIDHandler godoc
// @Summary Get stations by system ID
// @Description Get all stations for a specific system
// @Tags stations
// @Accept  json
// @Produce  json
// @Param systemID path int true "System ID"
// @Success 200 {array} models.Station
// @Router /systems/{systemID}/stations [get]
func GetStationsBySystemIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "systemID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid system ID")
		return
	}

	stations, err := service.GetStationsBySystemID(id)
	if err != nil {
		log.Printf("Error fetching stations for system %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve stations")
		return
	}
	if stations == nil {
		respondError(w, http.StatusNotFound, "No stations found for this system")
		return
	}
	respondJSON(w, http.StatusOK, stations)
} 