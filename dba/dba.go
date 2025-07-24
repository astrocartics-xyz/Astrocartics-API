package dba

import (
	"database/sql"
	"fmt"

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
	rows, err := db.Query("SELECT system_id, system_name, security_status, security_class, x_pos, y_pos, z_pos, constellation_id, spectral_class FROM systems ORDER BY system_name")
	if err != nil {
		return nil, fmt.Errorf("failed to query systems: %w", err)
	}
	defer rows.Close()
	var systems []models.System
	for rows.Next() {
		var s models.System
		if err := rows.Scan(&s.SystemID, &s.SystemName, &s.SecurityStatus, &s.SecurityClass, &s.XPos, &s.YPos, &s.ZPos, &s.ConstellationID, &s.SpectralClass); err != nil {
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
	sqlQuery := "SELECT system_id, system_name, security_status, security_class, x_pos, y_pos, z_pos, constellation_id, spectral_class FROM	systems WHERE system_id = $1 OR constellation_id = $1 ORDER BY system_name"
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

// GetSystemByName fetches a single system by its name.
func GetSystemByName(name string) (*models.System, error) {
	db := GetDB()
	var s models.System // Declare a variable of type System to hold the fetched data
	// Modify the SQL query to filter by system_name instead of system_id.
	// Use $1 as a placeholder for the name parameter.
	err := db.QueryRow("SELECT system_id, system_name, security_status, security_class, x_pos, y_pos, z_pos, constellation_id, spectral_class FROM systems WHERE system_name = $1", name).
		Scan(&s.SystemID, &s.SystemName, &s.SecurityStatus, &s.SecurityClass, &s.XPos, &s.YPos, &s.ZPos, &s.ConstellationID, &s.SpectralClass)
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

// GetStargateBySystemID fetches all stargates associated with a given system ID.
func GetStargateBySystemID(systemID int) ([]models.Stargate, error) {
	db := GetDB()
	// Prepare the SQL query with a WHERE clause to filter by system_id.
	// $1 is a placeholder for the first parameter passed to Query.
	rows, err := db.Query("SELECT stargate_id, stargate_name, system_id, destination_stargate_id, destination_system_id FROM stargates WHERE system_id = $1 ORDER BY stargate_name", systemID)
	if err != nil {
		// Return an error if the query itself fails (e.g., database connection issues, syntax error).
		return nil, fmt.Errorf("failed to query stargates by system ID: %w", err)
	}
	defer rows.Close() // Ensure rows are closed after the function returns
	var stargates []models.Stargate // Slice to hold the fetched stargate objects
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