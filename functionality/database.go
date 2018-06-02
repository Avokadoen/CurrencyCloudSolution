package functionality

//Sources:
// https://github.com/marni/imt2681_cloud/blob/master/mongodb

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"time"
)

//MongoInterface ...
// All of the functionality for db
type MongoInterface interface {
	Init()
	CountFix() int
	GetLocalFixer(date string) (interface{}, bool)
	AddDailyFix()
}

var MongoDB *MongoDBInfo

//MongoDBInfo ...
// Data for the db is contained in this struct
type MongoDBInfo struct {
	MongoURL               	string
	DatabaseName           	string
	RatesCollection        	string
}

//Init ...
// Initializes the database
func (db *MongoDBInfo) Init() {
	session, err := mgo.Dial(db.MongoURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// make sure date remains unique
	index := mgo.Index{
		Key:        []string{"date"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DatabaseName).C(db.RatesCollection).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

//CountFix ...
func (db *MongoDBInfo) CountFix() int {
	session, err := mgo.Dial(db.MongoURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// handle to "db"
	count, err := session.DB(db.DatabaseName).C(db.RatesCollection).Count()
	if err != nil {
		log.Printf("error in Count(): %v", err.Error())
		return -1
	}

	return count
}

//GetLocalFixer ...
// Retrieve local fixer data based on index
func (db *MongoDBInfo) GetLocalFixer(offset int) (interface{}, bool) {

	session, err := mgo.Dial(db.MongoURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fixerQry := session.DB(db.DatabaseName).C(db.RatesCollection).Find(nil).Sort("-date")

	var ratesData interface{}
	gotData := true

	err = fixerQry.Skip(offset).One(&ratesData)
	if err != nil {
		gotData = false
	}

	return ratesData, gotData
}

//AddDailyFix ...
// Retrieves daily data from fixer
func (db *MongoDBInfo) AddDailyFix() error{
	myClient := http.Client{
		Timeout: time.Second * 10,
	}

	rates, err := GetRates(&myClient, "EUR")
	if err != nil {
		return err
	}

	session, err := mgo.Dial(db.MongoURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.RatesCollection).Insert(rates)
	if err != nil {
		return err
	}
	return nil
}
