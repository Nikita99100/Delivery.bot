package models

func FillCourier(uuid string, status string, coordinates Coordinates) Courier {
	return Courier{
		UUID:    uuid,
		Status:  status,
		LastLat: coordinates.Lat,
		LastLon: coordinates.Long,
	}
}

type Courier struct {
	UUID    string  `json:"uuid"`
	Status  string  `json:"status"`
	LastLat float64 `json:"last_lat"`
	LastLon float64 `json:"last_lon"`
}
