package albums

import "github.com/google/uuid"

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums() []Album {
	return albums
}

func postAlbums(data CreateAlbumDto) (response Album, ok bool) {
	var newAlbum Album
	if data.Title == "" || data.Artist == "" {
		return newAlbum, false
	} else {
		newId := uuid.New().String()
		newAlbum = Album{Title: data.Title, Artist: data.Artist, Price: data.Price, ID: newId}
		albums = append(albums, newAlbum)
		return newAlbum, true
	}

}

func getAlbumById(id string) (res Album, ok bool) {
	for _, a := range albums {
		if a.ID == id {
			return a, true
		}
	}

	return Album{}, false
}
