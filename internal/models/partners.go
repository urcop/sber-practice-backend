package models

type Partners struct {
	PartnerId      string   `json:"partnerId"`
	Title          string   `json:"title"`
	Conditions     []string `json:"conditions"`
	AdditionalInfo string   `json:"AdditionalInfo"`
	Places         []struct {
		Address     string `json:"address"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
	} `json:"places"`
}
