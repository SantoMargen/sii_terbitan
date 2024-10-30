package category

import (
	"context"
	"siap_app/internal/app/entity/category"
)

type categoryUC interface {
	CreateCategory(ctx context.Context, input category.RequestCategory) error
	GetCategoryById(ctx context.Context, id int) (category.ResponseCategory, error)
	DeleteCategory(ctx context.Context, id int) error
	UpdateCategory(ctx context.Context, input category.RequestCategoryUpdate) error
}
