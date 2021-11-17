package dao

import (
	"github.com/amjadjibon/gotodo/model"
	"github.com/go-pg/pg/v10"
)

type Album struct {
	tableName struct{} `pg:"albums"`
	Id        int      `pg:"id"`
	Title     string   `pg:"title"`
	Artist_Id int      `pg:"artist_id"`
	Price     float64  `pg:"price"`
}

func getDB() *pg.DB {
	opt, err := pg.ParseURL("postgres://rootuser:rootpassword@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return pg.Connect(opt)
}

func CreateAlbum(inputObj *model.Album) error {
	db := getDB()
	defer db.Close()

	newAlbum := Album{
		Id:        inputObj.Id,
		Title:     inputObj.Title,
		Artist_Id: inputObj.ArtistId,
		Price:     inputObj.Price,
	}

	_, err := db.Model(&newAlbum).Insert()
	if err != nil {
		return err
	}

	return nil
}

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
	var pgAlbum Album

	db := getDB()
	defer db.Close()

	err := db.Model(&pgAlbum).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	modAlbum := &model.Album{
		Id:       pgAlbum.Id,
		Title:    pgAlbum.Title,
		ArtistId: pgAlbum.Artist_Id,
		Price:    pgAlbum.Price,
	}

	return modAlbum, nil
}

func UpdateAlbum(inputObj *model.Album) error {
	db := getDB()
	defer db.Close()

	updatedAlbum := Album{
		Title:     inputObj.Title,
		Artist_Id: inputObj.ArtistId,
		Price:     inputObj.Price,
	}

	_, err := db.Model(&updatedAlbum).Where("id = ?", inputObj.Id).Update()
	if err != nil {
		return err
	}

	return nil
}

func DeleteAlbum(id int) error {
	db := getDB()
	defer db.Close()

	pgAlbum := new(Album)

	_, err := db.Model(pgAlbum).Where("id = ?", id).Delete()
	if err != nil {
		return err
	}

	return nil
}
