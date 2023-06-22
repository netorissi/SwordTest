package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
)

func (i *impl) Get(ctx context.Context) ([]model.Task, error) {
	return i.repository.Task.Get(ctx)
}
