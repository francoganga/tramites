package database

import (
	"context"

	"github.com/francoganga/go_reno2/ent"
	"github.com/francoganga/go_reno2/pkg/domain/user"
)



type UserRepository struct {
    orm *ent.Client
}

func (r *UserRepository) Get(ctx context.Context, id string) (user.User, error) {

}
