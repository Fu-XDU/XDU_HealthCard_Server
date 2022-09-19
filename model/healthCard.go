package model

type HealthCard struct {
	ID       int    `json:"id"`
	Openid   string `gorm:"type:varchar(40);" json:"openid,omitempty"`
	StuID    string `gorm:"type:varchar(20);" json:"stu_id"`
	Passwd   string `json:"passwd,omitempty"`
	Location string `gorm:"type:varchar(1023);" json:"location"`
}

type HealthCardRequest struct {
	ID        int     `json:"id"`
	Openid    string  `json:"openid"`
	StuID     string  `json:"stu_id"`
	Passwd    string  `json:"passwd"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (h *HealthCardRequest) ToHealthCard(location string) *HealthCard {
	return &HealthCard{
		ID:       h.ID,
		Openid:   h.Openid,
		StuID:    h.StuID,
		Passwd:   h.Passwd,
		Location: location,
	}
}

type HealthCardLocation struct {
	Address    string `json:"address"`
	GeoApiInfo struct {
		Type     string `json:"type"`
		Info     string `json:"info"`
		Status   int    `json:"status"`
		Da       string `json:"$Da"`
		Position struct {
			Q   float64 `json:"Q"`
			R   float64 `json:"R"`
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"position"`
		Message          string      `json:"message"`
		LocationType     string      `json:"location_type"`
		Accuracy         interface{} `json:"accuracy"`
		IsConverted      bool        `json:"isConverted"`
		AddressComponent struct {
			Citycode         string        `json:"citycode"`
			Adcode           string        `json:"adcode"`
			BusinessAreas    []interface{} `json:"businessAreas"`
			NeighborhoodType string        `json:"neighborhoodType"`
			Neighborhood     string        `json:"neighborhood"`
			Building         string        `json:"building"`
			BuildingType     string        `json:"buildingType"`
			Street           string        `json:"street"`
			StreetNumber     string        `json:"streetNumber"`
			Country          string        `json:"country"`
			Province         string        `json:"province"`
			City             string        `json:"city"`
			District         string        `json:"district"`
			Township         string        `json:"township"`
		} `json:"addressComponent"`
		FormattedAddress string        `json:"formattedAddress"`
		Roads            []interface{} `json:"roads"`
		Crosses          []interface{} `json:"crosses"`
		Pois             []interface{} `json:"pois"`
	} `json:"geo_api_info"`
	Area     string `json:"area"`
	Province string `json:"province"`
	City     string `json:"city"`
}

type ConvertLocationResult struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Result    struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Address            string `json:"address"`
		FormattedAddresses struct {
			Recommend string `json:"recommend"`
			Rough     string `json:"rough"`
		} `json:"formatted_addresses"`
		AddressComponent struct {
			Nation       string `json:"nation"`
			Province     string `json:"province"`
			City         string `json:"city"`
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_component"`
		AdInfo struct {
			NationCode string `json:"nation_code"`
			Adcode     string `json:"adcode"`
			CityCode   string `json:"city_code"`
			Name       string `json:"name"`
			Location   struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Nation   string `json:"nation"`
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
		} `json:"ad_info"`
		AddressReference struct {
			Town struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"town"`
			StreetNumber struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance float64 `json:"_distance"`
			} `json:"street_number"`
			Street struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"street"`
			LandmarkL2 struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Distance int    `json:"_distance"`
				DirDesc  string `json:"_dir_desc"`
			} `json:"landmark_l2"`
		} `json:"address_reference"`
	} `json:"result"`
}

func (c *ConvertLocationResult) ToHealthCardLocation() *HealthCardLocation {
	return &HealthCardLocation{
		Address: c.Result.Address,
		GeoApiInfo: struct {
			Type     string `json:"type"`
			Info     string `json:"info"`
			Status   int    `json:"status"`
			Da       string `json:"$Da"`
			Position struct {
				Q   float64 `json:"Q"`
				R   float64 `json:"R"`
				Lng float64 `json:"lng"`
				Lat float64 `json:"lat"`
			} `json:"position"`
			Message          string      `json:"message"`
			LocationType     string      `json:"location_type"`
			Accuracy         interface{} `json:"accuracy"`
			IsConverted      bool        `json:"isConverted"`
			AddressComponent struct {
				Citycode         string        `json:"citycode"`
				Adcode           string        `json:"adcode"`
				BusinessAreas    []interface{} `json:"businessAreas"`
				NeighborhoodType string        `json:"neighborhoodType"`
				Neighborhood     string        `json:"neighborhood"`
				Building         string        `json:"building"`
				BuildingType     string        `json:"buildingType"`
				Street           string        `json:"street"`
				StreetNumber     string        `json:"streetNumber"`
				Country          string        `json:"country"`
				Province         string        `json:"province"`
				City             string        `json:"city"`
				District         string        `json:"district"`
				Township         string        `json:"township"`
			} `json:"addressComponent"`
			FormattedAddress string        `json:"formattedAddress"`
			Roads            []interface{} `json:"roads"`
			Crosses          []interface{} `json:"crosses"`
			Pois             []interface{} `json:"pois"`
		}{
			Type:   "complete",
			Info:   "SUCCESS",
			Status: 1,
			Da:     "jsonp_106929_",
			Position: struct {
				Q   float64 `json:"Q"`
				R   float64 `json:"R"`
				Lng float64 `json:"lng"`
				Lat float64 `json:"lat"`
			}{
				Q:   c.Result.Location.Lat,
				R:   c.Result.Location.Lng,
				Lng: c.Result.Location.Lng,
				Lat: c.Result.Location.Lat,
			},
			Message:      "Get ipLocation success.Get address success.",
			LocationType: "ip",
			Accuracy:     "null",
			IsConverted:  true,
			AddressComponent: struct {
				Citycode         string        `json:"citycode"`
				Adcode           string        `json:"adcode"`
				BusinessAreas    []interface{} `json:"businessAreas"`
				NeighborhoodType string        `json:"neighborhoodType"`
				Neighborhood     string        `json:"neighborhood"`
				Building         string        `json:"building"`
				BuildingType     string        `json:"buildingType"`
				Street           string        `json:"street"`
				StreetNumber     string        `json:"streetNumber"`
				Country          string        `json:"country"`
				Province         string        `json:"province"`
				City             string        `json:"city"`
				District         string        `json:"district"`
				Township         string        `json:"township"`
			}{
				Citycode:         c.Result.AdInfo.CityCode,
				Adcode:           c.Result.AdInfo.Adcode,
				BusinessAreas:    nil,
				NeighborhoodType: "",
				Neighborhood:     "",
				Building:         "",
				BuildingType:     "",
				Street:           c.Result.AddressComponent.Street,
				StreetNumber:     c.Result.AddressComponent.StreetNumber,
				Country:          c.Result.AddressComponent.Nation,
				Province:         c.Result.AddressComponent.Province,
				City:             c.Result.AddressComponent.City,
				District:         c.Result.AddressComponent.District,
				Township:         c.Result.FormattedAddresses.Recommend,
			},
			FormattedAddress: c.Result.FormattedAddresses.Recommend,
			Roads:            nil,
			Crosses:          nil,
			Pois:             nil,
		},
		Area:     c.Result.AddressComponent.Province + " " + c.Result.AddressComponent.City + " " + c.Result.AddressComponent.District,
		Province: c.Result.AddressComponent.Province,
		City:     c.Result.AddressComponent.City,
	}
}
