package service

import (
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

func GetSystemsByRegionID(id int) ([]models.System, error) {
	return dba.GetSystemsByRegionID(id)
}

func GetAllStargates() ([]models.Stargate, error) {
	return dba.GetAllStargates()
}

func GetStargateBySystemID(id int) ([]models.Stargate, error) {
	return dba.GetStargateBySystemID(id)
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