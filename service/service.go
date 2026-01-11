package service

import (
	"fmt"

	"github.com/astrocartics-xyz/Astrocartics-API/dba"
	"github.com/astrocartics-xyz/Astrocartics-API/models"
)

func GetAllRegions() ([]models.Region, error) {
	return dba.GetAllRegions()
}

func GetRegionByID(id int) (*models.Region, error) {
	return dba.GetRegionByID(id)
}

func GetRegionByName(name string) (*models.Region, error) {
	return dba.GetRegionByName(name)
}

func GetAllConstellations() ([]models.Constellation, error) {
	return dba.GetAllConstellations()
}

func GetConstellationByIDOrRegionID(id int) ([]models.Constellation, error) {
	return dba.GetConstellationByIDOrRegionID(id)
}

func GetConstellationByName(name string) (*models.Constellation, error) {
	return dba.GetConstellationByName(name)
}

func GetAllSystems() ([]models.System, error) {
	return dba.GetAllSystems()
}

func GetSystemByIDOrConstellationID(id int) ([]models.System, error) {
	return dba.GetSystemByIDOrConstellationID(id)
}

func GetSystemByName(name string) (*models.System, error) {
	return dba.GetSystemByName(name)
}

func GetSystemNameByID(systemID int) (string, error) {
	return dba.GetSystemNameByID(systemID)
}

func GetSystemsByRegionID(id int) ([]models.System, error) {
	return dba.GetSystemsByRegionID(id)
}

func GetAllStargates() ([]models.Stargate, error) {
	return dba.GetAllStargates()
}

func GetStargateBySystemID(id int) ([]models.Stargate, error) {
	return dba.GetStargateBySystemID(id)
}

func GetStargateByConstellationID(id int) ([]models.Stargate, error) {
        return dba.GetStargateByConstellationID(id)
}

func GetStargateByRegionID(id int) ([]models.Stargate, error) {
        return dba.GetStargateByRegionID(id)
}

func GetSpectralClassCounts() ([]models.SpectralClassCount, error) {
	return dba.GetSpectralClassCounts()
}

// Planet service functions
func GetAllPlanets() ([]models.Planet, error) {
	return dba.GetAllPlanets()
}

func GetPlanetByID(id int) (*models.Planet, error) {
	return dba.GetPlanetByID(id)
}

func GetPlanetByName(name string) (*models.Planet, error) {
	return dba.GetPlanetByName(name)
}

func GetPlanetsBySystemID(id int) ([]models.Planet, error) {
	return dba.GetPlanetsBySystemID(id)
}

// Station service functions
func GetAllStations() ([]models.Station, error) {
	return dba.GetAllStations()
}

func GetStationByID(id int) (*models.Station, error) {
	return dba.GetStationByID(id)
}

func GetStationByName(name string) (*models.Station, error) {
	return dba.GetStationByName(name)
}

func GetStationsBySystemID(id int) ([]models.Station, error) {
	return dba.GetStationsBySystemID(id)
}

// isValidKillMode validates if the mode is one of the supported modes
func isValidKillMode(mode string) bool {
	validModes := map[string]bool{
		"hour": true,
		"day": true,
		"week": true,
		"month": true,
	}
	return validModes[mode]
}

// Used for heat map display by regionID
func GetSystemHeatmapReportByRegionMode(regionID int, mode string) (models.HeatmapReport, error) {
	var empty models.HeatmapReport
	// validate mode
	if !isValidKillMode(mode) {
		return empty, fmt.Errorf("invalid mode: %s; supported: 'hour','day','week','month'", mode)
	}
	// dba returns: regionName, points, windowStart, windowEnd, error
	regionName, points, windowStart, windowEnd, err := dba.GetSystemHeatmapByRegionMode(regionID, mode)
	if err != nil {
		return empty, fmt.Errorf("failed to fetch system heatmap: %w", err)
	}
	// Compute total kills
	total := 0
	for _, p := range points {
		total += p.Kills
	}
	// Generate report
	report := models.HeatmapReport{
		Mode:        mode,
		RegionID:    regionID,
		RegionName:  regionName,
		WindowStart: windowStart,
		WindowEnd:   windowEnd,
		TotalKills:  total,
		Buckets:     points,
	}
	return report, nil
}
// GetLast15KillmailsBySystemID returns the last 15 killmails for a system.
func GetRecentKillmailsBySystemID(systemID int) ([]models.Killmails, error) {
	return dba.GetRecentKillmailsBySystemID(systemID)
}

// GetKillCountBySystemID retrieves kill counts by system ID and calculates the total
func GetKillCountBySystemID(systemID int, mode string) (string, int, []models.PeriodCount, error) {
	// Validate mode
	if !isValidKillMode(mode) {
		return "", 0, nil, fmt.Errorf("invalid mode: %s; supported: 'day', 'week', 'month'", mode)
	}
	// Fetch data from the dba layer
	systemName, buckets, err := dba.GetKillsBySystemID(systemID, mode)
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to fetch kills: %w", err)
	}
	// Use `systemName` if needed, or discard it with `_` if it’s not needed
	fmt.Printf("System Name: %s\n", systemName)
	// Calculate total kills
	total := 0
	for _, bucket := range buckets {
		total += bucket.Count
	}
	return systemName, total, buckets, nil
}

// GetKillCountByConstellationID retrieves kill counts by constellation ID and calculates the total
func GetKillCountByConstellationID(constellationID int, mode string) (string, int, []models.PeriodCount, error) {
	// Validate mode
	if !isValidKillMode(mode) {
		return "", 0, nil, fmt.Errorf("invalid mode: %s; supported: 'day', 'week', 'month'", mode)
	}
	// Fetch data from the dba layer
	constellationName, buckets, err := dba.GetKillsByConstellationID(constellationID, mode)
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to fetch kills: %w", err)
	}
	// Use `constellationName` if needed, or discard it with `_` if it’s not needed
	fmt.Printf("System Name: %s\n", constellationName)
	// Calculate total kills
	total := 0
	for _, bucket := range buckets {
		total += bucket.Count
	}
	return constellationName, total, buckets, nil
}

// GetKillCountByRegionID retrieves kill counts by region ID and calculates the total
func GetKillCountByRegionID(regionID int, mode string) (string, int, []models.PeriodCount, error) {
	// Validate mode
	if !isValidKillMode(mode) {
		return "", 0, nil, fmt.Errorf("invalid mode: %s; supported: 'day', 'week', 'month'", mode)
	}
	// Fetch data from the dba layer
	regionName, buckets, err := dba.GetKillsByRegionID(regionID, mode)
	if err != nil {
		return "", 0, nil, fmt.Errorf("failed to fetch kills: %w", err)
	}
	// Use `systemName` if needed, or discard it with `_` if it’s not needed
	fmt.Printf("System Name: %s\n", regionName)
	// Calculate total kills
	total := 0
	for _, bucket := range buckets {
		total += bucket.Count
	}
	return regionName, total, buckets, nil
}

// Get top regions by fetching top regions by kill count for a given time window
func GetTopRegionsByKills(mode string) ([]models.RegionKillCount, error) {
	return dba.GetTopRegionsByKills(mode)
}

// Get top constellations by fetching top constellations by kill count for a given time window
func GetTopConstellationsByKills(mode string) ([]models.ConstellationKillCount, error) {
	return dba.GetTopConstellationsByKills(mode)
}

// Get top systems by fetching top systems by kill count for a given time window
func GetTopSystemsByKills(mode string) ([]models.SystemKillCount, error) {
	return dba.GetTopSystemsByKills(mode)
}
