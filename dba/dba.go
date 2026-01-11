package dba

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astrocartics-xyz/Astrocartics-API/models"
)

// GetAllRegions fetches all regions from the database.
func GetAllRegions() ([]models.Region, error) {
	db := GetDB()
	rows, err := db.Query("SELECT region_id, region_name FROM regions ORDER BY region_name")
	if err != nil {
		return nil, fmt.Errorf("failed to query regions: %w", err)
	}
	defer rows.Close()
	var regions []models.Region
	for rows.Next() {
		var r models.Region
		if err := rows.Scan(&r.RegionID, &r.RegionName); err != nil {
			return nil, fmt.Errorf("failed to scan region row: %w", err)
		}
		regions = append(regions, r)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for regions: %w", err)
	}
	return regions, nil
}

// GetRegionByID fetching region by ID from the database.
func GetRegionByID(id int) (*models.Region, error) {
	db := GetDB()
	var s models.Region
	err := db.QueryRow("SELECT region_id, region_name FROM regions WHERE region_id = $1", id).
		Scan(&s.RegionID, &s.RegionName)
	if err == sql.ErrNoRows {
		return nil, nil // Not found
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query region by ID: %w", err)
	}
	return &s, nil
}

// GetRegionByName fetches all regions by name.
func GetRegionByName(name string) (*models.Region, error) {
	db := GetDB()
	var s models.Region
	// Modify the SQL query to filter by region_name
	err := db.QueryRow("SELECT region_id, region_name FROM regions WHERE region_name = $1", name).
		Scan(&s.RegionID, &s.RegionName)
	if err == sql.ErrNoRows {
		return nil, nil // Not found
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query region by name: %w", err)
	}
	return &s, nil
}

// GetAllConstellations fetches all constellations from the database.
func GetAllConstellations() ([]models.Constellation, error) {
	db := GetDB()
	// Query all columns from the 'constellations' table, ordered by constellation_name.
	rows, err := db.Query("SELECT constellation_id, constellation_name, region_id FROM constellations ORDER BY constellation_name")
	if err != nil {
		// Return an error if the query execution fails.
		return nil, fmt.Errorf("failed to query constellations: %w", err)
	}
	defer rows.Close() // Ensure the rows are closed after the function returns
	var constellations []models.Constellation // Slice to hold the fetched constellation objects
	for rows.Next() {
		var c models.Constellation // Variable to scan each row into
		// Scan the values from the current row into the Constellation struct fields.
		// The order of scanning must match the order of columns in the SELECT statement.
		if err := rows.Scan(&c.ConstellationID, &c.ConstellationName, &c.RegionID); err != nil {
			// Return an error if scanning a row fails (e.g., type mismatch, database issue).
			return nil, fmt.Errorf("failed to scan constellation row: %w", err)
		}
		constellations = append(constellations, c) // Add the scanned constellation to the slice
	}
	// Check for any errors that occurred during iteration over the rows.
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for constellations: %w", err)
	}
	return constellations, nil // Return the slice of constellations and no error
}

// GetConstellationByIDOrRegionID fetches a single constellation by its ConstellationID and RegionID.
func GetConstellationByIDOrRegionID(id int) ([]models.Constellation, error) {
	db := GetDB()
	// Base query with an OR condition for the single ID parameter
	// We use $1 for both conditions as it's the same input ID.
	sqlQuery := "SELECT constellation_id, constellation_name, region_id FROM constellations WHERE constellation_id = $1 OR region_id = $1 ORDER BY constellation_name"
	// Execute the query with the single ID parameter
	rows, err := db.Query(sqlQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query constellations by constellation or region ID: %w", err)
	}
	defer rows.Close()
	// Get constellation or constellations.
	var constellations []models.Constellation
	for rows.Next() {
		var c models.Constellation
		if err := rows.Scan(&c.ConstellationID, &c.ConstellationName, &c.RegionID); err != nil {
			return nil, fmt.Errorf("failed to scan constellation row: %w", err)
		}
		constellations = append(constellations, c)
	}
	// Error handling
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for constellations: %w", err)
	}
	// Return
	return constellations, nil
}

