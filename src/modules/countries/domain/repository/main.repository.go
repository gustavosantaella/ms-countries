package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"project/src/modules/countries/domain/entity"
)

func GetAll() []entity.CountryData {

	var countries []entity.CountryData
	pwd, _ := os.Getwd()
	countriesPath := filepath.Join(pwd, "src", "databases", "json", "countries.json")
	data, err := os.ReadFile(countriesPath)
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		return []entity.CountryData{}
	}

	if err := json.Unmarshal(data, &countries); err != nil {
		fmt.Println("Error parseando JSON:", err)
		return []entity.CountryData{}
	}

	return countries
}

func GetOneByIso3(iso3 string) entity.CountryData {

	var countries []entity.CountryData
	pwd, _ := os.Getwd()
	countriesPath := filepath.Join(pwd, "src", "databases", "json", "countries.json")
	data, err := os.ReadFile(countriesPath)
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		return entity.CountryData{}
	}

	if err := json.Unmarshal(data, &countries); err != nil {
		fmt.Println("Error parseando JSON:", err)
		return entity.CountryData{}
	}

	for _, country := range countries {
		if country.DialCode == iso3 {
			return country
		}
	}

	return entity.CountryData{}
}
