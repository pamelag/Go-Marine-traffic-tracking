package vessel

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"lab.identitii.com/identitii/discord/model"
	"fmt"
)

type Configuration struct {
	Imo            string
	LocationApiKey string
	PortsApiKey    string
	Timespan       int
}

var vesselLocation model.VesselLocation = model.VesselLocation{}
var configuration Configuration = Configuration{}

// This tests GET request with passing in a parameter.
func TestGetLocation(t *testing.T) {
	// Port models a Port
	t.Log("TestGetParams ")

	vesselLocations := make([]model.VesselLocation, 0)
	vesselLocations = append(vesselLocations, vesselLocation)
	imo, _ := strconv.Atoi(configuration.Imo)
	vesselLocationResponse, err := GetVesselLocation(imo, configuration.LocationApiKey, configuration.PortsApiKey, configuration.Timespan)

	if err != nil {
		t.Error("Test failed")
	}

	if compare(vesselLocation, vesselLocationResponse.VesselLocations[0]) {
		t.Log("Passed ")
	}
}

func setUp() {
	ports := make([]model.Port, 0)

	vesselLocation = model.VesselLocation{
		MMSI:        "229085000",
		LAT:         "36.078680",
		LON:         "-4.865917",
		SPEED:       "62",
		COURSE:      "263",
		TIMESTAMP:   "2017-01-30T05:14:56",
		SHIPNAME:    "MINERVA LYDIA",
		SHIPTYPE:    "80",
		IMO:         "9262900",
		CALLSIGN:    "9HA3058",
		FLAG:        "MT  AUGUSTA",
		CURRENTPORT: "2017-01-24T17:56:00",
		DESTINATION: "US CRP FOR ORDER",
		PORTS:       []model.Port{}}

	vesselLocation.PORTS = ports

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}
func shutdown() {
	// teardown code
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func compare(a model.VesselLocation, b model.VesselLocation) bool {

	if a.SHIPNAME == b.SHIPNAME {
		return true
	} else if a.MMSI == b.MMSI {
		return true
	} else if a.IMO == b.IMO {
		return true
	} else if a.DESTINATION == b.DESTINATION {
		return true
	} else if a.SHIPNAME == b.SHIPNAME {
		return true
	} else {
		return false
	}
}
