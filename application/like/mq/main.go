package main

import (
	"beyond/application/like/mq/internal/config"
	"beyond/application/like/mq/internal/logic"
	"beyond/application/like/mq/internal/svc"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/like.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	for _, consumer := range logic.Consumers(ctx, svcCtx) {
		serviceGroup.Add(consumer)
	}

	serviceGroup.Start()
	fmt.Println("like mq service is running")
}
