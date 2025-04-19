package main

import (
	"flag"
	"fmt"

	"consistent_hash/consistent_hash/proto"
	"consistent_hash/internal/config"
	consistenthashserverServer "consistent_hash/internal/server/consistenthashserver"
	proxyserviceServer "consistent_hash/internal/server/proxyservice"
	servicediscoveryServer "consistent_hash/internal/server/servicediscovery"
	"consistent_hash/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/consistenthash.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		proto_consistent_hash.RegisterServiceDiscoveryServer(grpcServer, servicediscoveryServer.NewServiceDiscoveryServer(ctx))
		proto_consistent_hash.RegisterConsistentHashServerServer(grpcServer, consistenthashserverServer.NewConsistentHashServerServer(ctx))
		proto_consistent_hash.RegisterProxyServiceServer(grpcServer, proxyserviceServer.NewProxyServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
