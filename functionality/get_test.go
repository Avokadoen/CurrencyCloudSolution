package functionality

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestGetBody(t *testing.T) {

	myClient := http.Client{
		Timeout: time.Second * 10,
	}

	testURL := "http://api.fixer.io/2017-04-09?base=EUR"

	rawBody := []byte(`{"base":"EUR","date":"2017-04-07","rates":{"AUD":1.4123,"BGN":1.9558,"BRL":3.3349,"CAD":1.4256,
	"CHF":1.0695,"CNY":7.3318,"CZK":26.563,"DKK":7.4363,"GBP":0.85573,"HKD":8.2596,"HRK":7.449,
	"HUF":310.36,"IDR":14161.0,"ILS":3.8784,"INR":68.378,"JPY":117.64,"KRW":1206.4,"MXN":19.947,
	"MYR":4.7144,"NOK":9.1583,"NZD":1.5249,"PHP":53.019,"PLN":4.2249,"RON":4.5178,
	"RUB":60.492,"SEK":9.5963,"SGD":1.4902,"THB":36.79,"TRY":3.9726,"USD":1.063,"ZAR":14.684}}`)

	getBodyBody, err := GetBody(testURL, &myClient)
	if err != nil {
		t.Error(err)
		return
	}

	testJson := RawFixer{}
	getBodyJson := RawFixer{}

	json.Unmarshal(rawBody, &testJson)
	json.Unmarshal(getBodyBody, &getBodyJson)

	if testJson.Date != getBodyJson.Date || testJson.LocalRate["NOK"] != getBodyJson.LocalRate["NOK"] {
		t.Error()
	}
}

func TestGetEURExchange(t *testing.T) {
	value, err := GetEURExchange("NOK")
	if err != nil{
		t.Error(err.Error())
	}else if value <= 0{
		t.Error("failed to get exchange value")
	}
}

func TestGetExchange(t *testing.T) {
	testInn := Parameters{}
	testInn.Currency1 = "NOK"
	testInn.Currency2 = "EUR"

	value, err := GetExchange(testInn)
	if err != nil {
		t.Error(err.Error())
	}else if value <= 0{
		t.Error("failed to get exchange value")
	}

	testInn.Currency1 = "NOK"
	testInn.Currency2 = "DKK"

	value, err = GetExchange(testInn)
	if err != nil {
		t.Error(err.Error())
	}else if value <= 0{
		t.Error("failed to get exchange value")
	}
}