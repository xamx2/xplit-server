package contexts

import (
	"context"

	"firebase.google.com/go/auth"
)

type fbAuthKey struct{}

func WithFirebaseAuth(ctx context.Context, client *auth.Client) context.Context {
	return context.WithValue(ctx, fbAuthKey{}, client)
}

func UseFirebaseAuth(ctx context.Context) *auth.Client {
	return ctx.Value(fbAuthKey{}).(*auth.Client)
}
