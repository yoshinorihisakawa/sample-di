package repository

import (
	"context"
	"database/sql"
	"github.com/yoshinorihisakawa/sample-di/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User)
}

type userRepository struct {
	db sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{*db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) {
	ur.db.Query("")
}
