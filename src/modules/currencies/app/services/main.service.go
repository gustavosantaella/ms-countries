package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"project/src/helpers"
	"strconv"
	"strings"
	"time"
)

func CurrencyConversion(base string, qoute string, amount string) (interface{}, error) {

	type CurrencyResponse struct {
		Base_currency  string      `json:"base_currency"`
		Quote_currency string      `json:"quote_currency"`
		Close_time     string      `json:"close_time"`
		Average_bid    string      `json:"average_bid"`
		Average_ask    string      `json:"average_ask"`
		High_bid       string      `json:"high_bid"`
		High_ask       string      `json:"high_ask"`
		Low_bid        string      `json:"low_bid"`
		Low_ask        string      `json:"low_ask"`
		Amount         interface{} `json:"amount"`
		Converted      interface{} `json:"converted"`
	}
	API_CURRENCY_URL, err := helpers.EnvGetProperty("CURRENCY_API")
	if err != nil {
		return nil, err
	}
	API_CURRENCY_URL = strings.ReplaceAll(API_CURRENCY_URL, "{BASE}", base)
	API_CURRENCY_URL = strings.ReplaceAll(API_CURRENCY_URL, "{QOUTE}", qoute)
	startDate := time.Now().Format("2006-01-02")
	endDate := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	API_CURRENCY_URL = strings.ReplaceAll(API_CURRENCY_URL, "{ENDDATE}", endDate)
	API_CURRENCY_URL = strings.ReplaceAll(API_CURRENCY_URL, "{STARTDATE}", startDate)
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
	var r2 []CurrencyResponse
	jsonData, err := json.Marshal(r["response"])
	if err != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(jsonData, &r2)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	if len(r2) == 0 {
		fmt.Println(API_CURRENCY_URL)
		return nil, errors.New("response is empty")
	}
	curremtR2 := r2[0]
	curremtR2.Amount = amount
	floatParsed, err := (strconv.ParseFloat(curremtR2.Average_ask, 64))
	floatParsedAmount, err2 := (strconv.ParseFloat(amount, 64))
	if err != nil || err2 != nil {
		return nil, err
	}
	curremtR2.Converted = floatParsed * floatParsedAmount
	return curremtR2, nil
}
