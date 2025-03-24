package services

import (
	"project/src/modules/countries/domain/entity"
	"project/src/modules/countries/domain/repository"
)

func GetAll() []entity.CountryData {

	countries := repository.GetAll()
	return countries
}

func GetOneByIso3(iso3 string) entity.CountryData {
	return repository.GetOneByIso3(iso3)

}
