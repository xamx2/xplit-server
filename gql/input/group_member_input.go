package input

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/xamx2/xplit-server/model"
)

type GroupMemberInput struct {
	Name graphql.NullString
}

func (i *GroupMemberInput) Decode(m *model.GroupMember) error {
	if i.Name.Set {
		m.Name = i.Name.Value
	}
	return nil
}
