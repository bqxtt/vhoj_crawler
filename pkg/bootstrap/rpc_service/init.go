package rpc_service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/rpc_config"
	crawler "github.com/ecnuvj/vhoj_crawler/pkg/handler"
	"github.com/ecnuvj/vhoj_crawler/pkg/sdk/crawlerpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitService() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", rpc_config.UserRpc.Address.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler, err := crawler.NewCrawlerHandler()
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	s := grpc.NewServer()
	crawlerpb.RegisterCrawlerServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
