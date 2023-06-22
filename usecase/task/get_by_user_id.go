package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
)

func (i *impl) GetByUserID(ctx context.Context, userID int) ([]model.Task, error) {
	return i.repository.Task.GetByUserID(ctx, userID)
}
