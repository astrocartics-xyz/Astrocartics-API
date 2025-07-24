package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/astrocartics-xyz/Astrocartics-API/service"
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

// GetRegionsHandler handles requests for all regions.
func GetRegionsHandler(w http.ResponseWriter, r *http.Request) {
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
	// Extract ID from URL path (e.g., /api/v1/regions/id/)
	idStr := r.URL.Path[len("/api/v1/regions/id/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid region ID")
		return
	}
	// Check database and serialize.
	region, err := service.GetRegionByID(id)
	if err != nil {
		log.Printf("Error fetching constellation %d: %v", id, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve region")
		return
	}
	if region == nil {
		respondError(w, http.StatusNotFound, "Region not found")
		return
	}
	respondJSON(w, http.StatusOK, region)
}

// GetRegionByNameHandler handles requests for a single region by name.
// Expected URL path: /api/v1/regions/name/{regionName}
func GetRegionByNameHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the system name from the URL path.
	regionName := strings.TrimPrefix(r.URL.Path, "/api/v1/regions/name/")
	if regionName == "" {
		respondError(w, http.StatusBadRequest, "Region name is missing in the URL")
		return
	}
	// Call the database function to get the system by name.
	region, err := service.GetRegionByName(regionName)
	if err != nil {
		log.Printf("Error fetching region by name '%s': %v", regionName, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve region")
		return
	}
	// If no system is found, respond with 404 Not Found.
	if region == nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Region '%s' not found", regionName))
		return
	}

	// Respond with the fetched region in JSON format.
	respondJSON(w, http.StatusOK, region)
}

// GetConstellationsHandler handles requests for all regions.
func GetConstellationsHandler(w http.ResponseWriter, r *http.Request) {
	constellations, err := service.GetAllConstellations()
	if err != nil {
		log.Printf("Error fetching regions: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve regions")
		return
	}
	respondJSON(w, http.StatusOK, constellations)
}

// GetConstellationByIDHandler handles requests for a single constellation by ID.
func GetConstellationByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path (e.g., /api/v1/constellations/id/)
	idStr := r.URL.Path[len("/api/v1/constellations/id/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid constellation or region ID")
		return
	}
	// Check database and serialize.
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

// GetConstellationByNameHandler handles requests for a single constellation by name.
// Expected URL path: /api/v1/constellations/name/{constellationName}
func GetConstellationByNameHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the constellation name from the URL path.
	constellationName := strings.TrimPrefix(r.URL.Path, "/api/v1/constellations/name/")
	if constellationName == "" {
		respondError(w, http.StatusBadRequest, "Constellation name is missing in the URL")
		return
	}
	// Call the database function to get the constellation by name.
	constellation, err := service.GetConstellationByName(constellationName)
	if err != nil {
		log.Printf("Error fetching constellation by name '%s': %v", constellationName, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve constellation")
		return
	}
	// If no constellation is found, respond with 404 Not Found.
	if constellation == nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf("System '%s' not found", constellationName))
		return
	}
	// Respond with the fetched constellation in JSON format.
	respondJSON(w, http.StatusOK, constellation)
}

// GetSystemsHandler handles requests for all systems.
func GetSystemsHandler(w http.ResponseWriter, r *http.Request) {
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
	// Extract ID from URL path (e.g., /api/v1/systems/id/30000142)
	idStr := r.URL.Path[len("/api/v1/systems/id/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid system or constellation ID")
		return
	}
	// Check database and serialize.
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

// GetSystemByNameHandler handles requests for a single system by name.
// Expected URL path: /api/v1/systems/name/{systemName}
func GetSystemByNameHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the system name from the URL path.
	systemName := strings.TrimPrefix(r.URL.Path, "/api/v1/systems/name/")
	if systemName == "" {
		respondError(w, http.StatusBadRequest, "System name is missing in the URL")
		return
	}
	// Call the database function to get the system by name.
	system, err := service.GetSystemByName(systemName)
	if err != nil {
		log.Printf("Error fetching system by name '%s': %v", systemName, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve system")
		return
	}
	// If no system is found, respond with 404 Not Found.
	if system == nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf("System '%s' not found", systemName))
		return
	}
	// Respond with the fetched system in JSON format.
	respondJSON(w, http.StatusOK, system)
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

// GetStargatesBySystemIDHandler handles requests to get stargates by system ID.
// Expected URL path: /api/v1/stargates/id/{systemID}
func GetStargateBySystemIDHandler(w http.ResponseWriter, r *http.Request) {
	// Assume the ID is the last segment after "/api/v1/stargates/id/".
	systemIDStr := r.URL.Path[len("/api/v1/stargates/id/"):]
	if systemIDStr == "" {
		respondError(w, http.StatusBadRequest, "system ID is missing in the URL")
		return
	}
	// Convert the string ID to an integer.
	systemID, err := strconv.Atoi(systemIDStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid system ID format")
		return
	}
	// Call the database function to get stargates by system ID.
	stargates, err := service.GetStargateBySystemID(systemID)
	if err != nil {
		log.Printf("Error fetching stargates by system ID %d: %v", systemID, err)
		respondError(w, http.StatusInternalServerError, "Failed to retrieve stargates for the specified system")
		return
	}
	// If no stargates are found, respond with 404 Not Found.
	if len(stargates) == 0 {
		respondError(w, http.StatusNotFound, fmt.Sprintf("No stargates found for system ID %d", systemID))
		return
	}
	// Respond with the fetched stargates in JSON format.
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