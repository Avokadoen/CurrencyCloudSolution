package functionality

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"bytes"
)

/*	Sources:
// https://elithrar.github.io/article/testing-http-handlers-go/

// Bytearray to string:
// https://stackoverflow.com/a/14230206
*/


func TestURLHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(URLHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound )
	}

	expected := "404 page not found"
	gotString := strings.TrimSpace(rr.Body.String())
	if  gotString != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			gotString, expected)
	}
}

func TestDialogFlowHandler(t *testing.T) {

	// TEST GET METHOD -- EXPECT INVALID
	req, err := http.NewRequest("GET", "/dialogflow", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr1 := httptest.NewRecorder()
	handler1 := http.HandlerFunc(DialogFlowHandler)

	handler1.ServeHTTP(rr1, req)

	if status := rr1.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed )
	}

	// TEST POST METHOD -- EXPECT VALID
	byteDialog := []byte(`{"result":{"parameters":{"currency1": "NOK", "currency2": "EUR"}}}`)
	ioDialog := bytes.NewReader(byteDialog)

	req, err = http.NewRequest("POST", "/dialogflow", ioDialog)
	if err != nil {
		t.Fatal(err)
	}

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(DialogFlowHandler)

	handler2.ServeHTTP(rr2, req)

	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK )
	}
	expected := `{"speech":"The exchange rate between NOK and EUR is"`
	gotString := rr2.Body.String()
	if  strings.Contains(expected, gotString){
		t.Errorf("handler returned unexpected body: got %v want %v",
			gotString, expected)
	}
}