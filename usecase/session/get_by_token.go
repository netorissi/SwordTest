package session

import (
	"context"

	"github.com/netorissi/SwordTest/model"
)

func (i *impl) GetByToken(ctx context.Context, token string) (*model.Session, error) {
	return i.repository.Session.GetByToken(ctx, token)
}
