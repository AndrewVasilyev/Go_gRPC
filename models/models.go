package models

type Plants struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	WateringDay int    `json:"watering_day"`
	Species     int    `json:"species"`
	Room        int    `json:"room"`
}

type Rooms struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Species struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type WateringDays struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
