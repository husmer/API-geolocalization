package pages

import (
	"encoding/json"
	"net/http"

	"groupie-tracker-geolocalization/data_structs"
	"groupie-tracker-geolocalization/helpers"
)

var artists []data_structs.Artist
var relations []data_structs.Relations

var artistsWithRelations []data_structs.ArtistWithRelations

// Create a new, excessively complicated struct because I have trouble building async filtering with many structs :( sorry!
func combineArtistsAndRelations(artists []data_structs.Artist, relations []data_structs.Relations) []data_structs.ArtistWithRelations {
	relationsMap := make(map[int]data_structs.Relations)
	for _, relation := range relations {
		relationsMap[relation.ID] = relation
	}

	// Create the combined struct
	var combinedArtists []data_structs.ArtistWithRelations
	for _, artist := range artists {
		// Check if a corresponding relation exists
		if relation, ok := relationsMap[artist.Id]; ok {
			// Iterate through DatesLocations and format the keys
			formattedDatesLocations := make(map[string][]string)
			for key, value := range relation.DatesLocations {
				formattedKey := helpers.FormatLocationText(key)
				formattedDatesLocations[formattedKey] = value
			}
			combinedArtist := data_structs.ArtistWithRelations{
				ID:           artist.Id,
				Image:        artist.Image,
				Name:         artist.Name,
				Members:      artist.Members,
				CreationDate: artist.CreationDate,
				FirstAlbum:   artist.FirstAlbum,
				Locations:    artist.Locations,
				ConcertDates: artist.ConcertDates,
				Relations: data_structs.Relations{
					ID:             relation.ID,
					DatesLocations: formattedDatesLocations,
				},
			}
			combinedArtists = append(combinedArtists, combinedArtist)
		}
	}

	return combinedArtists
}

func ParseApiData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data = make(map[string]interface{})

	artists = data_structs.GetArtists("https://groupietrackers.herokuapp.com/api/artists")
	relations = data_structs.GetRelations("https://groupietrackers.herokuapp.com/api/relation")

	// Combine artists and relations into a new struct
	artistsWithRelations = combineArtistsAndRelations(artists, relations)

	data["artistsWithRelations"] = artistsWithRelations
	json.NewEncoder(w).Encode(data)
}