// GetConstellationByName fetches a single constellation by its name.
func GetConstellationByName(name string) (*models.Constellation, error) {
	db := GetDB()
	var c models.Constellation // Declare a variable of type Constellation to hold the fetched data
	// Query a single row from the 'constellations' table where constellation_name matches the provided name.
	// Use $1 as a placeholder for the name parameter.
	err := db.QueryRow("SELECT constellation_id, constellation_name, region_id FROM constellations WHERE constellation_name = $1", name).
		Scan(&c.ConstellationID, &c.ConstellationName, &c.RegionID)
	// Check if no rows were returned (constellation not found).
	if err == sql.ErrNoRows {
		return nil, nil // Return nil for both Constellation and error to indicate "not found"
	}
	// Check for any other database errors.
	if err != nil {
		return nil, fmt.Errorf("failed to query constellation by name: %w", err)
	}
	return &c, nil // Return a pointer to the found Constellation and no error
}

// GetAllSystems fetches all systems from the database.
func GetAllSystems() ([]models.System, error) {
	db := GetDB()
	rows, err := db.Query(
		`SELECT s.system_id,
			s.system_name,
			s.security_status,
			s.security_class,
			s.x_pos,
			s.y_pos,
			s.z_pos,
			s.constellation_id,
			c.region_id,
			s.spectral_class
		FROM systems s
		JOIN constellations c ON c.constellation_id = s.constellation_id
		ORDER BY s.system_name`)
	if err != nil {
		return nil, fmt.Errorf("failed to query systems: %w", err)
	}
	defer rows.Close()
	var systems []models.System
	for rows.Next() {
		var s models.System
		if err := rows.Scan(&s.SystemID, &s.SystemName, &s.SecurityStatus, &s.SecurityClass, &s.XPos, &s.YPos, &s.ZPos, &s.ConstellationID, &s.RegionID, &s.SpectralClass); err != nil {
			return nil, fmt.Errorf("failed to scan system row: %w", err)
		}
		systems = append(systems, s)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for systems: %w", err)
	}
	return systems, nil
}

// GetSystemByIDOrConstellationID fetches a single system by its SystemID and ConstellationID.
func GetSystemByIDOrConstellationID(id int) ([]models.System, error) {
	db := GetDB()
	// Base query with an OR condition for the single ID parameter.
	// We use $1 for both conditions as it's the same input ID.
	sqlQuery := `SELECT s.system_id,
			s.system_name,
			s.security_status,
			s.security_class,
			s.x_pos,
			s.y_pos,
			s.z_pos,
			s.constellation_id,
			c.region_id,
			s.spectral_class
		FROM systems s
		JOIN constellations c ON c.constellation_id = s.constellation_id
		WHERE s.system_id = $1 OR s.constellation_id = $1
		ORDER BY s.system_name`
	// Execute the query with the single ID parameter.
	rows, err := db.Query(sqlQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query systems by system or constellation ID: %w", err)
	}
	defer rows.Close()
	// Get system or systems.
	var systems []models.System
	for rows.Next() {
		var s models.System
		if err := rows.Scan(
			&s.SystemID,
			&s.SystemName,
			&s.SecurityStatus,
			&s.SecurityClass,
			&s.XPos,
			&s.YPos,
			&s.ZPos,
			&s.ConstellationID,
			&s.RegionID,
			&s.SpectralClass,
		); err != nil {
			return nil, fmt.Errorf("failed to scan system row: %w", err)
		}
		systems = append(systems, s)
	}
	// Error handling
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for systems: %w", err)
	}
	// Returns
	return systems, nil
}

// GetSystemsByRegionID fetches all systems for a specific region ID.
func GetSystemsByRegionID(regionID int) ([]models.System, error) {
	db := GetDB()
	sqlQuery := `
		SELECT s.system_id, s.system_name, s.security_status, s.security_class, s.x_pos, s.y_pos, s.z_pos, s.constellation_id, c.region_id, s.spectral_class
		FROM systems s
		JOIN constellations c ON s.constellation_id = c.constellation_id
		WHERE c.region_id = $1
		ORDER BY s.system_name`
	rows, err := db.Query(sqlQuery, regionID)
	if err != nil {
		return nil, fmt.Errorf("failed to query systems by region ID: %w", err)
	}
	defer rows.Close()
	var systems []models.System
	for rows.Next() {
		var s models.System
		if err := rows.Scan(
			&s.SystemID,
			&s.SystemName,
			&s.SecurityStatus,
			&s.SecurityClass,
			&s.XPos,
			&s.YPos,
			&s.ZPos,
			&s.ConstellationID,
			&s.RegionID,
			&s.SpectralClass,
		); err != nil {
			return nil, fmt.Errorf("failed to scan system row: %w", err)
		}
		systems = append(systems, s)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for systems by region ID: %w", err)
	}
	return systems, nil
}

