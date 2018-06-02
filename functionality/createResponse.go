package functionality

import (
	"fmt"
	"regexp"
	"errors"
)

//CreateExchangeResponse ...
// @return: if error: empty string and error. else: The string presented by dialogflowbot
func CreateExchangeResponse(usrRequest Parameters) (string, error){

	// Check if user ask for a valid currency
	validString, err := regexp.MatchString("^[A-Za-z]{3}$", usrRequest.Currency1)
	if err == nil && validString == true {
		validString, err = regexp.MatchString("^[A-Za-z]{3}$", usrRequest.Currency2)
		if err == nil && validString == true {
			exchangeValue := 1.0

			if usrRequest.Currency1 != usrRequest.Currency2 {
				Exchange := 0.0
				err := error(nil)

				// calculating exchange has three formulas, one is in GetEURExchange
				if usrRequest.Currency1 == "EUR" {
					Exchange, err = GetEURExchange(usrRequest.Currency2)
					if err != nil {
						return "", err
					}
				} else {
					Exchange, err = GetExchange(usrRequest)
					if err != nil {
						return "", err
					}
				}
				// Round the exchange value
				num := Exchange + 0.0005
				exchangeValue = float64(int64(1000*(num))) / 1000
			}
			a := fmt.Sprint(exchangeValue)

			return `"` + "The exchange rate between " + usrRequest.Currency1 + " and " + usrRequest.Currency2 + " is " + a + `"`, nil
		}
	}
	return "", errors.New("invalid input")
}
