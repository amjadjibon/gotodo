package getalbumbyid

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/mkawserm/abesh/constant"
	"github.com/mkawserm/abesh/iface"
	abeshModel "github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"

	"github.com/amjadjibon/gotodo/dao"
)

type GetAlbums struct {
	mCM abeshModel.ConfigMap
}

func (g *GetAlbums) Name() string {
	return "getalbumbyid"
}

func (g *GetAlbums) Version() string {
	return "0.0.1"
}

func (g *GetAlbums) Category() string {
	return string(constant.CategoryService)
}

func (g *GetAlbums) ContractId() string {
	return "getalbumbyid"
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
	id := input.Metadata.Params["id"]

	parseInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, err
	}

	albums, err := dao.GetAlbumByID(int(parseInt))
	if err != nil {
		return nil, err
	}

	albumsByte, err := json.Marshal(albums)
	if err != nil {
		return nil, err
	}

	return abeshModel.GenerateOutputEvent(input.Metadata, g.ContractId(), "OK", 200, "application/json", albumsByte), nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&GetAlbums{})
}
