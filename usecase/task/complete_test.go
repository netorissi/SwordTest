package task_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/netorissi/SwordTest/event"
	"github.com/netorissi/SwordTest/mocks"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository"
	"github.com/netorissi/SwordTest/shared"
	"github.com/netorissi/SwordTest/test"
	"github.com/netorissi/SwordTest/usecase/task"
	"github.com/stretchr/testify/assert"
)

func TestComplete(t *testing.T) {
	shared.TimeNow = func() time.Time {
		return time.Date(2023, time.June, 22, 0, 0, 0, 0, time.Local)
	}

	ts := shared.TimeNow()

	cases := map[string]struct {
		input1      int
		input2      int
		outputErr   error
		prepareMock func(context.Context, *mocks.MockTaskRepository, *mocks.MockMessageBroker, *sync.WaitGroup)
	}{
		"success": {
			input1: 1,
			input2: 2,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository, b *mocks.MockMessageBroker, wg *sync.WaitGroup) {
				msr.EXPECT().GetByID(ctx, 1).Return(&model.Task{ID: 1, UserID: 2}, nil)
				msr.EXPECT().Update(ctx, model.Task{ID: 1, UserID: 2, CompletedAt: &ts}).Return(nil)
				wg.Add(1)
				msg := fmt.Sprintf("The tech 2 performed the task 1 on date %v", ts)
				b.EXPECT().Publish(event.EVENT_NOTIFY_MANAGERS, []byte(msg)).DoAndReturn(func(arg0, arg1 interface{}) error {
					defer wg.Done()
					return nil
				})
			},
		},
		"error": {
			input1:    1,
			input2:    2,
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository, b *mocks.MockMessageBroker, wg *sync.WaitGroup) {
				msr.EXPECT().GetByID(ctx, 1).Return(&model.Task{ID: 1, UserID: 2}, nil)
				msr.EXPECT().Update(ctx, model.Task{ID: 1, UserID: 2, CompletedAt: &ts}).Return(test.ErrAny)
			},
		},
		"error user id invalid": {
			input1:    1,
			input2:    3,
			outputErr: task.ErrUserInvalid,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository, b *mocks.MockMessageBroker, wg *sync.WaitGroup) {
				msr.EXPECT().GetByID(ctx, 1).Return(&model.Task{ID: 1, UserID: 2}, nil)
			},
		},
		"error task already completed": {
			input1:    1,
			input2:    2,
			outputErr: task.ErrTaskCompleted,
			prepareMock: func(ctx context.Context, msr *mocks.MockTaskRepository, b *mocks.MockMessageBroker, wg *sync.WaitGroup) {
				msr.EXPECT().GetByID(ctx, 1).Return(&model.Task{ID: 1, UserID: 2, CompletedAt: &ts}, nil)
			},
		},
	}

	for n, cs := range cases {
		t.Run(n, func(t *testing.T) {
			var (
				ctrl, ctx  = test.NewController(t)
				mockStore  = mocks.NewMockTaskRepository(ctrl)
				mockBroker = mocks.NewMockMessageBroker(ctrl)
				wg         = &sync.WaitGroup{}
			)

			s := task.New(&repository.Repository{
				Task: mockStore,
			}, mockBroker)

			cs.prepareMock(ctx, mockStore, mockBroker, wg)

			err := s.Complete(ctx, cs.input1, cs.input2)

			wg.Wait()

			assert.Equal(t, cs.outputErr, err)
		})
	}
}
