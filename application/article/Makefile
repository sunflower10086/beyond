gen-api:
	goctl api go --dir=./api --api ./api/article.api

gen-rpc:
	goctl rpc protoc ./rpc/article.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc

gen-model:
	goctl model mysql datasource --dir ./rpc/internal/model --table article --url "root:123456@tcp(127.0.0.1:3306)/beyond_article"