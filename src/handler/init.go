package handler

import (
	"context"

	"github.com/nihcet/zlog/src/model"
)

type HandlerInterface interface {
	Ping(context.Context, any) (*model.PingPong, error)
	Pong(context.Context, any) (*model.PingPong, error)
}

type Handler struct{}

func InitialHandler() HandlerInterface {
	return &Handler{}
}
