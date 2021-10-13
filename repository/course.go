package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/egaevan/online-learning/model"

	log "github.com/sirupsen/logrus"
)

type Course struct {
	DB *sql.DB
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &Course{
		DB: db,
	}
}

func (c *Course) FindOne(ctx context.Context, courseID int) (*model.CourseDetail, error) {
	query := `
			SELECT 
				course.id,
				course.name,
				course.price,
				course.count,
				category.id,
				category.name
			FROM 
				course
			JOIN
				category ON course.category_id = category.id
			WHERE
				course.id = ? AND flag_aktif = 1`

	course := model.CourseDetail{}

	err := c.DB.QueryRowContext(ctx, query, courseID).Scan(&course.Id, &course.Name, &course.Price, &course.Count, &course.Category.Id, &course.Category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found %s", err.Error())
		}
		return nil, err
	}

	return &course, nil
}

func (c *Course) Fetch(ctx context.Context) (result []model.Course, err error) {
	query := `
			SELECT 
				id,
				name,
				price,
				count
			FROM 
				course
			WHERE
				flag_aktif = 1`

	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]model.Course, 0)

	for rows.Next() {
		t := model.Course{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.Price,
			&t.Count,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (c *Course) Store(ctx context.Context, course model.Course) error {
	query := `
				INSERT INTO course
					(id,name,price,count)
				VALUES
					(?, ?, ?, ?, ?)
			`

	_, err := c.DB.ExecContext(ctx, query,
		course.Id, course.Name, course.Price, course.Count)

	if err != nil {
		return err
	}

	return nil
}

func (c *Course) Update(ctx context.Context, course model.CourseUpdate, courseID int) error {
	query := `
				UPDATE 
					course
				SET
					category_id = ?, 
					name = ?, 
					price = ?, 
					count = ?
				WHERE
					id = ?
			`

	_, err := c.DB.ExecContext(ctx, query,
		course.CategoryId, course.Name, course.Price, course.Count, courseID)

	if err != nil {
		return err
	}

	return nil

}

func (c *Course) Delete(ctx context.Context, courseID int) error {
	query := `
				UPDATE 
					course
				SET
					flag_aktif = 0
				WHERE
					id = ?
			`

	_, err := c.DB.ExecContext(ctx, query, courseID)

	if err != nil {
		return err
	}

	return nil
}

func (c *Course) Search(ctx context.Context, search string) (result []model.Course, err error) {
	query := `
			SELECT
				id,
				name,
				price,
				count
			FROM
				course
			WHERE
				name LIKE '%%%s%%' AND flag_aktif = 1`

	rows, err := c.DB.QueryContext(ctx, fmt.Sprintf(query, search))
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]model.Course, 0)

	for rows.Next() {
		t := model.Course{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.Price,
			&t.Count,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (c *Course) Sort(ctx context.Context, sort string) (result []model.Course, err error) {
	query := `
			SELECT
				id,
				name,
				price,
				count
			FROM
				course
			WHERE
				flag_aktif = 1 %s`

	rows, err := c.DB.QueryContext(ctx, fmt.Sprintf(query, sort))
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]model.Course, 0)

	for rows.Next() {
		t := model.Course{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.Price,
			&t.Count,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (c *Course) Statistic(ctx context.Context) (*model.StatisticResponse, error) {
	query := `SELECT COUNT(id) FROM user WHERE role = 1`
	query2 := `SELECT COUNT(id) FROM course`
	query3 := `SELECT COUNT(id) FROM course WHERE price = 0`

	statistic := model.StatisticResponse{}

	err := c.DB.QueryRowContext(ctx, query).Scan(&statistic.TotalUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found %s", err.Error())
		}
		return nil, err
	}

	err = c.DB.QueryRowContext(ctx, query2).Scan(&statistic.TotalCourse)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found %s", err.Error())
		}
		return nil, err
	}

	err = c.DB.QueryRowContext(ctx, query3).Scan(&statistic.TotalCourseFree)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found %s", err.Error())
		}
		return nil, err
	}

	return &statistic, nil

}

func (c *Course) FetchCategory(ctx context.Context) (result []model.CategoryDetail, err error) {
	query := `
			SELECT 
				id,
				name,
				count
			FROM 
				category`

	rows, err := c.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]model.CategoryDetail, 0)

	for rows.Next() {
		t := model.CategoryDetail{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.Count,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (c *Course) FetchPopularCategory(ctx context.Context, limit int) (result []model.CategoryDetail, err error) {
	query := `
			SELECT 
				id,
				name,
				count
			FROM 
				category
			ORDER BY
				count DESC
			LIMIT
				?`

	rows, err := c.DB.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]model.CategoryDetail, 0)

	for rows.Next() {
		t := model.CategoryDetail{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.Count,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}
