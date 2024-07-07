package svc

import (
	"beyond/application/article/api/internal/config"
	"beyond/application/article/rpc/article"
	"beyond/pkg/interceptors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/zrpc"
)

const (
	defaultOssConnectTimeout   = 2
	defaultOssReadWriteTimeout = 6
)

type ServiceContext struct {
	Config     config.Config
	OssClient  *oss.Client
	ArticleRPC article.Article
}

func NewServiceContext(c config.Config) *ServiceContext {
	if c.ALiYunOSS.ConnectTimeout == 0 {
		c.ALiYunOSS.ConnectTimeout = defaultOssConnectTimeout
	}
	if c.ALiYunOSS.ReadWriteTimeout == 0 {
		c.ALiYunOSS.ReadWriteTimeout = defaultOssReadWriteTimeout
	}
	ossClient, err := oss.New(c.ALiYunOSS.Endpoint, c.ALiYunOSS.AK, c.ALiYunOSS.AS)
	if err != nil {
		return nil
	}

	articleRPC := zrpc.MustNewClient(c.ArticleRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:     c,
		OssClient:  ossClient,
		ArticleRPC: article.NewArticle(articleRPC),
	}
}
