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

func TestGetByID(t *testing.T) {
	cases := map[string]struct {
		input       int
		output      *model.Task
		outputErr   error
		prepareMock func(context.Context, sqlmock.Sqlmock)
	}{
		"success": {
			input:  1,
			output: &model.Task{ID: 1, UserID: 2},
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`SELECT id, user_id, summary, completed_at FROM tasks WHERE id = ?;`)
				s.
					ExpectQuery(query).
					WithArgs(1).
					WillReturnError(nil).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "summary", "completed_at"}).AddRow(1, 2, "", nil))
			},
		},
		"error": {
			input:     1,
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`SELECT id, user_id, summary, completed_at FROM tasks WHERE id = ?;`)
				s.
					ExpectQuery(query).
					WithArgs(1).
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

			output, err := u.GetByID(ctx, cs.input)
			assert.Equal(t, cs.output, output)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
