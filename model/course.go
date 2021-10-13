package model

type Course struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Count string `json:"count"`
}

type CourseDetail struct {
	Id       int      `json:"id"`
	Category Category `json:"category"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Count    string   `json:"count"`
}

type CategoryDetail struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count string `json:"count"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CourseUpdate struct {
	Id         int    `json:"Id"`
	CategoryId int    `json:"course_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Count      string `json:"count"`
}

type StatisticResponse struct {
	TotalUser       int `json:"total_user"`
	TotalCourse     int `json:"total_course"`
	TotalCourseFree int `json:"total_course_free"`
}