// GetSystemByName fetches a single system by its name.
func GetSystemByName(name string) (*models.System, error) {
	db := GetDB()
	var s models.System // Declare a variable of type System to hold the fetched data
	// Modify the SQL query to filter by system_name instead of system_id.
	// Use $1 as a placeholder for the name parameter.
	err := db.QueryRow(`SELECT s.system_id,
		s.system_name,
		s.security_status,
		s.security_class,
		s.x_pos,
		s.y_pos,
		s.z_pos,
		s.constellation_id,
		c.region_id,
		s.spectral_class
	FROM systems s
	JOIN constellations c ON c.constellation_id = s.constellation_id
	WHERE s.system_name = $1
	LIMIT 1`, name).Scan(&s.SystemID, &s.SystemName, &s.SecurityStatus, &s.SecurityClass, &s.XPos, &s.YPos, &s.ZPos, &s.ConstellationID, &s.RegionID, &s.SpectralClass)
	// Check if no rows were returned (system not found).
	if err == sql.ErrNoRows {
		return nil, nil // Return nil for both System and error to indicate "not found"
	}
	// Check for any other database errors.
	if err != nil {
		return nil, fmt.Errorf("failed to query system by name: %w", err)
	}
	return &s, nil // Return a pointer to the found System and no error
}

// GetSystemIDByName fetches single system id by name
func GetSystemNameByID(systemID int) (string, error) {
	db := GetDB() // Get the database connection
	// Query to fetch the system name by its ID
	var systemName string
	query := `SELECT system_name FROM systems WHERE system_id = $1`
	err := db.QueryRow(query, systemID).Scan(&systemName)
	if err == sql.ErrNoRows {
		// If no rows are found, return an error
		return "", fmt.Errorf("system with ID %d not found", systemID)
	} else if err != nil {
		// Handle other query errors
		return "", fmt.Errorf("failed to fetch system name: %w", err)
	}
	return systemName, nil
}

// GetAllStargates fetches all stargate connections.
func GetAllStargates() ([]models.Stargate, error) {
	db := GetDB()
	rows, err := db.Query("SELECT stargate_id, stargate_name, system_id, destination_stargate_id, destination_system_id FROM stargates ORDER BY stargate_name")
	if err != nil {
		return nil, fmt.Errorf("failed to query stargates: %w", err)
	}
	defer rows.Close()
	var stargates []models.Stargate
	for rows.Next() {
		var sg models.Stargate
		if err := rows.Scan(&sg.StargateID, &sg.StargateName, &sg.SystemID, &sg.DestinationStargateID, &sg.DestinationSystemID); err != nil {
			return nil, fmt.Errorf("failed to scan stargate row: %w", err)
		}
		stargates = append(stargates, sg)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for stargates: %w", err)
	}
	return stargates, nil
}

// GetStargateBySystemID fetches stargates associated with a given system ID.
func GetStargateBySystemID(systemID int) ([]models.Stargate, error) {
	db := GetDB()
	// Prepare the SQL query with a WHERE clause to filter by system_id.
	rows, err := db.Query("SELECT stargate_id, stargate_name, system_id, destination_stargate_id, destination_system_id FROM stargates WHERE system_id = $1 ORDER BY stargate_name", systemID)
	if err != nil {
		// Return an error if the query itself fails (e.g., database connection issues, syntax error).
		return nil, fmt.Errorf("failed to query stargates by system ID: %w", err)
	}
	defer rows.Close() // Ensure rows are closed after the function returns
	var stargates = make([]models.Stargate, 0) // Slice to hold the fetched stargate objects
	for rows.Next() {
		var sg models.Stargate // Variable to scan each row into
		// Scan the values from the current row into the Stargate struct fields.
		if err := rows.Scan(&sg.StargateID, &sg.StargateName, &sg.SystemID, &sg.DestinationStargateID, &sg.DestinationSystemID); err != nil {
			// Return an error if scanning a row fails (e.g., type mismatch).
			return nil, fmt.Errorf("failed to scan stargate row: %w", err)
		}
		stargates = append(stargates, sg) // Add the scanned stargate to the slice
	}
	// Check for any errors that occurred during iteration over the rows.
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for stargates by system ID: %w", err)
	}
	return stargates, nil // Return the slice of stargates and no error
}

