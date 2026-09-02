package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/model"
)

type userResolver struct {
	model.User
}

func (r *userResolver) Groups(ctx context.Context) (rs []*groupResolver, err error) {
	db := contexts.UseDB(ctx)
	groups, err := contexts.UseCurrentUser(ctx).Groups(db, ctx)
	if err != nil {
		return nil, err
	}
	for _, g := range groups {
		rs = append(rs, &groupResolver{g})
	}
	return
}

func (r *userResolver) Group(ctx context.Context, args struct {
	ID graphql.ID
}) (*groupResolver, error) {
	db := contexts.UseDB(ctx)
	group, err := contexts.UseCurrentUser(ctx).Group(db, ctx, args.ID)
	if err != nil {
		return nil, err
	}
	return &groupResolver{Group: *group}, nil
}
