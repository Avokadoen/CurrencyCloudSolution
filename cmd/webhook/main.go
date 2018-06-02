package main

import (
	"bitbucket.org/Avokadoen/cloud-assignment3-docker/functionality"
)

//main ...
// For daily update of DB
func main() {
	functionality.MongoDB = functionality.GetADB()
	functionality.MongoDB.Init()
	functionality.MongoDB.AddDailyFix()
}