// GetStargateByConstellationID fetches stargates with a given constellation ID
func GetStargateByConstellationID(constellationID int) ([]models.Stargate, error) {
	db := GetDB()
	// Query
	rows, err := db.Query(`SELECT st.stargate_id, st.stargate_name, st.system_id,
		st.destination_stargate_id, st.destination_system_id
		FROM stargates st
		WHERE st.system_id IN (SELECT system_id FROM systems WHERE constellation_id = $1)
		ORDER BY st.stargate_name`, constellationID)
	// Check for errors
	if err != nil {
		return nil, fmt.Errorf("failed to query stargates by constellation id: %w", err)
	}
	defer rows.Close()
	// Slicing rows
	var stargates = make([]models.Stargate, 0)
	for rows.Next() {
		var sg models.Stargate
		if err := rows.Scan(&sg.StargateID, &sg.StargateName, &sg.SystemID, &sg.DestinationStargateID, &sg.DestinationSystemID); err != nil {
			return nil, fmt.Errorf("failed to scan stargate row: %w", err)
		}
		stargates = append(stargates, sg)
	}
	// Check for errors
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for stargates by constellation id: %w", err)
	}
	return stargates, nil
}

// GetStargateByRegionID fetches stargates with a given region ID.
func GetStargateByRegionID(regionID int) ([]models.Stargate, error) {
	db := GetDB()
	// Query
	rows, err := db.Query(`SELECT st.stargate_id, st.stargate_name, st.system_id, st.destination_stargate_id, st.destination_system_id
		FROM stargates st
		WHERE st.system_id IN (SELECT s.system_id
		FROM systems s
		JOIN constellations c ON s.constellation_id = c.constellation_id
		WHERE c.region_id = $1)
		ORDER BY st.stargate_name`, regionID)
	// Check for errors
	if err != nil {
		return nil, fmt.Errorf("failed to query stargates by region id: %w", err)
	}
	defer rows.Close()
	// Row slicing
	var stargates = make([]models.Stargate, 0)
	for rows.Next() {
		var sg models.Stargate
		if err := rows.Scan(&sg.StargateID, &sg.StargateName, &sg.SystemID, &sg.DestinationStargateID, &sg.DestinationSystemID); err != nil {
			return nil, fmt.Errorf("failed to scan stargate row: %w", err)
		}
		stargates = append(stargates, sg)
	}
	// Check for errors
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for stargates by region id: %w", err)
	}
	return stargates, nil
}

// GetSpectralClassCounts fetches counts of systems by spectral class.
func GetSpectralClassCounts() ([]models.SpectralClassCount, error) {
	db := GetDB()
	rows, err := db.Query("SELECT spectral_class, COUNT(system_id) AS system_count FROM systems WHERE spectral_class IS NOT NULL GROUP BY spectral_class ORDER BY system_count DESC")
	if err != nil {
		return nil, fmt.Errorf("failed to query spectral class counts: %w", err)
	}
	defer rows.Close()
	var counts []models.SpectralClassCount
	for rows.Next() {
		var scc models.SpectralClassCount
		if err := rows.Scan(&scc.SpectralClass, &scc.SystemCount); err != nil {
			return nil, fmt.Errorf("failed to scan spectral class count row: %w", err)
		}
		counts = append(counts, scc)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for spectral class counts: %w", err)
	}
	return counts, nil
}

// Planet database functions
func GetAllPlanets() ([]models.Planet, error) {
	db := GetDB()
	rows, err := db.Query("SELECT planet_id, planet_name, system_id, type, moon_count, asteroid_belt_count FROM planets ORDER BY planet_name")
	if err != nil {
		return nil, fmt.Errorf("failed to query planets: %w", err)
	}
	defer rows.Close()
	var planets []models.Planet
	for rows.Next() {
		var p models.Planet
		if err := rows.Scan(&p.PlanetID, &p.PlanetName, &p.SystemID, &p.Type, &p.MoonCount, &p.AsteroidBeltCount); err != nil {
			return nil, fmt.Errorf("failed to scan planet row: %w", err)
		}
		planets = append(planets, p)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for planets: %w", err)
	}
	return planets, nil
}

func GetPlanetByID(id int) (*models.Planet, error) {
	db := GetDB()
	var p models.Planet
	err := db.QueryRow("SELECT planet_id, planet_name, system_id, type, moon_count, asteroid_belt_count FROM planets WHERE planet_id = $1", id).
		Scan(&p.PlanetID, &p.PlanetName, &p.SystemID, &p.Type, &p.MoonCount, &p.AsteroidBeltCount)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query planet by ID: %w", err)
	}
	return &p, nil
}

