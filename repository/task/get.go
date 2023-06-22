package task

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

func (i *impl) Get(ctx context.Context) ([]model.Task, error) {
	var (
		resp  []model.Task
		query = `SELECT id, user_id, summary, completed_at FROM tasks;`
	)

	err := i.mySQLReader.SelectContext(ctx, &resp, query)
	if err != nil {
		shared.Logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
