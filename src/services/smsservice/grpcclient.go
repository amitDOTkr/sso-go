package smsservice

import (
	"log"

	// "github.com/amitdotkr/sso/sso-go/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GrpcClient() {
	cc, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("error: %v", err)
	}

	defer cc.Close()

	// gc := pb.NewOtpServiceClient(cc)
	// gc.OtpSend()
}
