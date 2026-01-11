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

// Systtem represents a point in Eve Online
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
	RegionID 	int	`json:"region_id"`
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

// swagger:model Station
type Station struct {
	StationID   int    `json:"station_id"`
	StationName string `json:"station_name"`
	SystemID    int    `json:"system_id"`
}

// swagger:model PeriodCount
type PeriodCount struct {
	Period string `json:"period"` // e.g. "2025-10-20"
    	Count  int    `json:"count"`
	DestroyedValue float64 `json:"destroyed_value"`
	DroppedValue   float64 `json:"dropped_value"`
}

// swagger:model Killmails
type Killmails struct {
	KillmailID     int64     `json:"killmail_id"`
        KillmailHash   string    `json:"killmail_hash"`
	SolarSystemID  int       `json:"system_id"`
	KillmailTime   string    `json:"killmail_time"`
	DestroyedValue float64   `json:"destroyed_value"`
	DroppedValue   float64   `json:"dropped_value"`
	TotalValue     float64   `json:"total_value"`
	FittedValue    float64   `json:"fitted_value"`
	VictimShip     int64     `json:"victim_ship"`
	KillShip       int64	 `json:"kill_ship"`
}

// For heatmap-by-mode results
type SystemPeriodHeatPoint struct {
	SystemID       int   `json:"system_id"`
	SystemName     string  `json:"system_name"`
	Kills          int     `json:"kills"`
	DestroyedValue float64 `json:"destroyed_value"`
	DroppedValue   float64 `json:"dropped_value"`
}

// swagger:model Heatmap
type HeatmapReport struct {
	Mode       string                           `json:"mode"`
	RegionID   int                              `json:"region_id"`
	RegionName string                           `json:"region_name"`
	TotalKills int                              `json:"total"`
	WindowStart string                  	    `json:"window_start"`  // ISO8601 start of the sliding window
	WindowEnd  string                  	    `json:"window_end"`    // ISO8601 end of the sliding window (usually now)
	Buckets    []SystemPeriodHeatPoint	    `json:"buckets"`
}

// Base struct for common fields
type KillStats struct {
	Mode    string        `json:"mode"`
	Total   int           `json:"total"`
	Buckets []PeriodCount `json:"buckets"`
}

// swagger:model SystemKills
type SystemKills struct {
	SystemID   int	      `json:"system_id"`
	SystemName string     `json:"system_name"`
	KillStats
}

// swagger:model ConstellationKills
type ConstellationKills struct {
	ConstellationID   int    `json:"constellation_id"`
	ConstellationName string `json:"constellation_name"`
	KillStats
}

// swagger:model RegionKills
type RegionKills struct {
	RegionID   int           `json:"region_id"`
	RegionName string        `json:"region_name"`
	KillStats
}

// swagger:model RegionKillCount
type RegionKillCount struct {
	RegionID   int		`json:"region_id"`
	RegionName string	`json:"region_name"`
	TotalKills int		`json:"total_kills"`
}

// swagger:model ConstellationKillCount
type ConstellationKillCount struct {
	ConstellationID   int	 `json:"constellation_id"`
	ConstellationName string `json:"constellation_name"`
	TotalKills int		 `json:"total_kills"`
}

// swagger:model SystemKillCount
type SystemKillCount struct {
	SystemID   int		`json:"system_id"`
	SystemName string	`json:"system_name"`
	TotalKills int		`json:"total_kills"`
}

