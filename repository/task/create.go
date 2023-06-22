package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

func (i *impl) Create(ctx context.Context, task model.Task) error {
	query := `INSERT INTO tasks (user_id, summary, completed_at) VALUES (?, ?, ?);`

	_, err := i.mySQLReader.ExecContext(ctx, query, task.UserID, task.Summary, task.CompletedAt)
	if err != nil {
		shared.Logger.Error(err.Error())
		return err
	}

	return nil
}
