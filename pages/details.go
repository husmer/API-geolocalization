package pages

import (
	"encoding/json"
	"fmt"
	"groupie-tracker-geolocalization/data_structs"
	"groupie-tracker-geolocalization/helpers"
	"net/http"
	"strconv"
)

func DetailsPageStarter(w http.ResponseWriter, r *http.Request) {
	// Get the search query from the URL query parameter "details"
	query := r.URL.Query().Get("details")
	fmt.Println(query)
	intQuery, err := strconv.Atoi(query)
	if err != nil {
		fmt.Println("failed strconv.Atoi the user query")
	}

	// Fetch the artists data from the API by making an HTTP request
	response, err := http.Get("http://localhost:8080/database")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var apiResponse struct {
		ArtistsWithRelations []data_structs.ArtistWithRelations `json:"artistsWithRelations"`
	}

	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// find artist by the id
	matchingArtist := apiResponse.ArtistsWithRelations[intQuery-1]

	err = helpers.AddCoordinates(&matchingArtist)
	fmt.Println(err)
	if err != nil {
		// Handle the error
		fmt.Println("Error adding coordinates:", err)
		// Error for the client
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Coordinates are:", matchingArtist.Coordinates)

	/*
			// testing
			fmt.Println("Testing:")
			// Iterate over the original keys and store unique ones in the map
			locationKeys := make(map[string]bool)

			for k := range matchingArtist.Relations.DatesLocations {
				locationKeys[k] = true
			}
			// A slice to hold the unique keys
			uniqueKeys := make([]string, 0, len(locationKeys))

			// Populate the keys slice with unique keys from the map
			for k := range locationKeys {
				uniqueKeys = append(uniqueKeys, k)
			}

			//fmt.Println("printing unqiueKeys[0]:", uniqueKeys[0])
			//fmt.Println("printing all keys:", uniqueKeys)
			//plusUniqueKeys := helpers.FormatLocationMaps(uniqueKeys)
			// fmt.Println("printing formatted keys:", plusKey)

			//resultLat, resultLong, err := helpers.LocationsToCoordinates(plusUniqueKeys[0]) // calling the function


		if err != nil {
			fmt.Println("Testing uniqueKeys error:", err)
		} else {
			fmt.Println("Expected latitude:", resultLat, "Expected longitude:", resultLong)
		}*/

	if err := data_structs.DetailsPage.Execute(w, matchingArtist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
