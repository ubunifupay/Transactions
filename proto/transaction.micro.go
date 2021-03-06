// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/transaction.proto

package transaction

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Transaction service

type TransactionService interface {
	GetTransactions(ctx context.Context, in *TransactionRequest, opts ...client.CallOption) (*TransactionReply, error)
}

type transactionService struct {
	c    client.Client
	name string
}

func NewTransactionService(name string, c client.Client) TransactionService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "transaction"
	}
	return &transactionService{
		c:    c,
		name: name,
	}
}

func (c *transactionService) GetTransactions(ctx context.Context, in *TransactionRequest, opts ...client.CallOption) (*TransactionReply, error) {
	req := c.c.NewRequest(c.name, "Transaction.GetTransactions", in)
	out := new(TransactionReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Transaction service

type TransactionHandler interface {
	GetTransactions(context.Context, *TransactionRequest, *TransactionReply) error
}

func RegisterTransactionHandler(s server.Server, hdlr TransactionHandler, opts ...server.HandlerOption) error {
	type transaction interface {
		GetTransactions(ctx context.Context, in *TransactionRequest, out *TransactionReply) error
	}
	type Transaction struct {
		transaction
	}
	h := &transactionHandler{hdlr}
	return s.Handle(s.NewHandler(&Transaction{h}, opts...))
}

type transactionHandler struct {
	TransactionHandler
}

func (h *transactionHandler) GetTransactions(ctx context.Context, in *TransactionRequest, out *TransactionReply) error {
	return h.TransactionHandler.GetTransactions(ctx, in, out)
}
