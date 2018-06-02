package functionality

import "gopkg.in/mgo.v2/bson"

type CurrencyLoad struct {
	Id              bson.ObjectId	`bson:"_id,omitempty"`
	WebHookURL      string       	`json:"webhookURL"`
	BaseCurrency    string       	`json:"baseCurrency"`
	TargetCurrency  string       	`json:"targetCurrency"`
	MinTriggerValue float64      	`json:"minTriggerValue"`
	MaxTriggerValue float64       	`json:"maxTriggerValue"`
}

type Rates struct {
	BaseCurrency   	string 			`json:"baseCurrency"`
	TargetCurrency 	string 			`json:"targetCurrency"`
}


type RawFixer struct {
	Base      		string          `json:"base"`
	Date      		string          `json:"date"`
	LocalRate 	map[string]float64 	`json:"rates"`
}


// 				Dialog flow data

type DialogFlowDataInn struct{
	ResultData		Result			`json:"result"`
}

type Result struct{
	Para 			Parameters		`json:"parameters"`
}

type Parameters struct{
	Currency1		string			`json:"currency1"`
	Currency2		string			`json:"currency2"`
}


//-------------------------------------------------


