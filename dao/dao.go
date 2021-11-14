package dao

import (
	"github.com/go-pg/pg/v10"

	"github.com/amjadjibon/gotodo/model"
)

type albumModel struct {
	Id       int64   `pg:"id"`
	Title    string  `pg:"title"`
	ArtistId int64   `pg:"artistId"`
	Price    float64 `pg:"price"`
}

func getDB(dsn string) *pg.DB {
	opt, err := pg.ParseURL(dsn)
	if err != nil {
		panic(err)
	}

	return pg.Connect(opt)
}

func GetAllAlbums() ([]*model.Album, error) {
	var alb []albumModel

	db := getDB("postgres://user:pass@localhost:5432/db_name?sslmode=disable")
	defer func() {
		_ = db.Close()
	}()

	err := db.Model(&alb).Select()
	if err != nil {
		return nil, err
	}

	var modAlbum []*model.Album

	for _, v := range alb {
		ab := &model.Album{
			Id:       v.Id,
			Title:    v.Title,
			ArtistId: v.ArtistId,
			Price:    v.Price,
		}

		modAlbum = append(modAlbum, ab)

		ab = new(model.Album)
	}

	return modAlbum, nil
}