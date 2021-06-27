package models

type Coordinates struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func FillCoordinates(lat float64, long float64) Coordinates {
	return Coordinates{
		Lat:  lat,
		Long: long,
	}
}
