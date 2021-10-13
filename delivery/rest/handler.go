package rest

import (
	"net/http"
	"strconv"

	"github.com/egaevan/online-learning/model"
	"github.com/egaevan/online-learning/usecase"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	CourseUsecae usecase.CourseUsecae
	UserUsecae   usecase.UserUsecae
}

type responseError struct {
	Message string `json:"message"`
}

const (
	isAdmin int = 1
)

func NewHandler(e *echo.Echo, courseUsecae usecase.CourseUsecae, userUsecae usecase.UserUsecae) {
	handler := &Handler{
		CourseUsecae: courseUsecae,
		UserUsecae:   userUsecae,
	}

	// Routing User
	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)
	e.DELETE("/user/:userID", handler.DeleteUser, JwtVerify)

	// Routing Course
	e.GET("/course", handler.GetCourse)
	e.GET("/course/:courseID", handler.GetDetailCourse)
	e.GET("/course-search", handler.SearchCourse)
	e.GET("/course-sort", handler.SortCourse)
	e.POST("/course", handler.SendCourse, JwtVerify)
	e.PATCH("/course/:courseID", handler.UpdateCourse, JwtVerify)
	e.DELETE("/course/:courseID", handler.DeleteCourse, JwtVerify)
	e.GET("/category", handler.GetCategory)
	e.GET("/category/:limit", handler.GetPopularCategory)

	e.GET("/statistic", handler.GetStatistic, JwtVerify)

}

func (h *Handler) GetDetailCourse(c echo.Context) error {
	courseIDParam := c.Param("courseID")

	if courseIDParam == "" {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	res, err := h.CourseUsecae.GetDetailCourse(c.Request().Context(), courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) GetCourse(c echo.Context) error {

	res, err := h.CourseUsecae.GetCourse(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) SendCourse(c echo.Context) error {
	dataReq := model.Course{}

	userInfo := c.Get("user").(*model.Token)

	if userInfo.Role != isAdmin {
		// unauthorized
		return echo.ErrUnauthorized
	}

	if err := c.Bind(&dataReq); err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid data request",
		})
		return echo.ErrBadRequest
	}

	res, err := h.CourseUsecae.SendCourse(c.Request().Context(), dataReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) UpdateCourse(c echo.Context) error {
	dataReq := model.CourseUpdate{}
	courseIDParam := c.Param("courseID")

	userInfo := c.Get("user").(*model.Token)

	if userInfo.Role != isAdmin {
		// unauthorized
		return echo.ErrUnauthorized
	}

	if courseIDParam == "" {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	if err := c.Bind(&dataReq); err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid data request",
		})
		return echo.ErrBadRequest
	}

	res, err := h.CourseUsecae.UpdateCourse(c.Request().Context(), dataReq, courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteCourse(c echo.Context) error {
	courseIDParam := c.Param("courseID")

	userInfo := c.Get("user").(*model.Token)

	if userInfo.Role != isAdmin {
		// unauthorized
		return echo.ErrUnauthorized
	}

	if courseIDParam == "" {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	err = h.CourseUsecae.DeleteCourse(c.Request().Context(), courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, responseError{
		Message: "Course has been deleted",
	})
}

func (h *Handler) SearchCourse(c echo.Context) error {
	search := c.QueryParam("search")

	if search == "" {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	res, err := h.CourseUsecae.SearchCourse(c.Request().Context(), search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) SortCourse(c echo.Context) error {
	sort := c.QueryParam("sort")

	if sort == "" {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	res, err := h.CourseUsecae.SortCourse(c.Request().Context(), sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) GetStatistic(c echo.Context) error {
	res, err := h.CourseUsecae.GetStatistic(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c echo.Context) error {
	dataReq := model.User{}
	if err := c.Bind(&dataReq); err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid data request",
		})
		return echo.ErrBadRequest
	}

	user, err := h.UserUsecae.Login(c.Request().Context(), dataReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: err.Error(),
		})
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "logged in",
		"token":   user.Token,
	})
}

func (h *Handler) Register(c echo.Context) error {
	dataReq := model.User{}
	if err := c.Bind(&dataReq); err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid data request",
		})

		return echo.ErrBadRequest
	}

	err := h.UserUsecae.CreateUser(c.Request().Context(), dataReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: err.Error(),
		})
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "success")
}

func (h *Handler) DeleteUser(c echo.Context) error {
	userIDParam := c.Param("userID")

	userInfo := c.Get("user").(*model.Token)

	if userInfo.Role != isAdmin {
		// unauthorized
		return echo.ErrUnauthorized
	}

	// admin
	if userIDParam == "" {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	err = h.UserUsecae.DeleteUser(c.Request().Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, responseError{
		Message: "User has been deleted",
	})
}

func (h *Handler) GetCategory(c echo.Context) error {

	res, err := h.CourseUsecae.GetCategory(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) GetPopularCategory(c echo.Context) error {
	limitParam := c.Param("limit")

	if limitParam == "" {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid parameter",
		})

		return echo.ErrBadRequest
	}

	res, err := h.CourseUsecae.GetPopularCategory(c.Request().Context(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}
