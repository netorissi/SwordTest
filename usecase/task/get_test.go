package task_test

import (
	"context"
	"testing"

	"github.com/netorissi/SwordTest/mocks"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository"
	"github.com/netorissi/SwordTest/test"
	"github.com/netorissi/SwordTest/usecase/task"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	cases := map[string]struct {
		output      []model.Task
		outputErr   error
		prepareMock func(context.Context, *mocks.MockTaskRepository)
	}{
		"success": {
			output: []model.Task{{ID: 1, UserID: 2}},
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository) {
				msr.EXPECT().Get(ctx).Return([]model.Task{{ID: 1, UserID: 2}}, nil)
			},
		},
		"error": {
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository) {
				msr.EXPECT().Get(ctx).Return(nil, test.ErrAny)
			},
		},
	}

	for n, cs := range cases {
		t.Run(n, func(t *testing.T) {
			var (
				ctrl, ctx = test.NewController(t)
				mockStore = mocks.NewMockTaskRepository(ctrl)
			)

			s := task.New(&repository.Repository{
				Task: mockStore,
			}, nil)

			cs.prepareMock(ctx, mockStore)

			output, err := s.Get(ctx)
			assert.Equal(t, cs.output, output)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
