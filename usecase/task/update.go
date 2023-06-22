package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
)

func (i *impl) Update(ctx context.Context, task model.Task) error {
	current, err := i.repository.Task.GetByID(ctx, task.ID)
	if err != nil {
		return err
	}

	if current.UserID != task.UserID {
		return ErrUserInvalid
	}

	current.Summary = task.Summary

	return i.repository.Task.Update(ctx, *current)
}
