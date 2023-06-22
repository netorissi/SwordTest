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

func TestUpdate(t *testing.T) {
	cases := map[string]struct {
		input       model.Task
		outputErr   error
		prepareMock func(context.Context, *mocks.MockTaskRepository)
	}{
		"success": {
			input: model.Task{ID: 1, UserID: 1, Summary: "123"},
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository) {
				msr.EXPECT().GetByID(ctx, 1).Return(&model.Task{ID: 1, UserID: 1}, nil)
				msr.EXPECT().Update(ctx, model.Task{ID: 1, UserID: 1, Summary: "123"}).Return(nil)
			},
		},
		"error": {
			input:     model.Task{ID: 1, UserID: 1, Summary: "123"},
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository) {
				msr.EXPECT().GetByID(ctx, 1).Return(&model.Task{ID: 1, UserID: 1}, nil)
				msr.EXPECT().Update(ctx, model.Task{ID: 1, UserID: 1, Summary: "123"}).Return(test.ErrAny)
			},
		},
		"error user invalid": {
			input:     model.Task{ID: 1, UserID: 2, Summary: "123"},
			outputErr: task.ErrUserInvalid,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository) {
				msr.EXPECT().GetByID(ctx, 1).Return(&model.Task{ID: 1, UserID: 1}, nil)
			},
		},
		"error get task by id": {
			input:     model.Task{ID: 1, UserID: 1, Summary: "123"},
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository) {
				msr.EXPECT().GetByID(ctx, 1).Return(nil, test.ErrAny)
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

			err := s.Update(ctx, cs.input)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
