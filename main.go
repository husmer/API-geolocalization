package main

import (
	"errors"
	"fmt"
	"groupie-tracker-geolocalization/pages"
	"net/http"
	"os"
)

func main() {
	fileServer := http.NewServeMux()
	fileServer.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	fileServer.HandleFunc("/", pages.MainPageStarter)
	fileServer.HandleFunc("/search", pages.SearchArtist)
	fileServer.HandleFunc("/database", pages.ParseApiData)
	fileServer.HandleFunc("/details", pages.DetailsPageStarter)

	fmt.Println("Running server at http://localhost:8080")
	fmt.Println("...to shut down server, press Ctrl+C")

	err := http.ListenAndServe(":8080", fileServer)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

}
