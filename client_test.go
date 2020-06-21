package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"testing"
)

func TestClient(t *testing.T) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := NewIncrementerClient(conn)
	request := &Request{
		MaximumValue: 20,
		StepValue:    2,
	}
	response, err := client.SetParams(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Num)
	grpclog.Fatalf("fail to dial: %v", response.Num)
}
