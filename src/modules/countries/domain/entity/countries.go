package entity

type Country struct {
	Id              int8    `json:"id"`
	Name            string  `json:"name"`
	Iso3            string  `json:"iso3"`
	Iso2            string  `json:"iso2"`
	Numeric_code    int8    `json:"numeric_code"`
	Phone_code      string  `json:"phone_code"`
	Capital         string  `json:"capital"`
	Currency        string  `json:"currency"`
	Currency_name   string  `json:"currency_name"`
	Currency_symbol string  `json:"currency_symbol"`
	Tld             string  `json:"tld"`
	Native          string  `json:"native"`
	Region          string  `json:"region"`
	Subregion       string  `json:"subregion"`
	Timezones       string  `json:"timezones"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Emoji           string  `json:"emoji"`
}

type CountryData struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	DialCode string `json:"dialCode"`
}
