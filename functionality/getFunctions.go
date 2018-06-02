package functionality

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"errors"
	"os"
)

// Sources:
// Removing decimals:
// https://stackoverflow.com/questions/18390266/how-can-we-truncate-float64-type-to-a-particular-precision-in-golang


//GetADB ...
// creates a db connection
func GetADB() *MongoDBInfo {

	db := MongoDBInfo{
		os.Getenv("DB_HOST"),
		"botcurrency",
		"ratesCollection",
	}

	_, err := mgo.Dial(db.MongoURL)
	if err != nil {
		panic(err)
	}

	return &db
}

// GetBody ...
// @Return: Body and error/nil
func GetBody(url string, myClient *http.Client) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		body := []byte(nil)
		return body, err
	}

	// Apply header to request
	req.Header.Set("User-Agent", "AkselHj")

	// Try to execute request
	res, doError := myClient.Do(req)
	if doError != nil {
		body := []byte(nil)
		return body, doError
	}

	defer res.Body.Close()

	body, readError := ioutil.ReadAll(res.Body)

	if readError != nil {
		return body, readError
	}

	return body, nil
}

//GetEURExchange ...
func GetEURExchange(target string)(float64, error){
	localData, got := MongoDB.GetLocalFixer(0)
	if got == false{
		return 0, errors.New("DB failed at GetLocalFixer()")
	}
	rates := RawFixer{}

	byteData, err := json.Marshal(localData)
	if err != nil{
		return 0, err
	}
	err = json.Unmarshal(byteData, &rates)
	if err != nil{
		return 0, err
	}
	return rates.LocalRate[target], nil
}

//GetExchange ...
func GetExchange(rates Parameters)(float64, error){

	rawFixer := RawFixer{}
	exchangeValue := 0.0

	localFixer, got := MongoDB.GetLocalFixer(0)
	if got == false{
		return 0, errors.New("could not retrieve rates in db")
	}

	fixerData, err := json.Marshal(localFixer)
	if err != nil{
		return 0, err
	}

	err = json.Unmarshal(fixerData, &rawFixer)
	if err != nil{
		return 0, err
	}

	if rawFixer.LocalRate[rates.Currency1] != 0 {
		if rates.Currency2 != "EUR" {
			exchangeValue = rawFixer.LocalRate[rates.Currency2] / rawFixer.LocalRate[rates.Currency1]
		} else {
			exchangeValue = 1.0 / rawFixer.LocalRate[rates.Currency1]
		}
	}

	return exchangeValue, nil
}

//GetRate ...
func GetRates(myClient *http.Client, base string) (interface{}, error) {

	var fixerInterface interface{}

	fixerURL := "http://api.fixer.io/latest?base=" + base

	fixerBody, err := GetBody(fixerURL, myClient)
	if err != nil {
		return fixerInterface, err
	}

	err = json.Unmarshal(fixerBody, &fixerInterface)
	if err != nil {
		return fixerInterface, err
	}

	return fixerInterface, err

}
