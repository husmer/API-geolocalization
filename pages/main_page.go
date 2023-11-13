package pages

import (
	"encoding/json"
	"groupie-tracker-geolocalization/data_structs"
	"groupie-tracker-geolocalization/helpers"
	"net/http"
	"strings"
)

func MainPageStarter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		data_structs.ErrorPage.Execute(w, nil)
		return
	}

	// Get the search query from the URL query parameter "q"
	query := strings.ToLower(r.URL.Query().Get("q"))

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

	// Use the FilterArtists function to filter artists
	matchingArtists := helpers.FilterArtists(query, apiResponse.ArtistsWithRelations)
	// Render the search suggestions template with matchingArtists
	if err := data_structs.HomePage.Execute(w, matchingArtists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
