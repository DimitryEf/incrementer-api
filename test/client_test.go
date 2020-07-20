package test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DimitryEf/incrementer-api/api"
	"github.com/DimitryEf/incrementer-api/tool"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
)

func TestHealthCheck(t *testing.T) {
	want := `{"alive": true}`
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	got := string(body)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func clearDB() error {
	//Порт указан в docker-compose.yml
	db, err := sql.Open("postgres", "port=5433 host=localhost user=postgres password=mysecretpassword dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE incrementer SET num = $1, maximum_value = $2, step_value = $3", 0, tool.MaximumInt64, 1)
	if err != nil {
		return err
	}
	return nil
}

func TestIncrementerBasicRequests(t *testing.T) {

	// Перед и после тестов очищаем базу
	err := clearDB()
	if err != nil {
		t.Fatalf("error on clear db: %v", err)
	}
	/*defer func() {
		err := clearDB()
		if err != nil {
			t.Fatalf("error on clear db: %v", err)
		}
	}()*/

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
		{
			method: "SetParams",
			request: &api.Request{
				StepValue: 3,
			},
			want: &api.Empty{},
		},
		{
			method:  "GetNumber",
			request: &api.Empty{},
			want:    &api.Response{Num: 1},
		},
		{
			method:  "IncrementNumber",
			request: &api.Empty{},
			want:    &api.Empty{},
		},
		{
			method:  "GetNumber",
			request: &api.Empty{},
			want:    &api.Response{Num: 4},
		},
		{
			method: "SetParams",
			request: &api.Request{
				MaximumValue: 2,
			},
			want: &api.Empty{},
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
			want:    &api.Response{Num: 0},
		},
		{
			method: "SetParams",
			request: &api.Request{
				StepValue: 1,
			},
			want: &api.Empty{},
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
		{
			method:  "IncrementNumber",
			request: &api.Empty{},
			want:    &api.Empty{},
		},
		{
			method:  "GetNumber",
			request: &api.Empty{},
			want:    &api.Response{Num: 2},
		},
		{
			method:  "IncrementNumber",
			request: &api.Empty{},
			want:    &api.Empty{},
		},
		{
			method:  "GetNumber",
			request: &api.Empty{},
			want:    &api.Response{Num: 0},
		},
		{
			method: "SetParams",
			request: &api.Request{
				MaximumValue: 10,
				StepValue:    3,
			},
			want: &api.Empty{},
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
			want:    &api.Response{Num: 3},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v ", tc.method), func(tt *testing.T) {

			switch tc.method {
			case "GetNumber":
				got, err := client.GetNumber(context.Background(), tc.request.(*api.Empty))
				if err != nil {
					t.Fatalf("fail to dial: %v", err)
				}

				if got.Num != tc.want.(*api.Response).Num {
					t.Fatalf("got %v; want %v", got.Num, tc.want.(*api.Response).Num)
				}
			case "IncrementNumber":
				got, err := client.IncrementNumber(context.Background(), tc.request.(*api.Empty))
				if err != nil {
					t.Fatalf("fail to dial: %v", err)
				}
				if got.Status != true {
					t.Fatalf("got %v; want %v", false, true)
				}
			case "SetParams":
				got, err := client.SetParams(context.Background(), tc.request.(*api.Request))
				if err != nil {
					t.Fatalf("fail to dial: %v", err)
				}
				if got.Status != true {
					t.Fatalf("got %v; want %v", false, true)
				}

			}

		})
	}

}
