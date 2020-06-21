package test

import (
	"context"
	"fmt"
	"github.com/DimitryEf/incrementer-api/api"
	"google.golang.org/grpc"
	"testing"
)

func TestIncrementerBasicRequests(t *testing.T) {

	// Предварительная настройка клиента
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := api.NewIncrementerClient(conn)

	// Тесты
	testCases := []struct {
		method  string
		request interface{}
		want    interface{}
	}{
		{
			method:  "GetNumber",
			request: &api.Empty{},
			want:    &api.Response{Num: 0},
		},
		{
			method:  "GetNumber",
			request: &api.Empty{},
			want:    &api.Response{Num: 0},
		},
		{
			method:  "IncrementNumber",
			request: &api.Empty{},
			want:    &api.Empty{},
		},
		{
			method:  "GetNumber",
			request: &api.Empty{},
			want:    &api.Response{Num: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v ", tc.method), func(t *testing.T) {

			switch tc.method {
			case "GetNumber":
				got, err := client.GetNumber(context.Background(), tc.request.(*api.Empty))
				if err != nil {
					t.Fatalf("fail to dial: %v", err)
				}

				if got.Num != tc.want.(*api.Response).Num {
					t.Errorf("got %v; want %v", got.Num, tc.want.(*api.Response).Num)
				}
			case "IncrementNumber":
				got, err := client.IncrementNumber(context.Background(), tc.request.(*api.Empty))
				if err != nil {
					t.Fatalf("fail to dial: %v", err)
				}
				if got.Status != true {
					t.Errorf("got %v; want %v", false, true)
				}
			case "SetParams":
				got, err := client.SetParams(context.Background(), tc.request.(*api.Request))
				if err != nil {
					t.Fatalf("fail to dial: %v", err)
				}
				if got.Status != true {
					t.Errorf("got %v; want %v", false, true)
				}

			}

		})
	}

}
