package loader

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/uptrace/bun"
	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/model"
)

func LoadGroupMember(ctx context.Context, ids []int32) (results []*dataloader.Result[model.GroupMember]) {
	var mems []model.GroupMember
	err := contexts.UseDB(ctx).NewSelect().Model(&mems).Where("id IN (?)", bun.List(ids)).Scan(ctx)

	m := map[int32]model.GroupMember{}
	for _, mem := range mems {
		m[mem.ID] = mem
	}

	for _, id := range ids {
		results = append(results, &dataloader.Result[model.GroupMember]{
			Data:  m[id],
			Error: err,
		})
	}
	return
}
