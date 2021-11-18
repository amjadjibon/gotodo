package createalbum

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

type CreateAlbum struct {
	mCM abeshModel.ConfigMap
}

func (g *CreateAlbum) Name() string {
	return "createalbum"
}

func (g *CreateAlbum) Version() string {
	return "0.0.1"
}

func (g *CreateAlbum) Category() string {
	return string(constant.CategoryService)
}

func (g *CreateAlbum) ContractId() string {
	return "createalbum"
}

func (g *CreateAlbum) New() iface.ICapability {
	return &CreateAlbum{}
}

func (g *CreateAlbum) SetConfigMap(cm abeshModel.ConfigMap) error {
	g.mCM = cm
	return nil
}

func (g *CreateAlbum) GetConfigMap() abeshModel.ConfigMap {
	return g.mCM
}

func (g *CreateAlbum) Serve(ctx context.Context, input *abeshModel.Event) (*abeshModel.Event, error) {
	inputObj := new(model.Album)

	err := json.Unmarshal(input.Value, inputObj)
	if err != nil {
		return nil, err
	}

	err = dao.CreateAlbum(inputObj)
	if err != nil {
		return nil, err
	}

	albumsByte, err := json.Marshal(inputObj)
	if err != nil {
		return nil, err
	}

	return abeshModel.GenerateOutputEvent(input.Metadata, g.ContractId(), "OK", 200, "application/json", albumsByte), nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&CreateAlbum{})
}
