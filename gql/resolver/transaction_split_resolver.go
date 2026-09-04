package resolver

import (
	"context"

	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/model"
)

type transactionSplitResolver struct {
	model.TransactionSplit
}

func (t transactionSplitResolver) Member(ctx context.Context) (*groupMemberResolver, error) {
	m, err := contexts.UseGroupMemberLoader(ctx).Load(ctx, t.MemberID)()
	if err != nil {
		return nil, err
	}
	return &groupMemberResolver{m}, nil
}
