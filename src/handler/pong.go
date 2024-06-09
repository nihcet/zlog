package handler

import (
	"context"
	"fmt"

	"github.com/nihcet/zlog/src/model"
	error_model "github.com/nihcet/zlog/src/model/error"
)

func (Handler) Pong(_ context.Context, _ any) (*model.PingPong, error) {
	fmt.Println("pong")
	return nil, error_model.NotFoundError
	// return nil, nil
	// return &model.PingPong{
	// 	Name:   "pong",
	// 	Number: 10110,
	// }, nil
}
