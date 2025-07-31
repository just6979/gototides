package main

// Tide corresponds to the Tide schema in the OpenAPI spec.
type Tide struct {
	Date    string  `json:"date"`
	Day     string  `json:"day"`
	Height  float64 `json:"height"`
	ISODate string  `json:"iso-date"`
	Prior   string  `json:"prior"`
	Time    string  `json:"time"`
	Type    string  `json:"type"`
}

// TidesResponse corresponds to the TidesResponse schema.
type TidesResponse struct {
	ReqLat        float64 `json:"req_lat"`
	ReqLon        float64 `json:"req_lon"`
	ReqTimestamp  string  `json:"req_timestamp"`
	RespLat       float64 `json:"resp_lat"`
	RespLon       float64 `json:"resp_lon"`
	Station       string  `json:"station"`
	StationTZ     string  `json:"station_tz"`
	Status        int     `json:"status"`
	WTI_copyright string  `json:"wti_copyright,omitempty"`
	Tides         []Tide  `json:"tides"`
}

// Station corresponds to the Station schema.
type Station struct {
	ID    string  `json:"id"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
	Name  string  `json:"name"`
	NOAA  bool    `json:"noaa"`
	Org   string  `json:"org"`
	OrgID string  `json:"org_id"`
}

// StationResponse corresponds to the StationResponse schema.
// Note: The OpenAPI spec example shows "station" as an array,
// but the schema defines it as a single object. This implementation
// follows the schema definition.
type StationResponse struct {
	Status  string  `json:"status"`
	Station Station `json:"station"`
}

// StationsResponse corresponds to the StationsResponse schema.
type StationsResponse struct {
	Status       string    `json:"status"`
	StationCount int       `json:"station_count"`
	Stations     []Station `json:"stations"`
}

type WorldTidesExtremesResponse struct {
	Status        int     `json:"status"`
	CallCount     int     `json:"callCount"`
	Copyright     string  `json:"copyright"`
	RequestLat    float64 `json:"requestLat"`
	RequestLon    float64 `json:"requestLon"`
	ResponseLat   float64 `json:"responseLat"`
	ResponseLon   float64 `json:"responseLon"`
	Atlas         string  `json:"atlas"`
	Station       string  `json:"station"`
	Timezone      string  `json:"timezone"`
	RequestDatum  string  `json:"requestDatum"`
	ResponseDatum string  `json:"responseDatum"`
	Extremes      []struct {
		Dt     int     `json:"dt"`
		Date   string  `json:"date"`
		Height float64 `json:"height"`
		Type   string  `json:"type"`
	} `json:"extremes"`
}

// --- Mock Data for Handlers ---

var mockStation = Station{
	ID:    "NOAA:8441241",
	Lat:   42.7101,
	Lon:   -70.7886,
	Name:  "Plum Island Sound (south end), Massachusetts",
	NOAA:  true,
	Org:   "NOAA",
	OrgID: "8441241",
}

var mockTide = Tide{
	Date:    "06-11",
	Day:     "Wed",
	Height:  0.483,
	ISODate: "2025-06-11T18:31:50-04:00",
	Prior:   "prior",
	Time:    "18:31",
	Type:    "Low",
}

var mockTidesResponse = TidesResponse{
	ReqLat:       42.6724,
	ReqLon:       -70.9443,
	ReqTimestamp: "Wed, 11 Jun 2025 10:08:36 EDT",
	RespLat:      42.7101,
	RespLon:      -70.7886,
	Station:      "Plum Island Sound (south end), Massachusetts",
	StationTZ:    "America/New_York",
	Status:       200,
	Tides:        []Tide{mockTide},
}

var mockStationResponse = StationResponse{
	Status:  "OK",
	Station: mockStation,
}

var mockStationsResponse = StationsResponse{
	Status:       "OK",
	StationCount: 1,
	Stations:     []Station{mockStation},
}
