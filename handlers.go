package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// respondWithJSON is a helper to write JSON responses.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getTidesByLocation(w http.ResponseWriter, r *http.Request) {
	// Example path: /json/tides/by-location/42.665,-70.9119
	location := strings.TrimPrefix(r.URL.Path, "/json/tides/by-location/")
	log.Printf("Request for tides by location: %s", location)
	// TODO: Implement logic to find station by location and get tides
	respondWithJSON(w, http.StatusOK, mockTidesResponse)
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
