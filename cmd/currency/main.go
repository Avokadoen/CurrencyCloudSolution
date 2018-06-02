// Author: Aksel Hjerpbakk
// Studentnr: 997816

package main

import (
	"bitbucket.org/Avokadoen/cloud-assignment3-docker/functionality"
	"net/http"
	"os"
)

//main ...
// For serving dialogflow and users
func main() {
	functionality.MongoDB = functionality.GetADB()
	functionality.MongoDB.Init()
	functionality.MongoDB.AddDailyFix()
	port := os.Getenv("PORT")
	http.HandleFunc("/", functionality.URLHandler)
	http.ListenAndServe(":"+port, nil)
}
