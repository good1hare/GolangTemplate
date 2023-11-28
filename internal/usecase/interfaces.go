package usecase

import (
	"github.com/good1hare/GolangTemplate/internal/entity"
)

type (
	User interface {
		GetUser(int) (entity.User, error)
		CreateOrUpdateUser(entity.User) (entity.User, error)
		DeleteUser(int) error
	}

	UserRepo interface {
		FindUser(int) (entity.User, error)
		SaveUser(entity.User) (entity.User, error)
		UpdateUser(entity.User) (entity.User, error)
		DeleteUser(int) error
	}
)
