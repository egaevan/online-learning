package usecase

import (
	"context"
	"fmt"

	"github.com/egaevan/online-learning/model"
	"github.com/egaevan/online-learning/repository"
	log "github.com/sirupsen/logrus"
)

const (
	lowPrice  = "ORDER BY price ASC"
	highPrice = "ORDER BY price DESC"
	freePrice = "AND price = 0"
)

type responseError struct {
	Message string `json:"message"`
}

type Course struct {
	CourseRepo repository.CourseRepository
}

func NewCourse(courseRepo repository.CourseRepository) CourseUsecae {
	return &Course{
		CourseRepo: courseRepo,
	}
}

func (c *Course) GetDetailCourse(ctx context.Context, CourseID int) (*model.CourseDetail, error) {

	prod, err := c.CourseRepo.FindOne(ctx, CourseID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return prod, nil
}

func (c *Course) GetCourse(ctx context.Context) ([]model.Course, error) {

	course, err := c.CourseRepo.Fetch(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	courseList := []model.Course{}

	for _, v := range course {
		var course model.Course

		course.Id = v.Id
		course.Name = v.Name
		course.Price = v.Price
		course.Count = v.Count

		courseList = append(courseList, course)
	}

	return courseList, nil
}

func (c *Course) SearchCourse(ctx context.Context, search string) ([]model.Course, error) {

	course, err := c.CourseRepo.Search(ctx, search)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	courseList := []model.Course{}

	for _, v := range course {
		var course model.Course

		course.Id = v.Id
		course.Name = v.Name
		course.Price = v.Price
		course.Count = v.Count

		courseList = append(courseList, course)
	}

	return courseList, nil
}

func (c *Course) SortCourse(ctx context.Context, sort string) ([]model.Course, error) {
	if sort == "high" {
		sort = highPrice
	} else if sort == "low" {
		sort = lowPrice
	} else if sort == "free" {
		sort = freePrice
	} else {
		return nil, fmt.Errorf("internal error")
	}

	course, err := c.CourseRepo.Sort(ctx, sort)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	courseList := []model.Course{}

	for _, v := range course {
		var course model.Course

		course.Id = v.Id
		course.Name = v.Name
		course.Price = v.Price
		course.Count = v.Count

		courseList = append(courseList, course)
	}

	return courseList, nil
}

func (c *Course) SendCourse(ctx context.Context, course model.Course) (*model.Course, error) {

	err := c.CourseRepo.Store(ctx, course)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &course, nil
}

func (c *Course) UpdateCourse(ctx context.Context, course model.CourseUpdate, courseID int) (*model.CourseUpdate, error) {

	err := c.CourseRepo.Update(ctx, course, courseID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &course, nil
}

func (c *Course) DeleteCourse(ctx context.Context, courseID int) error {

	err := c.CourseRepo.Delete(ctx, courseID)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (c *Course) GetStatistic(ctx context.Context) (*model.StatisticResponse, error) {

	stat, err := c.CourseRepo.Statistic(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return stat, nil
}

func (c *Course) GetCategory(ctx context.Context) ([]model.CategoryDetail, error) {

	category, err := c.CourseRepo.FetchCategory(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	categoryList := []model.CategoryDetail{}

	for _, v := range category {
		var category model.CategoryDetail

		category.Id = v.Id
		category.Name = v.Name
		category.Count = v.Count

		categoryList = append(categoryList, category)
	}

	return categoryList, nil
}

func (c *Course) GetPopularCategory(ctx context.Context, limit int) ([]model.CategoryDetail, error) {

	category, err := c.CourseRepo.FetchPopularCategory(ctx, limit)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	categoryList := []model.CategoryDetail{}

	for _, v := range category {
		var category model.CategoryDetail

		category.Id = v.Id
		category.Name = v.Name
		category.Count = v.Count

		categoryList = append(categoryList, category)
	}

	return categoryList, nil
}
