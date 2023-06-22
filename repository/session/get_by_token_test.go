package session_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository/session"
	"github.com/netorissi/SwordTest/test"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	cases := map[string]struct {
		input       string
		output      *model.Session
		outputErr   error
		prepareMock func(context.Context, sqlmock.Sqlmock)
	}{
		"success": {
			input:  "123",
			output: &model.Session{UserID: 2},
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`SELECT user_id, access_token FROM sessions WHERE access_token = ?;`)
				s.
					ExpectQuery(query).
					WithArgs("123").
					WillReturnError(nil).
					WillReturnRows(sqlmock.NewRows([]string{"user_id", "access_token"}).AddRow(2, ""))
			},
		},
		"error": {
			input:     "123",
			outputErr: test.ErrAny,
			prepareMock: func(ctx context.Context, s sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(`SELECT user_id, access_token FROM sessions WHERE access_token = ?;`)
				s.
					ExpectQuery(query).
					WithArgs("123").
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

			u := session.New(session.Options{
				MySQLReader: db,
				MySQLWriter: db,
			})

			cs.prepareMock(ctx, mk)

			output, err := u.GetByToken(ctx, cs.input)
			assert.Equal(t, cs.output, output)
			assert.Equal(t, cs.outputErr, err)
		})
	}
}
