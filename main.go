package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

type test struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.GET("/test", getTest)
    router.POST("/test", postTest)


    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    fmt.Println("test",c)
    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    // Add the new album to the slice.

    fmt.Println("newAlbum",newAlbum)

    albums = append(albums, newAlbum)
    fmt.Println()

    c.JSON(http.StatusCreated,newAlbum)
    // c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")
    fmt.Println(c.Params)
    fmt.Println(c)

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getTest(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "test")
}

func postTest(c *gin.Context) {
    var newTest test

    if err := c.BindJSON(&newTest); err != nil {
        return
    }

    fmt.Println("newAlbum",newTest)
    c.JSON(http.StatusCreated,newTest)
}
