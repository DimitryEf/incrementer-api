package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

/*func TestClient(t *testing.T) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := NewIncrementerClient(conn)
	empty := &Empty{}
	response, err := client.IncrementNumber(context.Background(), empty)
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	t.Logf("response: %v", response.Status)

	number, err := client.GetNumber(context.Background(), empty)
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	t.Errorf("response: %v", number.Num)
}*/

func TestClient2(t *testing.T) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := NewIncrementerClient(conn)

	testCases := []struct {
		method  string
		request interface{}
		want    interface{}
	}{
		{
			method:  "GetNumber",
			request: &Empty{},
			want:    &Response{Num: 0},
		},
		{
			method:  "GetNumber",
			request: &Empty{},
			want:    &Response{Num: 0},
		},
		{
			method:  "IncrementNumber",
			request: &Empty{},
			want:    &Empty{},
		},
		{
			method:  "GetNumber",
			request: &Empty{},
			want:    &Response{Num: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v ", tc.method), func(t *testing.T) {

			//t.Errorf("Printf:'%v'", tc.want)

			switch tc.method {
			case "GetNumber":
				got, err := client.GetNumber(context.Background(), tc.request.(*Empty))
				if err != nil {
					t.Fatalf("fail to dial: %v", err)
				}

				if got.Num != tc.want.(*Response).Num {
					t.Errorf("got %v; want %v", got.Num, tc.want.(*Response).Num)
				}
			case "IncrementNumber":
				got, err := client.IncrementNumber(context.Background(), tc.request.(*Empty))
				if err != nil {
					t.Fatalf("fail to dial: %v", err)
				}
				if got.Status != true {
					t.Errorf("got %v; want %v", false, true)
				}
			case "SetParams":
				got, err := client.SetParams(context.Background(), tc.request.(*Request))
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
