package contexts

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/xamx2/xplit-server/model"
)

type gmlKey struct{}

func WithGroupMemberLoader(ctx context.Context, l *dataloader.Loader[int32, model.GroupMember]) context.Context {
	return context.WithValue(ctx, gmlKey{}, l)
}

func UseGroupMemberLoader(ctx context.Context) *dataloader.Loader[int32, model.GroupMember] {
	return ctx.Value(gmlKey{}).(*dataloader.Loader[int32, model.GroupMember])
}
