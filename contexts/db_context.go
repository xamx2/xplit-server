package contexts

import (
	"context"

	"github.com/uptrace/bun"
)

type dbKey struct{}

func WithDB(ctx context.Context, db bun.IDB) context.Context {
	return context.WithValue(ctx, dbKey{}, db)
}

func UseDB(ctx context.Context) bun.IDB {
	return ctx.Value(dbKey{}).(bun.IDB)
}
