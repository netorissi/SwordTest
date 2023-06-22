package user

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

func (i *impl) GetByID(ctx context.Context, userID int) (*model.User, error) {
	var (
		resp  = new(model.User)
		query = `SELECT id, role FROM users WHERE id = ?;`
	)

	err := i.mySQLReader.GetContext(ctx, resp, query, userID)
	if err != nil {
		shared.Logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
