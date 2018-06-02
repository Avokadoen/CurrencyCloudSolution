package functionality

import "testing"

func TestCreateExchangeResponse(t *testing.T) {
	testValue := Parameters{}
	testValue.Currency1 = "illegal"
	testValue.Currency2 = "illegal"

	testString, err := CreateExchangeResponse(testValue)
	if err == nil || testString != ""{
		t.Error(err.Error())
	}

	testValue.Currency1 = "NOK"
	testValue.Currency2 = "NOK"
	testString, err = CreateExchangeResponse(testValue)
	if err != nil || testString == ""{
		t.Error("denied legal input")
	}else if testString != `"The exchange rate between NOK and NOK is 1"`{
		t.Error("wrong print on input")
	}
}