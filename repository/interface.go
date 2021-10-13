package repository

import (
	"context"

	"github.com/egaevan/online-learning/model"
)

type CourseRepository interface {
	FindOne(context.Context, int) (*model.CourseDetail, error)
	Fetch(context.Context) ([]model.Course, error)
	Store(context.Context, model.Course) error
	Update(context.Context, model.CourseUpdate, int) error
	Delete(context.Context, int) error
	Search(context.Context, string) ([]model.Course, error)
	Sort(context.Context, string) ([]model.Course, error)
	Statistic(ctx context.Context) (*model.StatisticResponse, error)
	FetchCategory(context.Context) ([]model.CategoryDetail, error)
	FetchPopularCategory(context.Context, int) ([]model.CategoryDetail, error)
}

type UserRepository interface {
	FindOne(context.Context, string, string) (model.User, error)
	Fetch(context.Context) error
	Store(context.Context, model.User) error
	Update(context.Context) error
	Delete(context.Context, int) error
}
