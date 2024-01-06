package partners

type PartnersResponse struct {
	PartnerId      string   `json:"partnerId"`
	Title          string   `json:"title"`
	Conditions     []string `json:"conditions"`
	AdditionalInfo string   `json:"AdditionalInfo"`
	Places         []struct {
		PlaceId     int    `json:"place_id"`
		PartnerId   string `json:"partner_id"`
		Address     string `json:"address"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
	} `json:"places"`
}

type PartnersRequest struct {
	Title          string   `json:"title"`
	Conditions     []string `json:"conditions"`
	AdditionalInfo string   `json:"additionalInfo"`
	Places         []struct {
		Address     string `json:"address"`
		Coordinates struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
	} `json:"places"`
}

type ResponseError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
