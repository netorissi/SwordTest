package user_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository/user"
	"github.com/netorissi/SwordTest/test"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	cases := map[string]struct {
		input       int
		output      *model.User
		outputErr   error
		prepareMock func(context.Context, sqlmock.Sqlmock)
	}{
		"success": {
			input:  1,
			output: &model.User{ID: 1},
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`SELECT id, role FROM users WHERE id = ?;`)
				s.
					ExpectQuery(query).
					WithArgs(1).
					WillReturnError(nil).
					WillReturnRows(sqlmock.NewRows([]string{"id", "role"}).AddRow(1, ""))
			},
		},
		"error": {
			input:     1,
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`SELECT id, role FROM users WHERE id = ?;`)
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

			u := user.New(user.Options{
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
