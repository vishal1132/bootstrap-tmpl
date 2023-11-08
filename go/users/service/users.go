package service

import (
	"context"

	"{{ module_name }}/users/model"
)

type userDB interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, entityType, entityVal string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type userService struct {
	db userDB
}

func NewUserService(db userDB) *userService {
	return &userService{
		db: db,
	}
}

func (s *userService) UpdateUserCompanyID(ctx context.Context, userID string, companyID string) (*model.User, error) {
	panic("implement me")
}

func (s *userService) UpdateUser(ctx context.Context, userID string, req *model.UserUpdateRequest) (*model.User, error) {
	panic("implement me")
}

func (s *userService) CreateUserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic("implement me")
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	panic("implement me")
}

func (s *userService) GetUser(ctx context.Context, email string) (*model.User, error) {
	panic("implement me")
}
