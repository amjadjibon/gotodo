package deletealbum

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/mkawserm/abesh/constant"
	"github.com/mkawserm/abesh/iface"
	abeshModel "github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"

	"github.com/amjadjibon/gotodo/dao"
	"github.com/amjadjibon/gotodo/model"
)

type DeleteAlbum struct {
	mCM abeshModel.ConfigMap
}

func (g *DeleteAlbum) Name() string {
	return "deletealbum"
}

func (g *DeleteAlbum) Version() string {
	return "0.0.1"
}

func (g *DeleteAlbum) Category() string {
	return string(constant.CategoryService)
}

func (g *DeleteAlbum) ContractId() string {
	return "deletealbum"
}

func (g *DeleteAlbum) New() iface.ICapability {
	return &DeleteAlbum{}
}

func (g *DeleteAlbum) SetConfigMap(cm abeshModel.ConfigMap) error {
	g.mCM = cm
	return nil
}

func (g *DeleteAlbum) GetConfigMap() abeshModel.ConfigMap {
	return g.mCM
}

func (g *DeleteAlbum) Serve(ctx context.Context, input *abeshModel.Event) (*abeshModel.Event, error) {
	id := input.Metadata.Params["id"]

	parseInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, err
	}

	err = dao.DeleteAlbum(int(parseInt))
	if err != nil {
		return nil, err
	}

	ab := model.Response{
		Message: "Album Deleted",
	}

	messageByte, err := json.Marshal(ab)
	if err != nil {
		return nil, err
	}

	return abeshModel.GenerateOutputEvent(input.Metadata, g.ContractId(), "OK", 200, "application/json", messageByte), nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&DeleteAlbum{})
}
