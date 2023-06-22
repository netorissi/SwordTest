package task

import (
	"context"
)

func (i *impl) Del(ctx context.Context, taskID int) error {
	return i.repository.Task.Del(ctx, taskID)
}
