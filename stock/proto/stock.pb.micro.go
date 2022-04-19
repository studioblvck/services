// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/stock.proto

package stock

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Stock service

func NewStockEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Stock service

type StockService interface {
	Quote(ctx context.Context, in *QuoteRequest, opts ...client.CallOption) (*QuoteResponse, error)
	Price(ctx context.Context, in *PriceRequest, opts ...client.CallOption) (*PriceResponse, error)
	History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error)
}

type stockService struct {
	c    client.Client
	name string
}

func NewStockService(name string, c client.Client) StockService {
	return &stockService{
		c:    c,
		name: name,
	}
}

func (c *stockService) Quote(ctx context.Context, in *QuoteRequest, opts ...client.CallOption) (*QuoteResponse, error) {
	req := c.c.NewRequest(c.name, "Stock.Quote", in)
	out := new(QuoteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockService) Price(ctx context.Context, in *PriceRequest, opts ...client.CallOption) (*PriceResponse, error) {
	req := c.c.NewRequest(c.name, "Stock.Price", in)
	out := new(PriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockService) History(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryResponse, error) {
	req := c.c.NewRequest(c.name, "Stock.History", in)
	out := new(HistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stock service

type StockHandler interface {
	Quote(context.Context, *QuoteRequest, *QuoteResponse) error
	Price(context.Context, *PriceRequest, *PriceResponse) error
	History(context.Context, *HistoryRequest, *HistoryResponse) error
}

func RegisterStockHandler(s server.Server, hdlr StockHandler, opts ...server.HandlerOption) error {
	type stock interface {
		Quote(ctx context.Context, in *QuoteRequest, out *QuoteResponse) error
		Price(ctx context.Context, in *PriceRequest, out *PriceResponse) error
		History(ctx context.Context, in *HistoryRequest, out *HistoryResponse) error
	}
	type Stock struct {
		stock
	}
	h := &stockHandler{hdlr}
	return s.Handle(s.NewHandler(&Stock{h}, opts...))
}

type stockHandler struct {
	StockHandler
}

func (h *stockHandler) Quote(ctx context.Context, in *QuoteRequest, out *QuoteResponse) error {
	return h.StockHandler.Quote(ctx, in, out)
}

func (h *stockHandler) Price(ctx context.Context, in *PriceRequest, out *PriceResponse) error {
	return h.StockHandler.Price(ctx, in, out)
}

func (h *stockHandler) History(ctx context.Context, in *HistoryRequest, out *HistoryResponse) error {
	return h.StockHandler.History(ctx, in, out)
}
