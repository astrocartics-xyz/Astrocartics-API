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

func GetAllStargates() ([]models.Stargate, error) {
	return dba.GetAllStargates()
}

func GetStargateBySystemID(id int) ([]models.Stargate, error) {
	return dba.GetStargateBySystemID(id)
}

func GetSpectralClassCounts() ([]models.SpectralClassCount, error) {
	return dba.GetSpectralClassCounts()
} 