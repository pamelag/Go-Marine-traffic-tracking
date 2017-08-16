Vessel Tracking

API Provider: https://www.marinetraffic.com/

API EndPoints:

Vessel Location: http://services.marinetraffic.com/api/exportvessel/API_KEY/timespan:60/imo:8907280/msgtype:extended/protocol:jsono
Responsible for retrieving vessel's current location.
API Params: IMO number.

Vessel Ports: http://services.marinetraffic.com/api/portcalls/API_KEY/imo:8907280/fromdate:2016-12-01/todate:2016-12-10/movetype:1/protocol:jsono
Responsible for retrieving vessel's port history.
API Params: imo, fromDate, toDate, moveType(moveType can either be 0 - Arrival or 1 - departure)

FrontEnd JS Library : Leaflet.js using OpenStreet Map. We get the location coordinates and set the lat, lon of the map view.



