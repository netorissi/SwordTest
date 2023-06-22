package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

func (i *impl) GetByID(ctx context.Context, taskID int) (*model.Task, error) {
	var (
		resp  = new(model.Task)
		query = `SELECT id, user_id, summary, completed_at FROM tasks WHERE id = ?;`
	)

	err := i.mySQLReader.GetContext(ctx, resp, query, taskID)
	if err != nil {
		shared.Logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
