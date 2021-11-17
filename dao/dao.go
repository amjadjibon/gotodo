package dao

import (
	"github.com/go-pg/pg/v10"

	"github.com/amjadjibon/gotodo/model"
)

type Album struct {
	tableName struct{} `pg:"album"`
	Id        int      `pg:"id"`
	Title     string   `pg:"title"`
	Artist_Id int      `pg:"artist_id"`
	Price     float64  `pg:"price"`
}

func getDB() *pg.DB {
	opt, err := pg.ParseURL("postgres://baseuser:basepassword@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return pg.Connect(opt)
}

//func getDB() *pg.DB {
//	opt := pg.Options{
//		User:     "rootuser",
//		Password: "rootpassword",
//		Database: "postgres",
//	}
//
//	return pg.Connect(&opt)
//}

func GetAllAlbums() ([]*model.Album, error) {
	db := getDB()
	defer db.Close()

	var pgAlbums []Album

	err := db.Model(&pgAlbums).Select()
	if err != nil {
		return nil, err
	}

	if pgAlbums == nil {
		println("\nError!\n")
	}

	var modAlbums []*model.Album

	for _, v := range pgAlbums {
		ab := &model.Album{
			Id:       v.Id,
			Title:    v.Title,
			ArtistId: v.Artist_Id,
			Price:    v.Price,
		}

		modAlbums = append(modAlbums, ab)

		ab = new(model.Album)
	}

	return modAlbums, nil
}

func GetAlbumByID(id int) (*model.Album, error) {
	pgAlbum := new(Album)
	// var modAlbum *model.Album

	db := getDB()
	defer db.Close()

	err := db.Model(pgAlbum).Where("pgAlbum.Id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	modAlbum := &model.Album{
		// Id:       pgAlbum.Id,
		Title: pgAlbum.Title,
		// ArtistId: pgAlbum.Artist_Id,
		// Price:    pgAlbum.Price,
	}

	return modAlbum, nil
}

type updateAlbumModel struct {
	Id       int64
	Price    float64 `pg:"price"`
	ArtistId int64   `pg:"artist_id"`
}

func UpdateAlbum(inputObj *model.UpdateModelInput) error {
	// db := getDB("postgres://rootuser:rootpassword@localhost:5432/postgres?sslmode=disable")
	db := getDB()
	defer func() {
		_ = db.Close()
	}()

	m := &updateAlbumModel{
		Price:    inputObj.Price,
		ArtistId: inputObj.ArtistId,
	}

	_, err := db.Model(&m).Where("id = ?", inputObj.Id).Update()
	if err != nil {
		return err
	}

	return nil
}

func CreateAlbum(inputObj *model.Album) error {
	// db := getDB("postgres://rootuser:rootpassword@localhost:5432/postgres?sslmode=disable")
	db := getDB()
	defer db.Close()

	_, err := db.Model(&inputObj).Insert()
	if err != nil {
		return err
	}

	return nil
}
