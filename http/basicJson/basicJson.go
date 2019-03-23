package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
{
	"location": "Zzyzx",
	"weather": "sunny",
	"temperature": 30,
	"celsius": true,
	"temp_forecast": [27, 25, 28],
	"wind": {
		"direction": "NW",
		"speed": 15
	}
}
*/

type weatherData struct {
	LocationName string   `json: locationName`
	Weather      string   `json: weather`
	Temperature  int      `json: temperature`
	Celsius      bool     `json: celsius`
	TempForecast []int    `json: temp_forecast`
	Wind         windData `json: wind`
}

type windData struct {
	Direction string `json: direction`
	Speed     int    `json: speed`
}

// data model for a client sending location request (the server will respose a weather data)
type loc struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// declare a location struct to store a location data send by client
	location := loc{}
	// the location data is inside the request body
	// we need a byte slice for unmarshalling (return by ReadAll function)
	// note: we use ReadAll for simplicity, it should be taken care for reading large data
	fmt.Println("isi body ", r.Body)
	jsn, err := ioutil.ReadAll(r.Body)
	fmt.Println("isi jsn ", jsn)
	if err != nil {
		log.Fatal("Error while reading r.Body: ", err)
	}
	// decode the recieve data using the Unmarshal function
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v ('&location'). If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	err = json.Unmarshal(jsn, &location)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}
	log.Printf("Received: %v\n", location)

	// use mock up data for response (we can get the data from weather service api or database or etc)
	weather := weatherData{
		LocationName: "Zzyzx",
		Weather:      "cloudy",
		Temperature:  31,
		Celsius:      true,
		TempForecast: []int{30, 32, 29},
		Wind: windData{
			Direction: "S",
			Speed:     28,
		},
	}

	// use Marshal for encoding go struct as json
	// Marshal returns the JSON encoding of v ('weather') in byte slice.
	weatherJson, err := json.Marshal(weather)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	// set the content type of response accordingly
	w.Header().Set("Content-Type", "application/json")
	// sending the response
	w.Write(weatherJson)
}

func server() {
	http.HandleFunc("/", weatherHandler)
	http.ListenAndServe(":8080", nil)
}

// mock up a client request
func client() {
	// create json data, marshall encode go struct to json
	locJson, err := json.Marshal(loc{Lat: 35.14325, Lon: -116.184})
	// set http request
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer(locJson))
	req.Header.Set("Content-Type", "application/json")
	// an http client will send the http request to the server and collect the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

func main() {
	go server()
	client()
}