func GetPlanetByName(name string) (*models.Planet, error) {
	db := GetDB()
	var p models.Planet
	err := db.QueryRow("SELECT planet_id, planet_name, system_id, type, moon_count, asteroid_belt_count FROM planets WHERE planet_name = $1", name).
		Scan(&p.PlanetID, &p.PlanetName, &p.SystemID, &p.Type, &p.MoonCount, &p.AsteroidBeltCount)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query planet by name: %w", err)
	}
	return &p, nil
}

func GetPlanetsBySystemID(systemID int) ([]models.Planet, error) {
	db := GetDB()
	rows, err := db.Query("SELECT planet_id, planet_name, system_id, type, moon_count, asteroid_belt_count FROM planets WHERE system_id = $1 ORDER BY planet_name", systemID)
	if err != nil {
		return nil, fmt.Errorf("failed to query planets by system ID: %w", err)
	}
	defer rows.Close()
	var planets []models.Planet
	for rows.Next() {
		var p models.Planet
		if err := rows.Scan(&p.PlanetID, &p.PlanetName, &p.SystemID, &p.Type, &p.MoonCount, &p.AsteroidBeltCount); err != nil {
			return nil, fmt.Errorf("failed to scan planet row: %w", err)
		}
		planets = append(planets, p)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for planets by system ID: %w", err)
	}
	return planets, nil
}

func GetAllStations() ([]models.Station, error) {
	db := GetDB()
	rows, err := db.Query("SELECT station_id, station_name, system_id FROM stations ORDER BY station_name")
	if err != nil {
		return nil, fmt.Errorf("failed to query stations: %w", err)
	}
	defer rows.Close()
	var stations []models.Station
	for rows.Next() {
		var s models.Station
		if err := rows.Scan(&s.StationID, &s.StationName, &s.SystemID); err != nil {
			return nil, fmt.Errorf("failed to scan station row: %w", err)
		}
		stations = append(stations, s)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for stations: %w", err)
	}
	return stations, nil
}

func GetStationByID(id int) (*models.Station, error) {
	db := GetDB()
	var s models.Station
	err := db.QueryRow("SELECT station_id, station_name, system_id FROM stations WHERE station_id = $1", id).
		Scan(&s.StationID, &s.StationName, &s.SystemID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query station by ID: %w", err)
	}
	return &s, nil
}

func GetStationByName(name string) (*models.Station, error) {
	db := GetDB()
	var s models.Station
	err := db.QueryRow("SELECT station_id, station_name, system_id FROM stations WHERE station_name = $1", name).
		Scan(&s.StationID, &s.StationName, &s.SystemID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query station by name: %w", err)
	}
	return &s, nil
}

func GetStationsBySystemID(systemID int) ([]models.Station, error) {
	db := GetDB()
	rows, err := db.Query("SELECT station_id, station_name, system_id FROM stations WHERE system_id = $1 ORDER BY station_name", systemID)
	if err != nil {
		return nil, fmt.Errorf("failed to query stations by system ID: %w", err)
	}
	defer rows.Close()
	var stations []models.Station
	for rows.Next() {
		var s models.Station
		if err := rows.Scan(&s.StationID, &s.StationName, &s.SystemID); err != nil {
			return nil, fmt.Errorf("failed to scan station row: %w", err)
		}
		stations = append(stations, s)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration for stations by system ID: %w", err)
	}
	return stations, nil
}

