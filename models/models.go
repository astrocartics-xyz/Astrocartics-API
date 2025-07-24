package models

// Region represents a region in EVE Online.
type Region struct {
	RegionID   int    `json:"region_id"`
	RegionName string `json:"region_name"`
}

// Constellation represents a constellation in EVE Online.
type Constellation struct {
	ConstellationID   int    `json:"constellation_id"`
	ConstellationName string `json:"constellation_name"`
	RegionID          int    `json:"region_id"`
}

// System represents a solar system in EVE Online.
// Includes spectral_class directly for map visualization.
type System struct {
	SystemID        int     `json:"system_id"`
	SystemName      string  `json:"system_name"`
	SecurityStatus  float64 `json:"security_status"`
	SecurityClass   *string `json:"security_class"` // Use pointer for nullable string
	XPos            float64 `json:"x_pos"`
	YPos            float64 `json:"y_pos"`
	ZPos            float64 `json:"z_pos"`
	ConstellationID int     `json:"constellation_id"`
	SpectralClass   *string `json:"spectral_class"` // Use pointer for nullable string
}

// Stargate represents a stargate connection between systems.
type Stargate struct {
	StargateID            int    `json:"stargate_id"`
	StargateName          string `json:"stargate_name"`
	SystemID              int    `json:"system_id"`
	DestinationStargateID int    `json:"destination_stargate_id"`
	DestinationSystemID   int    `json:"destination_system_id"`
}

// SpectralClassCount represents the count of systems for a given spectral class.
type SpectralClassCount struct {
	SpectralClass string `json:"spectral_class"`
	SystemCount   int    `json:"system_count"`
}

// Planet represents a planet in a system.
type Planet struct {
	PlanetID          int     `json:"planet_id"`
	PlanetName        string  `json:"planet_name"`
	SystemID          int     `json:"system_id"`
	Type              *string `json:"type"`
	MoonCount         int     `json:"moon_count"`
	AsteroidBeltCount int     `json:"asteroid_belt_count"`
}

// Station represents a station in a system.
type Station struct {
	StationID   int    `json:"station_id"`
	StationName string `json:"station_name"`
	SystemID    int    `json:"system_id"`
} 