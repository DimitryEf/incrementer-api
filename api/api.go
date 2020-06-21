package api

import (
	"context"
	"github.com/DimitryEf/incrementer-api/usecase"
)

// Api структура с rpc-хендлерами
type Api struct {
	Inc *usecase.Incrementer
}

// NewGrpcApi - конструктор для Api структуры
func NewGrpcApi(inc *usecase.Incrementer) *Api {
	return &Api{
		Inc: inc,
	}
}

// GetNumber - получить значение инкрементора
func (g *Api) GetNumber(ctx context.Context, request *Empty) (*Response, error) {
	num, err := g.Inc.GetNumber()
	if err != nil {
		return nil, err
	}
	return &Response{Num: num}, nil
}

// IncrementNumber - увеличить значение инкрементора
func (g *Api) IncrementNumber(ctx context.Context, request *Empty) (*Empty, error) {
	err := g.Inc.IncrementNumber()
	if err != nil {
		return nil, err
	}
	return &Empty{Status: true}, nil
}

// SetParams - установить параметры инкрементора: максимальное значение и значение шага
func (g *Api) SetParams(ctx context.Context, request *Request) (*Empty, error) {
	err := g.Inc.SetParams(request.MaximumValue, request.StepValue)
	if err != nil {
		return nil, err
	}
	return &Empty{Status: true}, nil
}