// GetSystemHeatmapByRegionMode queries per-period per-system metrics for a region.
// Returns region name and a slice ordered by period desc, kills desc.
func GetSystemHeatmapByRegionMode(regionID int, mode string) (string, []models.SystemPeriodHeatPoint, string, string, error) {
	db := GetDB()
	if db == nil {
		return "", nil, "", "", fmt.Errorf("database not initialized")
	}
	// Get interval
	interval, err := GetModeInterval(mode)
	if err != nil {
		return "", nil, "", "", fmt.Errorf("invalid mode: %w", err)
	}
        // Fetch region name
        var regionName string
        _ = db.QueryRow("SELECT region_name FROM regions WHERE region_id = $1", regionID).Scan(&regionName)
	// Get window
	var windowStart time.Time
	var windowEnd time.Time
	if err := db.QueryRow("SELECT (NOW() AT TIME ZONE 'UTC' - $1::interval) AS window_start, (NOW() AT TIME ZONE 'UTC') AS window_end", interval).Scan(&windowStart, &windowEnd); err != nil {
		return regionName, nil, "", "", fmt.Errorf("get window bounds: %w", err)
	}
	windowStartStr := windowStart.Format(time.RFC3339)
	windowEndStr := time.Now().UTC().Format(time.RFC3339)
	// query
	const query = `SELECT
		s.system_id,
		s.system_name,
		COUNT(k.killmail_id) AS kills,
		COALESCE(SUM(k.destroyed_value), 0) AS destroyed_value,
		COALESCE(SUM(k.dropped_value), 0) AS dropped_value
		FROM systems s
		JOIN constellations c ON s.constellation_id = c.constellation_id
		LEFT JOIN killmails k
		ON k.solar_system_id = s.system_id
		AND k.killmail_time >= (NOW() AT TIME ZONE 'UTC' - $2::interval)  -- sliding window
		WHERE c.region_id = $1
		GROUP BY s.system_id, s.system_name
		ORDER BY kills DESC;`
	// Check for errors
	rows, err := db.Query(query, regionID, interval)
	if err != nil {
		return regionName, nil, "", "", fmt.Errorf("query system heatmap by mode: %w", err)
	}
	defer rows.Close()
	// Run through the model
	heat := make([]models.SystemPeriodHeatPoint, 0, 128)
	/*for rows.Next() {
		var p models.SystemPeriodHeatPoint
		if err := rows.Scan(&p.SystemID, &p.SystemName, &p.Kills, &p.DestroyedValue, &p.DroppedValue); err != nil {
			return regionName, nil, fmt.Errorf("scan system heatmap row: %w", err)
		}
		heat = append(heat, p)
	}*/
	for rows.Next() {
	var id int
	var name string
	var kills int
	var destroyed, dropped float64
	// Check erors
	if err := rows.Scan(&id, &name, &kills, &destroyed, &dropped); err != nil {
		return regionName, nil, "", "", fmt.Errorf("scan system heatmap row: %w", err)
	}
	// Build heat
	heat = append(heat, models.SystemPeriodHeatPoint{
			SystemID:       id,
			SystemName:     name,
			Kills:          kills,
			DestroyedValue: destroyed,
			DroppedValue:   dropped,
		})
	}
	// Check for errors
	if err := rows.Err(); err != nil {
		return regionName, nil, "", "", fmt.Errorf("rows error: %w", err)
	}
	return regionName, heat, windowStartStr, windowEndStr, nil
}

// GetRecentKillmailsBySystemID returns the most recent 15 killmails for a given system.
func GetRecentKillmailsBySystemID(systemID int) ([]models.Killmails, error) {
	db := GetDB()
	// Build query
	query := `SELECT
		killmail_id,
		COALESCE(solar_system_id, 0) AS solar_system_id,
		killmail_time,
		COALESCE(destroyed_value, 0) AS destroyed_value,
		COALESCE(dropped_value, 0) AS dropped_value,
		COALESCE(killmail_hash, '') AS killmail_hash,
		COALESCE(total_value, 0) AS total_value,
		COALESCE(fitted_value, 0) AS fitted_value,
		COALESCE(victim_ship, 0) AS victim_ship,
		COALESCE(kill_ship, 0) AS kill_ship
		FROM killmails
		WHERE solar_system_id = $1
		ORDER BY killmail_time DESC
		LIMIT 15`
	rows, err := db.Query(query, systemID)
	if err != nil {
		return nil, fmt.Errorf("failed to query recent killmails by system: %w", err)
	}
	defer rows.Close()
	// Iterate over rows
	var kills = make([]models.Killmails, 0)
	for rows.Next() {
		var k models.Killmails
		if err := rows.Scan(&k.KillmailID, &k.SolarSystemID, &k.KillmailTime, &k.DestroyedValue, &k.DroppedValue, &k.KillmailHash, &k.TotalValue, &k.FittedValue, &k.VictimShip, &k.KillShip); err != nil {
			return nil, fmt.Errorf("failed to scan killmail row: %w", err)
		}
		kills = append(kills, k)
	}
	// Check for errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return kills, nil
}

