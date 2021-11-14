package model

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

type Album struct {
	Id       int64   `json:"id"`
	Title    string  `json:"title"`
	ArtistId int64   `json:"artistId"`
	Artist   *User   `json:"artist"`
	Price    float64 `json:"price"`
}

type UpdateModelInput struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Price    float64 `json:"price"`
	ArtistId int64  `json:"artist_id"`
}
