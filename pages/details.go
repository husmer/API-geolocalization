package pages

import (
	"encoding/json"
	"fmt"
	"groupie-tracker-geolocalization/data_structs"
	"groupie-tracker-geolocalization/helpers"
	"html/template"
	"net/http"
	"strconv"
)

func DetailsPageStarter(w http.ResponseWriter, r *http.Request) {
	// Get the search query from the URL query parameter "details"
	query := r.URL.Query().Get("details")
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

	tmpl, errTmpl := template.New("MapMarkerScript").ParseFiles("./static/details.html")
	if errTmpl != nil {
		fmt.Println("errTempl - in formting templ", errTmpl)
		// Redirect to the error page
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	errTmpl = tmpl.Execute(w, matchingArtist)
	if errTmpl != nil {
		fmt.Println("errTempl", errTmpl)
		// Redirect to the error page
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
}
