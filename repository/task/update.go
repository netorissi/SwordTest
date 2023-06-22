package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

func (i *impl) Update(ctx context.Context, task model.Task) error {
	query := `UPDATE tasks SET user_id = ?, summary = ?, completed_at = ? WHERE id = ?;`

	_, err := i.mySQLReader.ExecContext(ctx, query, task.UserID, task.Summary, task.CompletedAt, task.ID)
	if err != nil {
		shared.Logger.Error(err.Error())
		return err
	}

	return nil
}
