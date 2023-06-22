package task

import (
	"context"

	"github.com/netorissi/SwordTest/shared"
)

func (i *impl) Del(ctx context.Context, taskID int) error {
	query := `DELETE FROM tasks WHERE id = ?;`

	_, err := i.mySQLReader.ExecContext(ctx, query, taskID)
	if err != nil {
		shared.Logger.Error(err.Error())
		return err
	}

	return nil
}
