package task

import (
	"context"
	"errors"
	"fmt"

	"github.com/netorissi/SwordTest/event"
	"github.com/netorissi/SwordTest/shared"
)

var (
	ErrUserInvalid   = errors.New("user is invalid for this operation")
	ErrTaskCompleted = errors.New("the task already completed")
)

func (i *impl) Complete(ctx context.Context, taskID, userID int) error {
	task, err := i.repository.Task.GetByID(ctx, taskID)
	if err != nil {
		return err
	}

	if userID != task.UserID {
		return ErrUserInvalid
	}

	if task.CompletedAt != nil {
		return ErrTaskCompleted
	}

	now := shared.TimeNow()
	task.CompletedAt = &now

	if err := i.repository.Task.Update(ctx, *task); err != nil {
		return err
	}

	msg := fmt.Sprintf("The tech %d performed the task %d on date %v", userID, taskID, now)
	go i.broker.Publish(event.EVENT_NOTIFY_MANAGERS, []byte(msg))

	return nil
}
