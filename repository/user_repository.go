package repository

import (
	"context"
	"github.com/yoshinorihisakawa/sample-di/model"
	"github.com/yoshinorihisakawa/sample-di/database"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User)
}

type userRepository struct {
	db database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{*db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) {
	ur.db.Con().Query("")
}
