package handler

import (
	"context"
	"fmt"

	"github.com/nihcet/zlog/src/model"
)

func (Handler) Ping(_ context.Context, _ any) (*model.PingPong, error) {
	fmt.Println("ping")
	return &model.PingPong{
		Name:   "ping",
		Number: 42,
	}, nil
}
