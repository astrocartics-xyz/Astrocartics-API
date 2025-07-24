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

// GetRegionsHandler handles requests for all regions, or a region by name if a query param is provided.
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

// GetRegionByIDHandler handles requests for a single region by ID.
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

// GetConstellationsHandler handles requests for all constellations, or a constellation by name if a query param is provided.
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

// GetConstellationByIDHandler handles requests for a single constellation by ID.
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

// GetConstellationsByRegionIDHandler handles requests for constellations in a specific region.
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

// GetSystemsHandler handles requests for all systems, or a system by name if a query param is provided.
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

// GetSystemByIDHandler handles requests for a single system by ID.
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

// GetSystemsByConstellationIDHandler handles requests for systems in a specific constellation.
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

// GetStargatesHandler handles requests for all stargates.
func GetStargatesHandler(w http.ResponseWriter, r *http.Request) {
	stargates, err := service.GetAllStargates()
	if err != nil {
		log.Printf("Error fetching stargates: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve stargates")
		return
	}
	respondJSON(w, http.StatusOK, stargates)
}

// GetStargateBySystemIDHandler handles requests to get stargates by system ID.
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

// GetSpectralClassCountsHandler handles requests for system counts by spectral class.
func GetSpectralClassCountsHandler(w http.ResponseWriter, r *http.Request) {
	counts, err := service.GetSpectralClassCounts()
	if err != nil {
		log.Printf("Error fetching spectral class counts: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve spectral class counts")
		return
	}
	respondJSON(w, http.StatusOK, counts)
} 