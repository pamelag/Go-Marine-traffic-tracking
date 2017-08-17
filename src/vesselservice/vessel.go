package vesselservice

// VesselLocation models a VesselLocation
type VesselLocation struct {
	MMSI         string `json:"MMSI"`
	LAT          string `json:"LAT"`
	LON          string `json:"LON"`
	SPEED        string `json:"SPEED"`
	COURSE       string `json:"COURSE"`
	TIMESTAMP    string `json:"TIMESTAMP"`
	SHIPNAME     string `json:"SHIPNAME"`
	SHIPTYPE     string `json:"SHIPTYPE"`
	IMO          string `json:"IMO"`
	CALLSIGN     string `json:"CALLSIGN"`
	FLAG         string `json:"FLAG"`
	CURRENTPORT  string `json:"CURRENT_PORT"`
	LASTPORT     string `json:"LAST_PORT"`
	LASTPORTTIME string `json:"LAST_PORT_TIME"`
	DESTINATION  string `json:"DESTINATION"`
	ETA          string `json:"ETA"`
	LENGTH       string `json:"LENGTH"`
	WIDTH        string `json:"WIDTH"`
	DRAUGHT      string `json:"DRAUGHT"`
	GRT          string `json:"GRT"`
	DWT          string `json:"DWT"`
	YEARBUILT    string `json:"YEAR_BUILT"`
	PORTS        []Port `json:"PORTS"`
}

// Port models a Port
type Port struct {
	MMSI         string `json:"MMSI"`
	SHIPNAME     string `json:"SHIPNAME"`
	TIMESTAMPLT  string `json:"TIMESTAMP_LT"`
	TIMESTAMPUTC string `json:"TIMESTAMP_UTC"`
	MOVETYPE     string `json:"MOVE_TYPE"`
	TYPENAME     string `json:"TYPE_NAME"`
	PORTID       string `json:"PORT_ID"`
	PORTNAME     string `json:"PORT_NAME"`
}

//List models a vesseltrack location list
type List struct {
	MMSI      string `json:"MMSI"`
	NAME      string `json:"NAME"`
	INRANGE   string `json:"INRANGE"`
	DEST      string `json:"DEST"`
	ETATSTAMP string `json:"ETATSTAMP"`
	IMO       string `json:"IMO"`
	MSGID     string `json:"MSGID"`
}

// VesselInfoResponse models a response containing VesselInfo
type VesselInfoResponse struct {
	N    string `json:"n"`
	List []List `json:"list"`
}

// VesselLocationResponse struct
type VesselLocationResponse struct {
	VesselLocations []VesselLocation
	Error           error `json:",omitempty"`
	ErrorMsg        string
}
