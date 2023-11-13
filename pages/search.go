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

	fmt.Println("Testing:")
	i := 0
	for k := range matchingArtists[0].Relations.DatesLocations {
		keys[i] = k
		i++
	}
	fmt.Println("printing keys[0]:", keys[0])
	fmt.Println("printing all keys:", keys)
	resultLat, resultLong, err := helpers.LocationsToCoordinates("Tokyo") // calling the function

	if err != nil {
		fmt.Println("Testing expected error:", err)
	} else {
		fmt.Println("Expected latitude:", resultLat, "Expected longitude:", resultLong)
	}

	gotLat, gotLong, err := helpers.LocationsToCoordinates(keys[0]) // calling the function

	if err != nil {
		fmt.Println("Testing got error:", err)
	} else {
		fmt.Println("Real latitude:", gotLat, "Real longitude:", gotLong)
	}

	// testing end

	// Render the search results template with matchingArtists
	fp := path.Join("static", "search_results.html")
	tmpl := template.Must(template.ParseFiles(fp))

	if err := tmpl.Execute(w, matchingArtists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
