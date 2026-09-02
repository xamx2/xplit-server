package input

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/xamx2/xplit-server/model"
)

type GroupInput struct {
	Name graphql.NullString
}

func (i *GroupInput) Decode(g *model.Group) error {
	if i.Name.Set {
		g.Name = *i.Name.Value
	}
	return nil
}
