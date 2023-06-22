package session_test

import (
	"context"
	"testing"

	"github.com/netorissi/SwordTest/mocks"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository"
	"github.com/netorissi/SwordTest/test"
	"github.com/netorissi/SwordTest/usecase/session"
	"github.com/stretchr/testify/assert"
)

func TestGetByToken(t *testing.T) {
	cases := map[string]struct {
		input       string
		output      *model.Session
		outputErr   error
		prepareMock func(context.Context, *mocks.MockSessionRepository)
	}{
		"success": {
			input:  "123",
			output: &model.Session{UserID: 2},
			prepareMock: func(ctx context.Context, msr *mocks.MockSessionRepository) {
				msr.EXPECT().GetByToken(ctx, "123").Return(&model.Session{UserID: 2}, nil)
			},
		},
		"error": {
			input:     "123",
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, msr *mocks.MockSessionRepository) {
				msr.EXPECT().GetByToken(ctx, "123").Return(nil, test.ErrAny)
			},
		},
	}

	for n, cs := range cases {
		t.Run(n, func(t *testing.T) {
			var (
				ctrl, ctx = test.NewController(t)
				mockStore = mocks.NewMockSessionRepository(ctrl)
			)

			s := session.New(&repository.Repository{
				Session: mockStore,
			})

			cs.prepareMock(ctx, mockStore)

			output, err := s.GetByToken(ctx, cs.input)
			assert.Equal(t, cs.output, output)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
