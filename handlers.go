package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// helpers

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func startResponse() time.Time {
	startTime := time.Now().UTC()
	log.Printf("Starting response at: %s\n", startTime.Format("2006-01-02 15:04:05.000000000 -0700 EDT"))
	return startTime
}

func sendOKResponse(w http.ResponseWriter, startTime time.Time, payload interface{}) {
	sendResponse(w, http.StatusOK, startTime, payload)
}

func sendErrorResponse(w http.ResponseWriter, startTime time.Time, err error) {
	log.Println("Error:", err)
	sendResponse(w, http.StatusInternalServerError, startTime, errorResponse{
		Status:  "error",
		Message: err.Error()})
}

func sendResponse(w http.ResponseWriter, code int, startTime time.Time, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v\n", err)
		code = http.StatusInternalServerError
	}

	endTime := time.Now().UTC()
	log.Printf("Response prepared at: %s\n", endTime.Format("2006-01-02 15:04:05.000000000 -0700 EDT"))
	log.Printf("Total processing time: %s\n", endTime.Sub(startTime).String())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// API handlers

func getTidesByLocation(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/tides/by-location/42.665,-70.9119

	startTime := startResponse()

	location := strings.TrimPrefix(r.URL.Path, "/json/tides/by-location/")
	loc := strings.Split(location, ",")
	log.Printf("Request for tides by location: (%s, %s)\n", loc[0], loc[1])

	requestBase := "https://www.worldtides.info/api/v3?extremes&localtime&datum=CD&lat=%s&lon=%s&key=%s"
	requestUrl := fmt.Sprintf(requestBase, loc[0], loc[1], os.Getenv("WORLDTIDES_INFO_API_KEY"))
	log.Printf("Fetching tide data from %s\n", requestUrl)
	wtiResponse, err := http.Get(requestUrl)
	if err != nil {
		sendErrorResponse(w, startTime, fmt.Errorf("unable to fetch tides for %s: %w", location, err))
		return
	}
	defer wtiResponse.Body.Close()

	wtiJson, _ := io.ReadAll(wtiResponse.Body)
	var wtiData WorldTidesExtremesResponse
	if err := json.Unmarshal(wtiJson, &wtiData); err != nil {
		log.Print("Can not unmarshal JSON\n")
	}

	// wtiString, _ := json.MarshalIndent(wtiData, "", "  ")
	// log.Println(string(wtiString))

	tz, err := time.LoadLocation(wtiData.Timezone)
	if err != nil {
		tz = time.UTC
	}
	localNow := startTime.In(tz)
	log.Printf("Local Now: %s\n", localNow.Format("2006-01-02 15:04:05 -0700 EDT"))

	var tides []Tide
	for _, extreme := range wtiData.Extremes {
		prior := "future"
		tideTime, _ := time.Parse("2006-01-02T15:04:05-07:00", extreme.Date)
		if tideTime.Before(localNow) {
			prior = "prior"
		}

		var newTide Tide
		newTide.ISODate = extreme.Date
		newTide.Type = extreme.Type
		newTide.Height = extreme.Height
		newTide.Date = tideTime.Format("01-02")
		newTide.Day = tideTime.Format("Mon")
		newTide.Time = tideTime.Format("15:04")
		newTide.Prior = prior

		tides = append(tides, newTide)
	}

	var tidesResponse TidesResponse
	tidesResponse.Tides = tides
	tidesResponse.Status = wtiData.Status
	tidesResponse.Station = wtiData.Station
	tidesResponse.StationTZ = wtiData.Timezone
	tidesResponse.ReqLat = wtiData.RequestLat
	tidesResponse.ReqLon = wtiData.RequestLon
	tidesResponse.RespLat = wtiData.ResponseLat
	tidesResponse.RespLon = wtiData.ResponseLon
	tidesResponse.WTI_copyright = wtiData.Copyright
	tidesResponse.ReqTimestamp = localNow.Format("Mon, 2 Jan 2006 15:04:05 EDT")

	sendOKResponse(w, startTime, tidesResponse)
}

func getTidesByStation(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/tides/by-station/NOAA:8440452
	startTime := startResponse()

	station := strings.TrimPrefix(r.URL.Path, "/json/tides/by-station/")
	log.Printf("Request for tides by station: %s", station)
	// TODO: Implement logic to get tides for the given station

	sendOKResponse(w, startTime, mockTidesResponse)
}

func getStations(w http.ResponseWriter, r *http.Request) {
	startTime := startResponse()

	log.Printf("Request for all stations")
	// TODO: Implement logic to retrieve all cached stations

	sendOKResponse(w, startTime, mockStationsResponse)
}

func refreshStations(w http.ResponseWriter, r *http.Request) {
	startTime := startResponse()

	log.Printf("Request to refresh stations cache")
	// TODO: Implement logic to refresh the cache and return new stations

	sendOKResponse(w, startTime, mockStationsResponse)
}

func getStationByID(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/station/by-id/NOAA:8441241
	startTime := startResponse()

	stationID := strings.TrimPrefix(r.URL.Path, "/json/station/by-id/")
	log.Printf("Request for station by ID: %s", stationID)
	// TODO: Implement logic to find a station by its ID
	response := mockStationResponse
	response.Station.ID = stationID // Echo back the requested ID

	sendOKResponse(w, startTime, response)
}

func getStationByNearest(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/station/by-nearest/42.665,-70.9119
	startTime := startResponse()

	location := strings.TrimPrefix(r.URL.Path, "/json/station/by-nearest/")
	log.Printf("Request for nearest station to location: %s", location)
	// TODO: Implement logic to find the nearest station

	sendOKResponse(w, startTime, mockStationResponse)
}
