package vessel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"lab.identitii.com/identitii/discord/model"
	logger "lab.identitii.com/identitii/go-logger"
)

var log = logger.GetLogger("VesselService::Server")

const (
	apiLocation = "http://services.marinetraffic.com/api/exportvessel/%s/timespan:%d/imo:%d/msgtype:extended/protocol:jsono"
	apiPortCall = "http://services.marinetraffic.com/api/portcalls/%s/imo:%d/fromdate:%s/todate:%s/movetype:%d/protocol:jsono"
)

//GetVesselLocation - Service to return the vessel info
func GetVesselLocation(imo int, apiLocationKey string, apiPortKey string, timeSpan int) (model.VesselLocationResponse, error) {
	log.Println("GetVesselLocation ")
	location := getVesselLocation(imo, apiLocationKey, timeSpan)

	if location.Error != nil {
		return model.VesselLocationResponse{}, location.Error
	}
	if len(location.VesselLocations) > 0 {
		location := getPorts(imo, apiPortKey, location)
		if location.Error != nil {
			return model.VesselLocationResponse{}, location.Error
		}
	}

	log.Info("result ", location.VesselLocations)

	return location, nil
}

func getVesselLocation(imo int, apiKey string, timeSpan int) model.VesselLocationResponse {
	vesselLocations := make([]model.VesselLocation, 0)
	if imo == 0 {
		log.Debug("IMO is not provided")
	}
	url := fmt.Sprintf(apiLocation, apiKey, timeSpan, imo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Debug(err)
		return model.VesselLocationResponse{VesselLocations: nil, Error: err}
	}

	client := &http.Client{}
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Debug("Do: ", err1)
		return model.VesselLocationResponse{VesselLocations: nil, Error: err1}
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Debug(err2)
		return model.VesselLocationResponse{VesselLocations: nil, Error: err2}
	}

	err3 := json.Unmarshal(body, &vesselLocations)
	if err3 != nil {
		log.Debug(err3)
		return model.VesselLocationResponse{VesselLocations: nil, Error: err3}
	}
	return model.VesselLocationResponse{VesselLocations: vesselLocations, Error: nil}
}

func getPorts(imo int, apiPortKey string, response model.VesselLocationResponse) model.VesselLocationResponse {
	ports := make([]model.Port, 0)

	moveType := 1

	if imo == 0 {
		log.Debug("IMO is not provided")
	}

	fromDate := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	toDate := time.Now().Format("2006-01-02")

	url := fmt.Sprintf(apiPortCall, apiPortKey, imo, fromDate, toDate, moveType)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Debug(err)
		return model.VesselLocationResponse{VesselLocations: nil, Error: err}
	}

	client := &http.Client{}
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Debug("Do: ", err1)
		return model.VesselLocationResponse{VesselLocations: nil, Error: err1}
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Debug(err2)
		return model.VesselLocationResponse{VesselLocations: nil, Error: err2}
	}
	err3 := json.Unmarshal(body, &ports)
	if err3 != nil {
		log.Debug(err3)
		return model.VesselLocationResponse{VesselLocations: nil, Error: err3}
	}

	response.VesselLocations[0].PORTS = ports

	return response
}

//GetVessels - Find the vessel by the IMO
func GetVessels(imo string) (vesselInfoResponse model.VesselInfoResponse, errc error) {
	vesselInfoResponse.List = make([]model.List, 0)

	req, err := http.NewRequest("GET", "https://www.vesselfinder.com/vessels/livesearch?term="+imo, nil)
	if err != nil {
		log.Debug(err)
	}
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Debug(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	err1 := json.Unmarshal(body, &vesselInfoResponse)

	if err1 != nil {
		fmt.Println("unmarshal error", err1)
	}

	log.Info("GetVessels : Service response ", vesselInfoResponse)
	return
}
