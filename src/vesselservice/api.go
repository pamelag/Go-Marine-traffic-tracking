package vessel

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"lab.identitii.com/identitii/discord/model"
	"lab.identitii.com/identitii/go-config"
)

//AddVesselRoutes - Routes
func AddVesselRoutes(router *mux.Router) {
	router.HandleFunc("/vessel/getLocation/{imo}", GetLocationHandleFunc).Methods("GET")
	router.HandleFunc("/vessel/getVessels/{param}", GetVesselsHandleFunc).Methods("GET")
}

//GetLocationHandleFunc - Gets Vessel Locations
func GetLocationHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imoParam := vars["imo"]
	imo, errNumFmt := strconv.Atoi(imoParam)

	resp := model.VesselLocationResponse{}
	var errc error

	if errNumFmt != nil {
		resp.Error = nil
		resp.ErrorMsg = "imo not a number"
	}

	resp, errc = GetVesselLocation(imo, config.MustString("vesseltracking.APIKeyLocation"), config.MustString("vesseltracking.APIKeyPortCalls"), config.MustInt("vesseltracking.TimeSpan"))
	if errc != nil {
		resp.Error = nil
		resp.ErrorMsg = "API Error"
	}

	data, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//GetVesselsHandleFunc - Gets a collection of Vessels currently sailing
func GetVesselsHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["param"]

	vesselResponse, error := GetVessels(param)
	if error != nil {
		log.Debug(error)
	}
	js, err := json.Marshal(vesselResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
