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

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// respondWithJSON is a helper to write JSON responses.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getTidesByLocation(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/tides/by-location/42.665,-70.9119

	utcNow := time.Now().UTC()
	log.Printf("UTC Now: %s\n", utcNow)

	location := strings.TrimPrefix(r.URL.Path, "/json/tides/by-location/")
	loc := strings.Split(location, ",")
	log.Printf("Request for tides by location: (%s, %s)\n", loc[0], loc[1])

	requestBase := "https://www.worldtides.info/api/v3?extremes&localtime&datum=CD&lat=%s&lon=%s&key=%s"
	requestUrl := fmt.Sprintf(requestBase, loc[0], loc[1], os.Getenv("WORLDTIDES_INFO_API_KEY"))
	log.Printf("Fetching tide data from %s\n", requestUrl)
	wtiResponse, err := http.Get(requestUrl)
	if err != nil {
		respondWithJSON(w, http.StatusOK, errorResponse{"Error", "unable to fetch tides for " + location})
	}
	wtiJson, _ := io.ReadAll(wtiResponse.Body)
	var wtiData WorldTidesExtremesResponse
	if err := json.Unmarshal(wtiJson, &wtiData); err != nil {
		log.Print("Can not unmarshal JSON\n")
	}
	log.Println(wtiData)

	tz := wtiData.Timezone
	localNow := utcNow.In(time.FixedZone(tz, 0))
	log.Printf("Local Now: %s\n", localNow.Format("Mon, 2 Jan 2006 15:04:05 MST"))

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
	tidesResponse.ReqTimestamp = localNow.Format("Mon, 2 Jan 2005 15:04:05 MST")

	respondWithJSON(w, http.StatusOK, tidesResponse)
}

func getTidesByStation(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/tides/by-station/NOAA:8440452
	station := strings.TrimPrefix(r.URL.Path, "/json/tides/by-station/")
	log.Printf("Request for tides by station: %s", station)
	// TODO: Implement logic to get tides for the given station
	respondWithJSON(w, http.StatusOK, mockTidesResponse)
}

func getStations(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request for all stations")
	// TODO: Implement logic to retrieve all cached stations
	respondWithJSON(w, http.StatusOK, mockStationsResponse)
}

func refreshStations(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to refresh stations cache")
	// TODO: Implement logic to refresh the cache and return new stations
	respondWithJSON(w, http.StatusOK, mockStationsResponse)
}

func getStationByID(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/station/by-id/NOAA:8441241
	stationID := strings.TrimPrefix(r.URL.Path, "/json/station/by-id/")
	log.Printf("Request for station by ID: %s", stationID)
	// TODO: Implement logic to find a station by its ID
	response := mockStationResponse
	response.Station.ID = stationID // Echo back the requested ID
	respondWithJSON(w, http.StatusOK, response)
}

func getStationByNearest(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/station/by-nearest/42.665,-70.9119
	location := strings.TrimPrefix(r.URL.Path, "/json/station/by-nearest/")
	log.Printf("Request for nearest station to location: %s", location)
	// TODO: Implement logic to find the nearest station
	respondWithJSON(w, http.StatusOK, mockStationResponse)
}
