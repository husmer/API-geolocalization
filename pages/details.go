package pages

import (
	"encoding/json"
	"fmt"
	"groupie-tracker-geolocalization/data_structs"
	"html/template"
	"net/http"
	"path"
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

	// Render the search results template with matchingArtists
	fp := path.Join("static", "details.html")
	tmpl := template.Must(template.ParseFiles(fp))

	if err := tmpl.Execute(w, matchingArtist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
