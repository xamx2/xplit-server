package resolver

import (
	"context"

	"github.com/xamx2/xplit-server/contexts"
)

type queryResolver struct{}

func (*queryResolver) CurrentUser(ctx context.Context) *userResolver {
	u := contexts.UseAuth(ctx).MustGetUser(ctx)
	return &userResolver{*u}
}
