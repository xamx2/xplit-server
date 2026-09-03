package resolver

import (
	"context"

	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/model"
)

type groupResolver struct {
	model.Group
}

func (r *groupResolver) Members(ctx context.Context) (rs []groupMemberResolver, err error) {
	db := contexts.UseDB(ctx)
	members, err := r.Group.Members(db, ctx)
	if err != nil {
		return nil, err
	}
	for _, member := range members {
		rs = append(rs, groupMemberResolver{member})
	}
	return
}

func (r *groupResolver) Transactions(ctx context.Context) (rs []transactionResolver, err error) {
	db := contexts.UseDB(ctx)
	ts, err := r.Group.Transactions(db, ctx)
	if err != nil {
		return nil, err
	}
	for _, t := range ts {
		rs = append(rs, transactionResolver{t})
	}
	return
}
