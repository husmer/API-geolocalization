#### Auditing directions: https://github.com/01-edu/public/tree/master/subjects/groupie-tracker/geolocalization/audit #### 

---  

# Groupie Tracker Geolocalization #

**Description**  
Groupie Trackers consists on receiving a given API and manipulating the data contained in it, in order to create a site, displaying the information. Now with an added search, filtering and geolocalization (using Google Maps API) function.

**Authors**
Oliver Vilu  

**How to run**  
1. go run .  

2. Runs on http://localhost:8080/

3. Use the search bar to find the artist or just scroll through the database

3. Press the purple "Details" button to get to the map

**Implementation details: algorithm**  
Implemented with Go server HTTP requests. CSS, HTML for styling. Search function and suggestions are based on Go templates. Database filtering is based on JS to implement asynchronous filtering. Concert locations are displayed using Google Maps API.