package helpers

import (
	"groupie-tracker-filters/data_structs"
	"strconv"
	"strings"
)

func FilterArtists(query string, artists []data_structs.ArtistWithRelations) []data_structs.ArtistWithRelations {
	var matchingArtists []data_structs.ArtistWithRelations

	for _, artist := range artists {
		queryAsInt, _ := strconv.Atoi(query)
		if strings.Contains(strings.ToLower(artist.Name), query) {
			matchingArtists = append(matchingArtists, artist)
		} else if strings.ToLower(artist.FirstAlbum) == query {
			matchingArtists = append(matchingArtists, artist)
		} else if artist.CreationDate == queryAsInt {
			matchingArtists = append(matchingArtists, artist)
		} else if artist.Relations.DatesLocations[query] != nil {
			matchingArtists = append(matchingArtists, artist)
		} else {
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), query) {
					matchingArtists = append(matchingArtists, artist)
				}
			}
		}
	}

	return matchingArtists
}
