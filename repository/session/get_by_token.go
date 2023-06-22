package session

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

func (i *impl) GetByToken(ctx context.Context, token string) (*model.Session, error) {
	var (
		resp  = new(model.Session)
		query = `SELECT user_id, access_token FROM sessions WHERE access_token = ?;`
	)

	err := i.mySQLReader.GetContext(ctx, resp, query, token)
	if err != nil {
		shared.Logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
