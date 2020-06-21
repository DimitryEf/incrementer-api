package main

import "context"

type GrpcApi struct {
	Inc *Incrementer
}

func NewGrpcApi(inc *Incrementer) *GrpcApi {
	return &GrpcApi{
		Inc: inc,
	}
}

func (g *GrpcApi) GetNumber(ctx context.Context, request *Empty) (*Response, error) {
	num, err := g.Inc.GetNumber()
	if err != nil {
		return nil, err
	}
	return &Response{Num: num}, nil
}

func (g *GrpcApi) IncrementNumber(ctx context.Context, request *Empty) (*Empty, error) {
	err := g.Inc.IncrementNumber()
	if err != nil {
		return nil, err
	}
	return &Empty{Status: true}, nil
}

func (g *GrpcApi) SetParams(ctx context.Context, request *Request) (*Empty, error) {
	err := g.Inc.SetParams(request.MaximumValue, request.StepValue)
	if err != nil {
		return nil, err
	}
	return &Empty{Status: true}, nil
}
