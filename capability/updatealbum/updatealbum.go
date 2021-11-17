package updatealbum

import (
	"context"
	"encoding/json"

	"github.com/mkawserm/abesh/constant"
	"github.com/mkawserm/abesh/iface"
	abeshModel "github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"

	"github.com/amjadjibon/gotodo/dao"
	"github.com/amjadjibon/gotodo/model"
)

type UpdateAlbum struct {
	mCM abeshModel.ConfigMap
}

func (g *UpdateAlbum) Name() string {
	return "getalbums"
}

func (g *UpdateAlbum) Version() string {
	return "0.0.1"
}

func (g *UpdateAlbum) Category() string {
	return string(constant.CategoryService)
}

func (g *UpdateAlbum) ContractId() string {
	return "updatealbum"
}

func (g *UpdateAlbum) New() iface.ICapability {
	return &UpdateAlbum{}
}

func (g *UpdateAlbum) SetConfigMap(cm abeshModel.ConfigMap) error {
	g.mCM = cm
	return nil
}

func (g *UpdateAlbum) GetConfigMap() abeshModel.ConfigMap {
	return g.mCM
}

func (g *UpdateAlbum) Serve(ctx context.Context, input *abeshModel.Event) (*abeshModel.Event, error) {
	inputObj := new(model.UpdateModelInput)
	err := json.Unmarshal(input.Value, inputObj)
	if err != nil {
		return nil, err
	}

	err = dao.UpdateAlbum(inputObj)
	if err != nil {
		return nil, err
	}

	albums := &model.Album{
		// Id:       inputObj.Id,
		Title: inputObj.Title,
		// ArtistId: inputObj.ArtistId,
		// Price: inputObj.Price,
	}

	albumsByte, err := json.Marshal(albums)
	if err != nil {
		return nil, err
	}

	return abeshModel.GenerateOutputEvent(input.Metadata, g.ContractId(), "OK", 200, "application/json", albumsByte), nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&UpdateAlbum{})
}
