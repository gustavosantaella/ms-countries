package repository

import (
	"project/src/databases/supabase"
	"project/src/modules/countries/domain/entity"
)

func GetAll() []entity.Country {
	var countries []entity.Country

	supabase.DB.Model(&countries).Select()

	return countries
}

func GetOneByIso3(iso3 string) entity.Country {

	var country entity.Country
	supabase.DB.Model(&country).Where("iso3 = ?", iso3).Select()

	return country
}
