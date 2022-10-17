package v1

import (
	"example.com/my-gin/pkg/app"
	"example.com/my-gin/pkg/e"
	"example.com/my-gin/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `form:"id" valid:"Required"`
	Title  string  `form:"title" valid:"Required"`
	Artist string  `form:"artist" valid:"Required"`
	Price  float64 `form:"price" valid:"Required"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	fmt.Println("postalbum,,,,,", setting.UserId)
	var (
		newAlbum album
		appG     = app.Gin{C: c}
	)

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	//if err := c.BindJSON(&newAlbum); err != nil {
	//	c.IndentedJSON(400, "params error")
	//	return
	//}
	httpCode, errCode := app.BindAndValid(c, &newAlbum)
	//fmt.Println(newAlbum)
	//return
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
		//c.IndentedJSON(httpCode, nil)
		//return
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(111111)

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
