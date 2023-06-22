package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

func (i *impl) GetByUserID(ctx context.Context, userID int) ([]model.Task, error) {
	var (
		resp  []model.Task
		query = `SELECT id, user_id, summary, completed_at FROM tasks WHERE user_id = ?;`
	)

	err := i.mySQLReader.SelectContext(ctx, &resp, query, userID)
	if err != nil {
		shared.Logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