// GetKillsBySystemID fetches kill counts grouped by time periods
func GetKillsBySystemID(systemID int, mode string) (string, []models.PeriodCount, error) {
	db := GetDB()
	// Fetch the system name using GetSystemNameByID
	systemName, err := GetSystemNameByID(systemID)
	if err != nil {
		return "", nil, fmt.Errorf("failed to fetch system name: %w", err)
	}
	// Query to group kills by time periods
	query := `SELECT DATE_TRUNC($2, killmail_time) AS period,
		COUNT(*) AS count,
		COALESCE(SUM(destroyed_value), 0) AS destroyed_value,
		COALESCE(SUM(dropped_value), 0) AS dropped_value
		FROM killmails
		WHERE solar_system_id = $1
		GROUP BY period
		ORDER BY period DESC`
	// Check for rows
	rows, err := db.Query(query, systemID, mode)
	if err != nil {
        	return systemName, nil, fmt.Errorf("failed to query kills: %w", err)
	}
	defer rows.Close()
    	// Process the results into a slice of PeriodCount
	var buckets []models.PeriodCount
	totalKills := 0
	for rows.Next() {
		var bucket models.PeriodCount
		if err := rows.Scan(&bucket.Period, &bucket.Count, &bucket.DestroyedValue, &bucket.DroppedValue); err != nil {
			return systemName, nil, fmt.Errorf("failed to scan bucket: %w", err)
		}
		totalKills += bucket.Count
		buckets = append(buckets, bucket)
	}
	// Check for any errors during row iteration
	if err := rows.Err(); err != nil {
		return systemName, nil, fmt.Errorf("row iteration error: %w", err)
	}
	return systemName, buckets, nil
}

// GetKillsByConstellationID fetches kill counts grouped by time periods
func GetKillsByConstellationID(constellationID int, mode string) (string, []models.PeriodCount, error) {
	db := GetDB()
	// Fetch the constellation name by using GetConstellationByIDorRegionID
	constellation, err := GetConstellationByIDOrRegionID(constellationID)
	if err != nil {
		return "", nil, fmt.Errorf("failed to fetch constellation name: %w", err)
	}
	// Check if the region was found (GetRegionByID might return nil, nil if not found)
	if constellation == nil {
		return "", nil, fmt.Errorf("region not found")
	}
	// Extract the name string from the struct
	constellationName := constellation[0].ConstellationName
	// Query to group kills by time periods by using JOIN 'killmails' with 'systems' to filter by constellation_id
	query := `SELECT DATE_TRUNC($2, k.killmail_time) AS period,
		COUNT(*) AS count,
		COALESCE(SUM(k.destroyed_value), 0) AS destroyed_value,
		COALESCE(SUM(k.dropped_value), 0) AS dropped_value
		FROM killmails k
		JOIN systems s ON k.solar_system_id = s.system_id
		WHERE s.constellation_id = $1
		GROUP BY period
		ORDER BY period DESC;`
	// Check for errors
	rows, err := db.Query(query, constellationID, mode)
	if err != nil {
		return constellationName, nil, fmt.Errorf("failed to query kills: %w", err)
	}
	defer rows.Close()
	// Process the results into a slice of PeriodCount
	var buckets []models.PeriodCount
	totalKills := 0
	for rows.Next() {
		var bucket models.PeriodCount
		if err := rows.Scan(&bucket.Period, &bucket.Count, &bucket.DestroyedValue, &bucket.DroppedValue); err != nil {
			return constellationName, nil, fmt.Errorf("failed to scan bucket: %w", err)
		}
		totalKills += bucket.Count
		buckets = append(buckets, bucket)
	}
	// Check for erros
	if err := rows.Err(); err != nil {
		return constellationName, nil, fmt.Errorf("row iteration error: %w", err)
	}
	return constellationName, buckets, nil
}

// GetKillsByRegionID fetches kill counts grouped by time periods
func GetKillsByRegionID(regionID int, mode string) (string, []models.PeriodCount, error) {
	db := GetDB()
	// Fetch the region name using GetRegionNameByID
	region, err := GetRegionByID(regionID)
	if err != nil {
		return "", nil, fmt.Errorf("failed to fetch region name: %w", err)
	}
	// Check if the region was found
	if region == nil {
		return "", nil, fmt.Errorf("region not found")
	}
	// Extract the name string from the struct
	regionName := region.RegionName
	// Query to group kills by time period using JOIN systems to filter killmails by the region_id of the systems
	query := `SELECT DATE_TRUNC($2, k.killmail_time) AS period,
		COUNT(*) AS count,
                COALESCE(SUM(k.destroyed_value), 0) AS destroyed_value,
                COALESCE(SUM(k.dropped_value), 0) AS dropped_value
		FROM killmails k
		JOIN systems s ON k.solar_system_id = s.system_id
		JOIN constellations c ON s.constellation_id = c.constellation_id
		WHERE c.region_id = $1
		GROUP BY period
		ORDER BY period DESC;`
	// Check for errors
	rows, err := db.Query(query, regionID, mode)
	if err != nil {
		return regionName, nil, fmt.Errorf("failed to query kills: %w", err)
	}
	defer rows.Close()
	// Process the results into a slice of PeriodCount
	var buckets []models.PeriodCount
	totalKills := 0
	for rows.Next() {
		var bucket models.PeriodCount
		if err := rows.Scan(&bucket.Period, &bucket.Count, &bucket.DestroyedValue, &bucket.DroppedValue); err != nil {
			return regionName, nil, fmt.Errorf("failed to scan bucket: %w", err)
		}
		totalKills += bucket.Count
		buckets = append(buckets, bucket)
	}
	// Check for any errors during row iteration
	if err := rows.Err(); err != nil {
		return regionName, nil, fmt.Errorf("row iteration error: %w", err)
	}
	return regionName, buckets, nil
}

