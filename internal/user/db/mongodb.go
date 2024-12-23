package db

import (
	"context"

	"github.com/morheus9/go_rest/internal/user"
	"github.com/morheus9/go_rest/pkg/logging"
)

type db struct {
}

func (d *db) Create(context.Context, user.User) (string, error) {
	panic("implement me")
}

func (d *db) FindOne(context.Context, string) (user.User, error) {
	panic("implement me")
}

func (d *db) Update(context.Context, user.User) error {
	panic("implement me")
}

func (d *db) Delete(context.Context, string) error {
	panic("implement me")
}
func NewStorage(collection string, logger *logging.Logger) user.Storage {
	return &db{}
}
