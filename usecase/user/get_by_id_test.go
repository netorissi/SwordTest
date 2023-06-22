package user_test

import (
	"context"
	"testing"

	"github.com/netorissi/SwordTest/mocks"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository"
	"github.com/netorissi/SwordTest/test"
	"github.com/netorissi/SwordTest/usecase/user"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	cases := map[string]struct {
		input       int
		output      *model.User
		outputErr   error
		prepareMock func(context.Context, *mocks.MockUserRepository)
	}{
		"success": {
			input:  2,
			output: &model.User{ID: 2},
			prepareMock: func(ctx context.Context, msr *mocks.MockUserRepository) {
				msr.EXPECT().GetByID(ctx, 2).Return(&model.User{ID: 2}, nil)
			},
		},
		"error": {
			input:     2,
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, msr *mocks.MockUserRepository) {
				msr.EXPECT().GetByID(ctx, 2).Return(nil, test.ErrAny)
			},
		},
	}

	for n, cs := range cases {
		t.Run(n, func(t *testing.T) {
			var (
				ctrl, ctx = test.NewController(t)
				mockStore = mocks.NewMockUserRepository(ctrl)
			)

			s := user.New(&repository.Repository{
				User: mockStore,
			})

			cs.prepareMock(ctx, mockStore)

			output, err := s.GetByID(ctx, cs.input)
			assert.Equal(t, cs.output, output)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
