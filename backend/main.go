package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID        string  `json:"id"`
	Geschl    int8   `json:"gesch"`
	Name      string  `json:"name"`
	Alter     int8    `json:"alter"`
	Preferenz int8	  `json:"prefer"`
}

var albums = []album{
	{ID: "1", Geschl: 1, Name: "John Coltrane", Alter: 21, Preferenz:1},
	{ID: "2", Geschl: 1, Name: "Gerry Mulligan", Alter: 20, Preferenz:2},
	{ID: "3", Geschl: 2, Name: "Sarah Vaughan", Alter: 77, Preferenz:1},
	{ID: "4", Geschl: 2, Name: "Gunilla Coltrane", Alter: 21, Preferenz:2},
	{ID: "5", Geschl: 1, Name: "Peter hanfstengel", Alter: 29, Preferenz:1},
	{ID: "6", Geschl: 1, Name: "Laura Vaughan", Alter: 77, Preferenz:1},
	{ID: "7", Geschl: 2, Name: "Arni Coltrane", Alter: 21, Preferenz:1},
	{ID: "8", Geschl: 2, Name: "Andy Mulligan", Alter: 36, Preferenz:2},
	{ID: "9", Geschl: 1, Name: "Sarah Ludrigan", Alter: 44, Preferenz:2},
	{ID: "10", Geschl: 1, Name: "John Coltrane", Alter: 31, Preferenz:2},
	{ID: "11", Geschl: 2, Name: "Gerry Mulligan", Alter: 18, Preferenz:1},
	{ID: "12", Geschl: 2, Name: "Sarah Vaughan", Alter: 27, Preferenz:1},
	{ID: "13", Geschl: 1, Name: "John Coltrane", Alter: 29, Preferenz:2},
	{ID: "14", Geschl: 1, Name: "Gerry Mulligan", Alter: 59, Preferenz:2},
	{ID: "15", Geschl: 2, Name: "Sarah Wagenknechtl", Alter: 57, Preferenz:1},
	{ID: "16", Geschl: 2, Name: "Lara Sutter", Alter: 31, Preferenz:1},
	{ID: "17", Geschl: 2, Name: "Nora Korkenzieher", Alter: 49, Preferenz:1},
	{ID: "18", Geschl: 2, Name: "Lena Podgorni", Alter: 27, Preferenz:1},
	{ID: "19", Geschl: 1, Name: "Paul Siebenkäs", Alter: 21, Preferenz:2},
	{ID: "20", Geschl: 1, Name: "Ludwig Gerstel", Alter: 19, Preferenz:2},
	{ID: "21", Geschl: 2, Name: "Ludmilla Zauberbein", Alter: 38, Preferenz:1},
	{ID: "22", Geschl: 1, Name: "Leonard Kahlenfeld", Alter: 41, Preferenz:2},
	{ID: "23", Geschl: 3, Name: "Pitron Lagosna", Alter: 49, Preferenz:2},
	{ID: "24", Geschl: 3, Name: "Zeltis Kanonalis", Alter: 72, Preferenz:1},
}

// getAlbums responds with the list of all albums as JSON.
func getPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func postPerson(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
			return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getPersonID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
			if a.ID == id {
					c.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.GET("/persons/:id", getPersonID)
	router.POST("/post_person", postPerson)

	router.Run("localhost:8080")
}


