package functionality

import (
	"encoding/json"
	"net/http"
	"log"
	"strings"
)

//DialogFlowHandler ...
func DialogFlowHandler(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {

		r.Header.Set("Content-Type", "application/json")

		dialogData := DialogFlowDataInn{}

		err := json.NewDecoder(r.Body).Decode(&dialogData)
		if err != nil {
			log.Print(err.Error())
			return
		}
		// Declare an empty response
		responseBody := []byte(nil)

		response, err := CreateExchangeResponse(dialogData.ResultData.Para)
		if err == nil {
			// Define response if CreateExchangeResponse was successful
			responseBody = []byte(`{"speech":` + response + `, "textDisplay": ` + response + `}`)
		}
		defer r.Body.Close()

		// If body was defined, present response content
		if responseBody != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseBody)
		}else{
			http.Error(w, "invalid request", http.StatusNoContent)
		}

	} else {
		http.Error(w, "invalid method used", http.StatusMethodNotAllowed)
	}

}

//URLHandler ...
// used to identify what user is requesting
func URLHandler(w http.ResponseWriter, r *http.Request) {
	URLCheck := r.URL.Path
	URLSplit := strings.Split(URLCheck, "/")

	checkString := strings.ToLower(URLSplit[1])

	if checkString == "dialogflow" {
		DialogFlowHandler(w, r)
	} else{
		http.NotFound(w, r)
	}

}