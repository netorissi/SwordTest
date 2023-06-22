package test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

var ErrAny = errors.New("any error")

func NewController(t *testing.T) (*gomock.Controller, context.Context) {
	return gomock.NewController(t), context.Background()
}
