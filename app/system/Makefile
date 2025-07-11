init:
	goctl env -w GOCTL_EXPERIMENTAL=off
	go install github.com/zeromicro/go-zero/tools/goctl@latest

api:
	goctl api go --api desc/desc.api --style "go_zero" --dir .

grpc:
	goctl rpc protoc desc/desc.proto --style "go_zero" --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --client=true -m

db:
	gentool -c "./gen.yaml"

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags no_k8s -o app-system-linux-amd64 system.go

