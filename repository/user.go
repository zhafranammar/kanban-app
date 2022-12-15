package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	users := entity.User{}
	err := r.db.Where("id = ?",id).First(&users)
	if err != nil{
		if errors.Is(err.Error, gorm.ErrRecordNotFound){
			return users,nil
		}
		return users, err.Error
	}
	return users, err.Error // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	users := entity.User{}
	err := r.db.Where("email = ?",email).First(&users)
	if err != nil{
		if errors.Is(err.Error, gorm.ErrRecordNotFound){
			return users,nil
		}
		return users, err.Error
	}
	return users, err.Error // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.Create(&user)
	return user, err.Error // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	return entity.User{}, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	users := entity.User{}
	err :=  r.db.Where("id = ?", id).Delete(&users).Error
	return err // TODO: replace this
}
