package database

import (
	"context"

	"github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	Bun *bun.DB
}

func New(bun *bun.DB) UserRepository {
	return UserRepository{Bun: bun}
}

func (r *UserRepository) Get(ctx context.Context, id string) (user.User, error) {
	um := new(User)

	err := r.Bun.NewSelect().
		Model(um).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return user.User{}, err
	}

	u := user.User{
		ID:       um.ID,
		Username: um.Username,
		Roles:    make([]user.Role, 0),
	}

	return u, nil
}
