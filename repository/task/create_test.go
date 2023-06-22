package task_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository/task"
	"github.com/netorissi/SwordTest/test"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	cases := map[string]struct {
		input       model.Task
		outputErr   error
		prepareMock func(context.Context, sqlmock.Sqlmock)
	}{
		"success": {
			input: model.Task{
				UserID:  1,
				Summary: "test",
			},
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`INSERT INTO tasks (user_id, summary, completed_at) VALUES (?, ?, ?);`)
				s.
					ExpectExec(query).
					WithArgs(1, "test", nil).
					WillReturnError(nil).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
		},
		"error": {
			input: model.Task{
				UserID:  1,
				Summary: "test",
			},
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`INSERT INTO tasks (user_id, summary, completed_at) VALUES (?, ?, ?);`)
				s.
					ExpectExec(query).
					WithArgs(1, "test", nil).
					WillReturnError(test.ErrAny)
			},
		},
	}

	for n, cs := range cases {
		t.Run(n, func(t *testing.T) {
			var (
				_, ctx = test.NewController(t)
				db, mk = test.NewMySQL()
			)

			u := task.New(task.Options{
				MySQLReader: db,
				MySQLWriter: db,
			})

			cs.prepareMock(ctx, mk)

			err := u.Create(ctx, cs.input)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
