package bootstrap

import (
	"github.com/ecnuvj/vhoj_crawler/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_crawler/pkg/bootstrap/rpc_service"
)

func Init() {
	database.Init()
	rpc_service.InitService()
}
