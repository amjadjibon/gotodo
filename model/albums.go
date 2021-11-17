package model

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

type Album struct {
	Id       int     `json:"id"`
	Title    string  `json:"title"`
	ArtistId int     `json:"artist_id"`
	Price    float64 `json:"price"`
}

type Response struct {
	Message string `json:"msg"`
}
