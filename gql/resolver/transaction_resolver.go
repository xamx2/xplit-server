package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/model"
)

type transactionResolver struct {
	model.Transaction
}

func (r transactionResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.Transaction.CreatedAt}
}

func (t transactionResolver) Splits(ctx context.Context) (rs []*transactionSplitResolver, err error) {
	tss, err := t.Transaction.Splits(contexts.UseDB(ctx), ctx)
	if err != nil {
		return nil, err
	}
	for _, ts := range tss {
		rs = append(rs, &transactionSplitResolver{ts})
	}
	return
}

func (t transactionResolver) Member(ctx context.Context) (*groupMemberResolver, error) {
	m, err := t.Transaction.Member(contexts.UseDB(ctx), ctx)
	if err != nil {
		return nil, err
	}
	return &groupMemberResolver{*m}, nil
}
