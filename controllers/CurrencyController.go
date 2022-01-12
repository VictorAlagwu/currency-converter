package controllers

import (
	"converter/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Currency struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Symbol string `json:"symbol"`
}

type Converter struct {
	Amount string
	FromCurrency string
	ToCurrency string
}

func JSONHandler()([]Currency, error) {
	currencies := make([]Currency, 0)

	jsonFile, err := ioutil.ReadFile("currency.json")

	if err != nil {
		fmt.Println(err)
	}

	currentJson := []byte(jsonFile)
	err = json.Unmarshal(currentJson, &currencies)
	if err != nil {
		panic(err)
	}

	return currencies, nil
}

func GetRate(w http.ResponseWriter, r *http.Request)  {
	inputs := Converter {
		Amount: r.FormValue("amount"),
		FromCurrency: r.FormValue("fromCurrency"),
		ToCurrency: r.FormValue("toCurrency"),
	}
	amount, err := strconv.ParseFloat(inputs.Amount, 64)
	var currentValue float64 = 0.00000
	if err != nil {
		log.Println(err)
	}
	//Rate as of Jan 11, 2022, 19:32 UTC
	if inputs.FromCurrency == inputs.ToCurrency {
		currentValue = amount
	}

	if (inputs.FromCurrency == "NGN") && (inputs.ToCurrency == "GHS") {
		currentValue = amount * 0.0149
	}

	if (inputs.FromCurrency == "NGN") && (inputs.ToCurrency == "KSH") {
		currentValue = amount * 0.2742
	}

	if (inputs.FromCurrency == "GHS") && (inputs.ToCurrency == "NGN") {
		currentValue = amount * 66.9611
	}

	if (inputs.FromCurrency == "GHS") && (inputs.ToCurrency == "KSH") {
		currentValue = amount * 18.3562
	}

	if (inputs.FromCurrency == "KSH") && (inputs.ToCurrency == "NGN") {
		currentValue = amount * 3.6475
	}

	if (inputs.FromCurrency == "KSH") && (inputs.ToCurrency == "GHS") {
		currentValue = amount * 0.0545
	}
	responses.JSON(w, http.StatusOK, currentValue)
}



