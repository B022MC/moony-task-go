package model

type AddressByCoordinatesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Location           *Location           `json:"location"`
		Address            string              `json:"address"`
		FormattedAddresses *FormattedAddresses `json:"formatted_addresses"`
		AddressComponent   *AddressComponent   `json:"address_component"`
		AdInfo             *AdInfo             `json:"ad_info"`
		AddressReference   *AddressReference   `json:"address_reference"`
	} `json:"result"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type FormattedAddresses struct {
	Recommend       string `json:"recommend"`
	Rough           string `json:"rough"`
	StandardAddress string `json:"standard_address"`
}

type AddressComponent struct {
	Nation       string `json:"nation"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
}

type AdInfo struct {
	NationCode    string    `json:"nation_code"`
	Adcode        string    `json:"adcode"`
	PhoneAreaCode string    `json:"phone_area_code"`
	CityCode      string    `json:"city_code"`
	Name          string    `json:"name"`
	Location      *Location `json:"location"`
	Nation        string    `json:"nation"`
	Province      string    `json:"province"`
	City          string    `json:"city"`
	District      string    `json:"district"`
}

type AddressReference struct {
	StreetNumber *ReferenceInfo `json:"street_number"`
	Crossroad    *ReferenceInfo `json:"crossroad"`
	Town         *ReferenceInfo `json:"town"`
	Street       *ReferenceInfo `json:"street"`
	LandmarkL2   *ReferenceInfo `json:"landmark_l2"`
}

type ReferenceInfo struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Location *Location `json:"location"`
	Distance float64   `json:"_distance"`
	DirDesc  string    `json:"_dir_desc"`
}
