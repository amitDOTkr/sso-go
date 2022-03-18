binary:
	CGO_ENABLED=0 go build main.go

image:
	docker build -t amitdotkr/sso-go:dev .

push:
	docker push amitdotkr/sso-go:dev

docker:
	CGO_ENABLED=0 go build main.go && \
	docker build -t amitdotkr/sso-go:dev . && \
	docker push amitdotkr/sso-go:dev && \
	rm main

pb:
	protoc --go_out=. --go-grpc_out=. ./src/protos/*.proto