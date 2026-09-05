package resolver

import (
	"context"

	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/gql/argument"
	"github.com/xamx2/xplit-server/gql/input"
	"github.com/xamx2/xplit-server/model"
)

type mutationResolver struct{}

func (*mutationResolver) CreateGroup(ctx context.Context, args struct {
	Input input.GroupInput
}) (*groupResolver, error) {
	g := new(model.Group)
	if err := args.Input.Decode(g); err != nil {
		return nil, err
	}
	db, u := contexts.UseDB(ctx), contexts.UseAuth(ctx).MustGetUser(ctx)
	if err := u.CreateGroup(db, ctx, g); err != nil {
		return nil, err
	}
	return &groupResolver{*g}, nil
}

func (*mutationResolver) CreateGroupMember(ctx context.Context, args struct {
	argument.GroupArgument
	Input input.GroupMemberInput
}) (*groupMemberResolver, error) {
	gm := new(model.GroupMember)
	if err := args.Input.Decode(gm); err != nil {
		return nil, err
	}
	m, err := args.Member(ctx)
	if err != nil {
		return nil, err
	}
	if err := m.CreateGroupMember(contexts.UseDB(ctx), ctx, gm); err != nil {
		return nil, err
	}
	return &groupMemberResolver{*gm}, nil
}

func (*mutationResolver) CreateTransaction(ctx context.Context, args struct {
	argument.GroupArgument
	Input input.TransactionInput
}) (*transactionResolver, error) {
	t := new(model.Transaction)
	if err := args.Input.Decode(t); err != nil {
		return nil, err
	}
	m, err := args.Member(ctx)
	if err != nil {
		return nil, err
	}
	if err := m.CreateTransaction(contexts.UseDB(ctx), ctx, t); err != nil {
		return nil, err
	}
	return &transactionResolver{*t}, nil
}
