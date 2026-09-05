package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/dundunlabs/grapher"
	"github.com/dundunlabs/grapher/explorer/graphiql"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/rs/cors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/xamx2/xplit-server/contexts"
	"github.com/xamx2/xplit-server/gql"
	"github.com/xamx2/xplit-server/handler"
	"github.com/xamx2/xplit-server/loader"
	"google.golang.org/api/option"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	db := bun.NewDB(sqldb, pgdialect.New()).WithQueryHook(bundebug.NewQueryHook(bundebug.FromEnv()))
	defer db.Close()

	fbopt := option.WithAuthCredentialsFile(option.ServiceAccount, os.Getenv("FIREBASE_CREDENTIALS_FILE"))
	fbapp, err := firebase.NewApp(ctx, nil, fbopt)
	if err != nil {
		panic(err)
	}
	fbauth, err := fbapp.Auth(ctx)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", &handler.AuthHandler{
		Next: &grapher.Handler{
			Schema:   gql.NewSchema(),
			Explorer: graphiql.NewExplorer(),
		},
	})

	gml := dataloader.NewBatchedLoader(loader.LoadGroupMember)

	corsOpts := cors.Options{
		AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowedHeaders: []string{"*"},
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: cors.New(corsOpts).Handler(mux),
		BaseContext: func(_ net.Listener) context.Context {
			httpCtx := contexts.WithDB(ctx, db)
			httpCtx = contexts.WithFirebaseAuth(httpCtx, fbauth)
			httpCtx = contexts.WithGroupMemberLoader(httpCtx, gml)
			return httpCtx
		},
	}

	go srv.ListenAndServe()
	fmt.Println("Server started on :8080")

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	fmt.Println("Shutting down server...")
	srv.Shutdown(ctx)
}
