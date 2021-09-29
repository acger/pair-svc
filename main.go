package main

import (
	"flag"
	"fmt"

	"github.com/acger/pair-svc/internal/config"
	"github.com/acger/pair-svc/internal/server"
	"github.com/acger/pair-svc/internal/svc"
	"github.com/acger/pair-svc/pair"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/pair.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewPairServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pair.RegisterPairServer(grpcServer, srv)
		reflection.Register(grpcServer)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)


	s.Start()
}