// Get time slice window
func GetModeInterval(mode string) (string, error) {
	switch mode {
		case "hour":
			return "1 hour", nil
		case "day":
			return "1 day", nil
		case "week":
			return "1 week", nil
		case "month":
			return "1 month", nil
		default:
			return "", fmt.Errorf("invalid time window: %s", mode)
	}
}

// Get top regions by kill count
func GetTopRegionsByKills(mode string) ([]models.RegionKillCount, error) {
	db := GetDB()
	// Get the interval
	interval, err := GetModeInterval(mode)
	if err != nil {
		return nil, err
	}
	// Query for killmails
	query := fmt.Sprintf(`SELECT r.region_id, r.region_name, COUNT(*) as kill_count
		FROM killmails k
		JOIN systems s ON k.solar_system_id = s.system_id
		JOIN constellations c ON s.constellation_id = c.constellation_id
		JOIN regions r ON c.region_id = r.region_id
		WHERE k.killmail_time > NOW() - INTERVAL '%s'
		GROUP BY r.region_id, r.region_name
		ORDER BY kill_count DESC
		LIMIT 10`, interval)
	// Check for errors
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query top regions: %w", err)
	}
	defer rows.Close()
	// Models call
	var results []models.RegionKillCount
	for rows.Next() {
		var r models.RegionKillCount
		if err := rows.Scan(&r.RegionID, &r.RegionName, &r.TotalKills); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return results, nil
}

// Get top constellations by kill count
func GetTopConstellationsByKills(mode string) ([]models.ConstellationKillCount, error) {
	db := GetDB()
	// Get time interval
	interval, err := GetModeInterval(mode)
	if err != nil {
		return nil, err
	}
	// Qeury
	query := fmt.Sprintf(`SELECT c.constellation_id, c.constellation_name, COUNT(*) as kill_count
		FROM killmails k
		JOIN systems s ON k.solar_system_id = s.system_id
		JOIN constellations c ON s.constellation_id = c.constellation_id
		WHERE k.killmail_time > NOW() - INTERVAL '%s'
		GROUP BY c.constellation_id, c.constellation_name
		ORDER BY kill_count DESC
		LIMIT 10`, interval)
	// Chjeck for errors
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query top constellations: %w", err)
	}
	defer rows.Close()
	// Model calls
	var results []models.ConstellationKillCount
	for rows.Next() {
		var r models.ConstellationKillCount
		if err := rows.Scan(&r.ConstellationID, &r.ConstellationName, &r.TotalKills); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return results, nil
}

// Get top systems by kill count
func GetTopSystemsByKills(mode string) ([]models.SystemKillCount, error) {
	db := GetDB()
	// Get interval
	interval, err := GetModeInterval(mode)
	if err != nil {
		return nil, err
	}
	// Query
	query := fmt.Sprintf(`SELECT s.system_id, s.system_name, COUNT(*) as kill_count
		FROM killmails k
		JOIN systems s ON k.solar_system_id = s.system_id
		WHERE k.killmail_time > NOW() - INTERVAL '%s'
		GROUP BY s.system_id, s.system_name
		ORDER BY kill_count DESC
		LIMIT 10`, interval)
	// Check for errors
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query top systems: %w", err)
	}
	defer rows.Close()
	// Models call
	var results []models.SystemKillCount
	for rows.Next() {
		var r models.SystemKillCount
		if err := rows.Scan(&r.SystemID, &r.SystemName, &r.TotalKills); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return results, nil
}

