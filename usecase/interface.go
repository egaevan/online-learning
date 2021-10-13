package usecase

import (
	"context"

	"github.com/egaevan/online-learning/model"
)

type CourseUsecae interface {
	GetDetailCourse(context.Context, int) (*model.CourseDetail, error)
	GetCourse(context.Context) ([]model.Course, error)
	SendCourse(context.Context, model.Course) (*model.Course, error)
	UpdateCourse(context.Context, model.CourseUpdate, int) (*model.CourseUpdate, error)
	DeleteCourse(context.Context, int) error
	SearchCourse(context.Context, string) ([]model.Course, error)
	SortCourse(context.Context, string) ([]model.Course, error)
	GetStatistic(ctx context.Context) (*model.StatisticResponse, error)
	GetCategory(context.Context) ([]model.CategoryDetail, error)
	GetPopularCategory(context.Context, int) ([]model.CategoryDetail, error)
}

type UserUsecae interface {
	Login(ctx context.Context, user model.User) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	DeleteUser(context.Context, int) error
}
