package argument

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/model"
)

type GroupArgument struct {
	GroupID graphql.ID
}

func (arg GroupArgument) Member(ctx context.Context) (gm model.GroupMember, err error) {
	err = contexts.UseDB(ctx).NewSelect().Model(&gm).Where("group_id = ? AND user_id = ?", arg.GroupID, contexts.UseCurrentUser(ctx).ID).Scan(ctx)
	return
}
