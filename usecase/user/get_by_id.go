package user

import (
	"context"

	"github.com/netorissi/SwordTest/model"
)

func (i *impl) GetByID(ctx context.Context, userID int) (*model.User, error) {
	return i.repository.User.GetByID(ctx, userID)
}
