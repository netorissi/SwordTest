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

func TestGet(t *testing.T) {
	cases := map[string]struct {
		output      []model.Task
		outputErr   error
		prepareMock func(context.Context, sqlmock.Sqlmock)
	}{
		"success": {
			output: []model.Task{{ID: 1, UserID: 2}},
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`SELECT id, user_id, summary, completed_at FROM tasks;`)
				s.
					ExpectQuery(query).
					WillReturnError(nil).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "summary", "completed_at"}).AddRow(1, 2, "", nil))
			},
		},
		"error": {
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`SELECT id, user_id, summary, completed_at FROM tasks;`)
				s.
					ExpectQuery(query).
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

			output, err := u.Get(ctx)
			assert.Equal(t, cs.output, output)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
