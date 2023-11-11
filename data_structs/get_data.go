package data_structs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetArtists(artistUrl string) []Artist {
	var data []Artist
	res, err := http.Get(artistUrl)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func GetDates(api string) []Dates {
	var datesData DatesData
	res, err := http.Get(api)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = json.NewDecoder(res.Body).Decode(&datesData)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var dates []Dates
	for _, localData := range datesData.Index {
		dates = append(dates, Dates{
			ID:    localData.ID,
			Dates: localData.Dates,
		})
	}
	return dates

}

func GetLocations(locationsUrl string) []Locations {
	var locationsData LocationsData
	res, err := http.Get(locationsUrl)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = json.NewDecoder(res.Body).Decode(&locationsData)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var locations []Locations

	for _, localData := range locationsData.Index {
		locations = append(locations, Locations{
			ID:        localData.ID,
			Locations: localData.Locations,
			Dates:     localData.Dates,
		})
	}

	return locations
}

func GetRelations(relationsUrl string) []Relations {
	var relationsData RelationsData
	res, err := http.Get(relationsUrl)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = json.NewDecoder(res.Body).Decode(&relationsData)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var relations []Relations

	for _, localData := range relationsData.Index {
		relations = append(relations, Relations{
			ID:             localData.ID,
			DatesLocations: localData.DatesLocations,
		})
	}

	return relations
}
