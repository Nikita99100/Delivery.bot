package models

type Task struct {
	OrderNumber       string `json:"order_number"`
	Type              string `json:"type"`
	Route             Route  `json:"route"`
	ProdComplTimeFrom string `json:"prod_compl_time_from"`
	ProdComplTimeTo   string `json:"prod_compl_time_to"`
	CourierUUID       string `json:"courier_uuid"`
	PhoneNumber       string `json:"phone_number"`
}
type Route struct {
	City    string `json:"city"`
	Street  string `json:"street"`
	House   string `json:"house"`
	Comment string `json:"comment"`
	Country string `json:"country"`
}
