package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
)

func (i *impl) Create(ctx context.Context, task model.Task) error {
	return i.repository.Task.Create(ctx, task)
}
