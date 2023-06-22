package task_test

import (
	"context"
	"testing"

	"github.com/netorissi/SwordTest/mocks"
	"github.com/netorissi/SwordTest/repository"
	"github.com/netorissi/SwordTest/test"
	"github.com/netorissi/SwordTest/usecase/task"
	"github.com/stretchr/testify/assert"
)

func TestDel(t *testing.T) {
	cases := map[string]struct {
		input       int
		outputErr   error
		prepareMock func(context.Context, *mocks.MockTaskRepository)
	}{
		"success": {
			input: 1,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository) {
				msr.EXPECT().Del(ctx, 1).Return(nil)
			},
		},
		"error": {
			input:     1,
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository) {
				msr.EXPECT().Del(ctx, 1).Return(test.ErrAny)
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

			err := s.Del(ctx, cs.input)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
