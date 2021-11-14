package data

import (
	"entgo.io/ent/examples/o2o2types/ent"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewArticleRepo)

type Data struct {
	db *ent.Client
}

func NewData() *Data {
	return &Data{}
}
