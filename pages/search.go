package pages

import (
	"encoding/json"
	"fmt"
	"groupie-tracker-geolocalization/data_structs"
	"groupie-tracker-geolocalization/helpers"
	"html/template"
	"net/http"
	"path"
	"strings"
)

func SearchArtist(w http.ResponseWriter, r *http.Request) {
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

	// testing

	keys := make([]string, len(matchingArtists[0].Relations.DatesLocations))

	fmt.Println("testing")
	i := 0
	for k := range matchingArtists[0].Relations.DatesLocations {
		keys[i] = k
		i++
	}
	helpers.LocationsToCoordinates(keys[0])

	// Render the search results template with matchingArtists
	fp := path.Join("static", "search_results.html")
	tmpl := template.Must(template.ParseFiles(fp))

	if err := tmpl.Execute(w, matchingArtists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
