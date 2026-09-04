package resolver

import (
	"context"

	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/model"
)

type settlementResolver struct {
	model.Settlement
}

func (r settlementResolver) From(ctx context.Context) (*groupMemberResolver, error) {
	gm, err := contexts.UseGroupMemberLoader(ctx).Load(ctx, r.FromID)()
	if err != nil {
		return nil, err
	}
	return &groupMemberResolver{gm}, nil
}

func (r settlementResolver) To(ctx context.Context) (*groupMemberResolver, error) {
	gm, err := contexts.UseGroupMemberLoader(ctx).Load(ctx, r.ToID)()
	if err != nil {
		return nil, err
	}
	return &groupMemberResolver{gm}, nil
}
