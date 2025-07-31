package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/joho/godotenv"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("templates", "index.mustache")
	http.ServeFile(w, r, fp)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()

	// Static file routes
	mux.HandleFunc("/", getIndex)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// API routes
	mux.HandleFunc("/json/tides/by-location/", getTidesByLocation)
	mux.HandleFunc("/json/tides/by-station/", getTidesByStation)
	mux.HandleFunc("/json/stations", getStations)
	mux.HandleFunc("/json/stations/refresh", refreshStations)
	mux.HandleFunc("/json/station/by-id/", getStationByID)
	mux.HandleFunc("/json/station/by-nearest/", getStationByNearest)

	port := ":5000"
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
