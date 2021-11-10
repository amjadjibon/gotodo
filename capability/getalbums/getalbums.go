package getalbums

import (
	"context"
	"encoding/json"
	"github.com/amjadjibon/gotodo/model"
	"github.com/mkawserm/abesh/constant"
	"github.com/mkawserm/abesh/iface"
	abeshModel "github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"
)

type GetAlbums struct {
	mCM abeshModel.ConfigMap
}


func (g *GetAlbums) Name() string {
	return "getalbums"
}

func (g *GetAlbums) Version() string {
	return "0.0.1"
}

func (g *GetAlbums) Category() string {
	return string(constant.CategoryService)
}

func (g *GetAlbums) ContractId() string {
	return "getalbums"
}

func (g *GetAlbums) New() iface.ICapability {
	return &GetAlbums{}
}

func (g *GetAlbums) SetConfigMap(cm abeshModel.ConfigMap) error {
	g.mCM = cm
	return nil
}

func (g *GetAlbums) GetConfigMap() abeshModel.ConfigMap {
	return g.mCM
}

func (g *GetAlbums) Serve(ctx context.Context, input *abeshModel.Event) (*abeshModel.Event, error) {
	var albums []model.Album

	album1 := model.Album{
		Id:       1,
		Title:    "New1",
		ArtistId: 1,
		Artist:   &model.User{
			Id:    1,
			Name:  "Artist1",
			Genre: "New",
		},
		Price:    100,
	}

	albums = append(albums, album1)

	albumsByte, err := json.Marshal(albums)
	if err != nil {
		return nil, err
	}

	return abeshModel.GenerateOutputEvent(input.Metadata, g.ContractId(), "OK", 200, "application/json", albumsByte), nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&GetAlbums{})
}
