package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/src/helpers"
	"strconv"
	"strings"
)

func CurrencyConversion(base string, qoute string, amount string) (interface{}, error) {

	type CurrencyResponse struct {
		Avg_bid string `json:"avg_bid"`
		Avg_ask string `json:"avg_ask"`
		Avg_mid string `json:"avg_mid"`
	}

	type ResponseConverted struct {
		Base      string  `json:"base"`
		Rate      string  `json:"rate"`
		Qoute     string  `json:"qoute"`
		Amount    float64 `json:"amount"`
		Converted float64 `json:"converted"`
	}
	API_CURRENCY_URL, err := helpers.EnvGetProperty("CURRENCY_API_RATES")
	if err != nil {
		return nil, err
	}
	API_CURRENCY_URL = strings.ReplaceAll(API_CURRENCY_URL, "{BASE}", base)
	API_CURRENCY_URL = strings.ReplaceAll(API_CURRENCY_URL, "{QOUTE}", qoute)
	request, err := http.NewRequest("GET", API_CURRENCY_URL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	var r map[string]interface{}
	json.NewDecoder(response.Body).Decode(&r)
	closeError := response.Body.Close()
	if closeError != nil {
		return nil, closeError
	}
	jsonData, err := json.Marshal(r["rates"])
	if err != nil {
		return nil, err
	}
	r2 := CurrencyResponse{}
	unmarshalError := json.Unmarshal(jsonData, &r2)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	// if len(r2) == 0 {
	// 	fmt.Println(API_CURRENCY_URL)
	// 	return nil, errors.New("response is empty")
	// }
	curremtR2 := r2
	fmt.Println("before float conversion")
	floatParsed, err := (strconv.ParseFloat(curremtR2.Avg_ask, 64))
	floatParsedAmount, err2 := (strconv.ParseFloat(amount, 64))
	if err != nil || err2 != nil {
		fmt.Println(err, err2)
		return nil, err
	}
	fmt.Println(floatParsed, floatParsedAmount)
	converted := floatParsed * floatParsedAmount
	convertedResponse := ResponseConverted{
		Rate:      curremtR2.Avg_ask,
		Base:      base,
		Qoute:     qoute,
		Amount:    floatParsedAmount,
		Converted: converted,
	}
	return convertedResponse, nil
}

func AllCurrencies() (interface{}, error) {
	API_CURRENCY_URL, err := helpers.EnvGetProperty("CURRENCY_API_CURRENCIES")
	if err != nil {
		return nil, err
	}

	response, err := http.Get(API_CURRENCY_URL)
	if err != nil {
		fmt.Println("error request", err)
		return nil, err
	}

	var allCurrencies map[string]string
	err1 := json.NewDecoder(response.Body).Decode(&allCurrencies)
	if err1 != nil {
		fmt.Println("error1", err1)
		return nil, err1
	}

	return allCurrencies, nil
}
