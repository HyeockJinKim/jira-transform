package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HyeockJinKim/jira-transform/config"
	"github.com/HyeockJinKim/jira-transform/destination"
	"github.com/HyeockJinKim/jira-transform/destination/ghdst"
	"github.com/HyeockJinKim/jira-transform/servers"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conf, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %+v\n", conf)
	server := startServer(ctx, conf)
	// signal handling
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	cancel()
	stopCtx, stopCancel := context.WithTimeout(context.Background(), time.Duration(conf.HTTP.ShutdownTimeoutMs)*time.Millisecond)
	defer stopCancel()
	if err := server.Stop(stopCtx); err != nil {
		panic(err)
	}
}

func startServer(ctx context.Context, conf *config.Config) *servers.Server {
	const (
		github = "github"
	)
	server := servers.NewServer(servers.ServerArgs{
		Addr:   conf.HTTP.Addr,
		APIKey: conf.HTTP.APIKey,
		Handler: servers.NewHandler(servers.HandlerArgs{
			Destinations: map[string]destination.Destination{
				github: ghdst.NewDestination(ghdst.Args{
					Token: conf.Mirror.Github.Token,
					Owner: conf.Mirror.Github.Owner,
				}),
			}}),
	})
	go func() {
		if err := server.Start(ctx); err != nil {
			panic(err)
		}
	}()
	return server
}
