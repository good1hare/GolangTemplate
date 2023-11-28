package repo

import (
	"context"
	"database/sql"
	queries "github.com/good1hare/GolangTemplate/db"
	"github.com/good1hare/GolangTemplate/internal/entity"
)

type UserRepo struct {
	*sql.DB
	*queries.Queries
}

func NewUserRepo(sql *sql.DB) *UserRepo {
	q := queries.New(sql) // Create a new queries.

	return &UserRepo{sql, q}
}

func ConvertUser(userOld queries.User) entity.User {
	entityUser := entity.User{
		Id:        int(userOld.ID),
		Name:      userOld.Name.String,
		Phone:     int(userOld.Phone.Int32),
		CreatedAt: userOld.CreatedAt.Time,
		UpdatedAt: userOld.UpdatedAt.Time,
	}
	return entityUser
}

func (r *UserRepo) FindUser(userId int) (entity.User, error) {
	ctx := context.Background()
	qUser, err := r.Queries.GetUser(ctx, int64(userId))
	if err != nil {
		return entity.User{}, err
	}
	return ConvertUser(qUser), nil
}

func (r *UserRepo) SaveUser(user entity.User) (entity.User, error) {
	ctx := context.Background()
	result, err := r.Queries.CreateUser(ctx, queries.CreateUserParams{Name: sql.NullString{String: user.Name, Valid: true}, Phone: sql.NullInt32{Int32: int32(user.Phone), Valid: true}})
	if err != nil {
		return entity.User{}, err
	}

	insertedUserId, err := result.LastInsertId()
	if err != nil {
		return entity.User{}, err
	}

	qUser, err := r.Queries.GetUser(ctx, insertedUserId)
	if err != nil {
		return entity.User{}, err
	}
	return ConvertUser(qUser), nil
}

func (r *UserRepo) UpdateUser(user entity.User) (entity.User, error) {
	return user, nil
}

func (r *UserRepo) DeleteUser(userId int) error {
	ctx := context.Background()
	err := r.Queries.DeleteUser(ctx, int64(userId))
	if err != nil {
		return err
	}
	return nil
}
