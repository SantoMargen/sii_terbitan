package user

import (
	"context"
	"siap_app/internal/app/entity"
	"siap_app/internal/app/entity/user"
)

type userUC interface {
	CreateUser(ctx context.Context, input user.RegisterRequest) error
	CreateUserByAdmin(ctx context.Context, input user.RegisterByAdminRequest) error
	LoginUser(ctx context.Context, ipAddress string, input user.LoginRequest) (user.ResponseLogin, error)
	LogoutUser(ctx context.Context, email string) error
	UpdateRoleUser(ctx context.Context, userId int, input user.UpdateRoleRequest) error
	UpdatePasswordUser(ctx context.Context, userId int, input user.UpdatePaswordeRequest) error
	GetUsers(ctx context.Context, input entity.Pagination) ([]user.User, int64, error)
}
