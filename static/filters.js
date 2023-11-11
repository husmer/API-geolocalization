document.addEventListener("DOMContentLoaded", function() {
    // Function to populate the table based on the filtered data
    function populateTable(filteredData) {
        const table = document.getElementById("myTable");
        // Clear the table body
        table.innerHTML = '<tr class="header"><th>Image</th><th>Name</th><th>Members</th><th>Creation Date</th><th>First Album</th><th>Locations</th></tr>';
        
        filteredData.forEach(item => {
            const row = table.insertRow();
            row.innerHTML = `
                <td style="width: 15%";><img src="${item.image}" alt="${item.name}" width="250"></td>
                <td style="width: 15%";>${item.name}</td>
                <td style="width: 15%";>${item.members.join("<br>")}</td>
                <td style="width: 5%";>${item.creationDate}</td>
                <td style="width: 10%";>${item.firstAlbum}</td>
                <td style="width: 25%";>${formatDatesLocations(item.relations.datesLocations)}</td>
            `;
        });
    }
    // Fetch the data from the Go server
    fetch("/database")
        .then(response => response.json())
        .then(data => {
            const artists = data.artistsWithRelations;
            // populate the (Locations:) filtering dropdown menu
            populateLocations(artists);
    
            // Filtering wide-criteria
            var creationDateFromFilter = 1940;
            var creationDateToFilter = 2022;
            
            var firstAlbumFrom = 1940;
            var firstAlbumTo = 2022;
            
            var concertLocation = "anyLocation"

            // Create an array to store the values of selected checkboxes
            var selectedMemberCounts = [];
            
            // Access the input element
            var rangeCreationToInput = document.getElementById("rangeValueCreationTo");
            var rangeCreationFromInput = document.getElementById("rangeValueCreationFrom");
            var rangeAlbumToInput = document.getElementById("rangeValueAlbumTo");
            var rangeAlbumFromInput = document.getElementById("rangeValueAlbumFrom");
            
            // Access the location element
            var selectedConcertLocation = document.getElementById("locations");

            // Access the checkbox elements for number of members
            var selectedNrOfMbrs = document.querySelectorAll("input[name='nmbr']");


            // Add an event listener to update creationDateToFilter when the input changes
            rangeCreationToInput.addEventListener("input", function() {
                creationDateToFilter = this.value;
                // Call the filterAndPopulateTable function to apply the filter and update the table
                filterAndPopulateTable();
            });
            
            rangeCreationFromInput.addEventListener("input", function() {
                creationDateFromFilter = this.value;
                filterAndPopulateTable();
            });
            
            rangeAlbumFromInput.addEventListener("input", function() {
                firstAlbumFrom = this.value;
                filterAndPopulateTable();
            });
            
            rangeAlbumToInput.addEventListener("input", function() {
                firstAlbumTo = this.value;
                filterAndPopulateTable();
            });
            
            selectedConcertLocation.addEventListener("change", function() {
                concertLocation = this.value;
                console.log(concertLocation);
                filterAndPopulateTable();
            });

            // Add an event listener to update the selectedMemberCounts array when checkboxes change
            selectedNrOfMbrs.forEach(checkbox => {
                checkbox.addEventListener("change", function() {
                    selectedMemberCounts = Array.from(selectedNrOfMbrs)
                        .filter(checkbox => checkbox.checked)
                        .map(checkbox => checkbox.value);
                    filterAndPopulateTable();
                });
            });
            
            function filterAndPopulateTable() {
                var filteredData = artists.filter(item => {
                    var creationYear = item.creationDate;
                    var firstAlbumYear = parseInt(item.firstAlbum.split("-")[2], 10);
                    var artistConcertLocations = item.relations.datesLocations;
                    var memberCount = item.members.length.toString();
        
                    // Check if the number of members is in the selectedMemberCounts array
                    var memberFilter = selectedMemberCounts.length === 0 || selectedMemberCounts.includes(memberCount);
        
                    var creationDateFilter = creationYear >= creationDateFromFilter && creationYear <= creationDateToFilter;
                    var firstAlbumFilter = firstAlbumYear >= firstAlbumFrom && firstAlbumYear <= firstAlbumTo;
        
                    // Check the location filter
                    var locationFilter = concertLocation === "anyLocation" || Object.keys(artistConcertLocations).includes(concertLocation);
        
                    return memberFilter && creationDateFilter && firstAlbumFilter && locationFilter;
                });
            
                // Populate the table with the filtered data
                populateTable(filteredData);
            }
    
            filterAndPopulateTable();
        })
        .catch(error => {
            console.error("Error fetching data:", error);
    });
    // Function to format datesLocations as a string and show them in the Locations column
    function formatDatesLocations(datesLocations) {
        let result = '';
        for (const location in datesLocations) {
            if (result) {
                result += '<br>'; // Add line break if there's already content
            }
            result += `${(location)}: ${datesLocations[location].join('<br>')}`;
        }
        return result;
    }
    // Function to populate the locations in the select element
    function populateLocations(places) {
        const selectElement = document.getElementById('locations'); // Get the select element
        
        const uniqueLocations = [];
        
        // Iterate through the relations array and extract unique location names
        places.forEach(item => {
            const datesLocations = item.relations.datesLocations;
            for (const location in datesLocations) {
                if (!uniqueLocations.includes(location)) {
                    uniqueLocations.push(location);
                }
            }
        });
        uniqueLocations.sort();
        // Create an option for each location
        uniqueLocations.forEach(location => {
            const option = document.createElement('option');
            option.value = location;
            option.text = location;
            selectElement.appendChild(option);
        });
    }
});
