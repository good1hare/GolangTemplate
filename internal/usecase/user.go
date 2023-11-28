package usecase

import (
	"github.com/good1hare/GolangTemplate/internal/entity"
)

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u *UserUseCase) GetUser(userId int) (entity.User, error) {
	user, err := u.repo.FindUser(userId)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserUseCase) CreateOrUpdateUser(user entity.User) (entity.User, error) {
	resultUser, err := u.repo.FindUser(user.Id)
	if err == nil {
		resultUser, err = u.repo.UpdateUser(user)
		return resultUser, err
	}
	resultUser, err = u.repo.SaveUser(user)
	if err != nil {
		return resultUser, err
	}
	return resultUser, nil
}

func (u *UserUseCase) DeleteUser(userId int) error {
	err := u.repo.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
