package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/eduardoacuna/self-esteem/database"
	"github.com/eduardoacuna/self-esteem/env"
	"github.com/eduardoacuna/self-esteem/log"
	"github.com/eduardoacuna/self-esteem/server"
)

func main() {
	ctx := context.WithValue(context.Background(), "ID", "system")
	log.Setup(os.Stdout)
	env.Setup(ctx)
	err := database.Setup(ctx, env.DBUser, env.DBName, env.DBHost, env.DBPort)
	if err != nil {
		log.Fatal(ctx, "error in db setup", "err", err)
	}
	defer database.Close(ctx)
	server.SetupRoutes(ctx)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-signals
		log.Warn(ctx, "shutting server down")
		database.Close(ctx)
		os.Exit(0)
	}()

	server.ListenAndServe(ctx, env.Bin, env.Address, env.Port)
}
