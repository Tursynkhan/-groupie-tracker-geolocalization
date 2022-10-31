package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type YandexGeoCoderResp struct {
	Response struct {
		GeoObjectCollection struct {
			MetaDataProperty struct {
				GeocoderResponseMetaData struct {
					Request string `json:"request"`
					Results string `json:"results"`
					Found   string `json:"found"`
				} `json:"GeocoderResponseMetaData"`
			} `json:"metaDataProperty"`
			FeatureMember []struct {
				GeoObject struct {
					MetaDataProperty struct {
						GeocoderMetaData struct {
							Precision string `json:"precision"`
							Text      string `json:"text"`
							Kind      string `json:"kind"`
							Address   struct {
								CountryCode string `json:"country_code"`
								Formatted   string `json:"formatted"`
								Components  []struct {
									Kind string `json:"kind"`
									Name string `json:"name"`
								} `json:"Components"`
							} `json:"Address"`
							AddressDetails struct {
								Country struct {
									AddressLine        string `json:"AddressLine"`
									CountryNameCode    string `json:"CountryNameCode"`
									CountryName        string `json:"CountryName"`
									AdministrativeArea struct {
										AdministrativeAreaName string `json:"AdministrativeAreaName"`
									} `json:"AdministrativeArea"`
								} `json:"Country"`
							} `json:"AddressDetails"`
						} `json:"GeocoderMetaData"`
					} `json:"metaDataProperty"`
					Name        string `json:"name"`
					Description string `json:"description"`
					BoundedBy   struct {
						Envelope struct {
							LowerCorner string `json:"lowerCorner"`
							UpperCorner string `json:"upperCorner"`
						} `json:"Envelope"`
					} `json:"boundedBy"`
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}

func GetCoordOfCity(datesLocations map[string][]string) [][]string {
	coords := [][]string{}
	for city := range datesLocations {
		city = strings.ReplaceAll(city, " ", "+")
		url := "https://geocode-maps.yandex.ru/1.x/?apikey=a8ff2dc8-4e35-4012-85c3-ddcba1569483&format=json&geocode=" + city
		resp, err := http.Get(url)
		if err != nil {
			log.Print(err)
		}
		defer resp.Body.Close()
		cityInfo := YandexGeoCoderResp{}
		json.NewDecoder(resp.Body).Decode(&cityInfo)
		pos := strings.Split(cityInfo.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Point.Pos, " ")
		pos[0], pos[1] = pos[1], pos[0]
		coords = append(coords, pos)
	}
	return coords
}
