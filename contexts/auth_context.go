package contexts

import (
	"context"

	"github.com/xamx2/xplit-server/model"
)

func UseCurrentUser(ctx context.Context) *model.User {
	// TODO: Implement logic to retrieve the current user from the context
	return &model.User{ID: 1}
}
