package contexts

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"sync"

	"github.com/xamx2/xplit-server/model"
)

type authKey struct{}

func WithAuth(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, authKey{}, &AuthContext{req: req})
}

func UseAuth(ctx context.Context) *AuthContext {
	return ctx.Value(authKey{}).(*AuthContext)
}

type AuthContext struct {
	req  *http.Request
	once sync.Once
	u    *model.User
	err  error
}

func (authCtx *AuthContext) GetUser(ctx context.Context) (*model.User, error) {
	authCtx.once.Do(func() {
		idtoken := strings.TrimPrefix(authCtx.req.Header.Get("Authorization"), "Bearer ")
		token, err := UseFirebaseAuth(ctx).VerifyIDToken(ctx, idtoken)
		if err != nil {
			authCtx.err = err
			return
		}

		authCtx.u = &model.User{FirebaseUID: token.UID}
		if name, ok := token.Claims["name"].(string); ok {
			authCtx.u.Name = &name
		}

		db := UseDB(ctx)
		authCtx.err = db.NewSelect().Model(authCtx.u).WherePK("firebase_uid").Scan(ctx)
		if authCtx.err == sql.ErrNoRows {
			_, authCtx.err = db.NewInsert().Model(authCtx.u).Exec(ctx)
		}
	})
	return authCtx.u, authCtx.err
}

func (authCtx *AuthContext) MustGetUser(ctx context.Context) *model.User {
	u, err := authCtx.GetUser(ctx)
	if err != nil {
		panic(err)
	}
	return u
}
