package controller

import (
	"monelog/model"
	"monelog/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// TaskController タスク関連のAPIを管理するコントローラー
// @title Blog CMS API
// @version 1.0
// @description ブログCMSのAPI
// @BasePath /
type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

// GetAllTasks ユーザーのすべてのタスクを取得
// @Summary ユーザーのタスク一覧を取得
// @Description ログインユーザーのすべてのタスクを取得する
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} model.TaskResponse
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func (tc *taskController) GetAllTasks(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	tasksRes, err := tc.tu.GetAllTasks(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tasksRes)
}

// GetTaskById 指定されたIDのタスクを取得
// @Summary 特定のタスクを取得
// @Description 指定されたIDのタスクを取得する
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskId path int true "タスクID"
// @Success 200 {object} model.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{taskId} [get]
func (tc *taskController) GetTaskById(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	id := c.Param("taskId")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID"})
	}
	
	taskRes, err := tc.tu.GetTaskById(userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, taskRes)
}

// CreateTask 新しいタスクを作成
// @Summary 新しいタスクを作成
// @Description ユーザーの新しいタスクを作成する
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.TaskRequest true "タスク情報"
// @Success 201 {object} model.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func (tc *taskController) CreateTask(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	var request model.TaskRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	taskRes, err := tc.tu.CreateTask(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, taskRes)
}

// UpdateTask 既存のタスクを更新
// @Summary タスクを更新
// @Description 指定されたIDのタスクを更新する
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskId path int true "タスクID"
// @Param task body model.TaskRequest true "更新するタスク情報"
// @Success 200 {object} model.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{taskId} [put]
func (tc *taskController) UpdateTask(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	id := c.Param("taskId")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID"})
	}
	
	var request model.TaskRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	request.UserId = userId
	taskRes, err := tc.tu.UpdateTask(request, userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, taskRes)
}

// DeleteTask タスクを削除
// @Summary タスクを削除
// @Description 指定されたIDのタスクを削除する
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskId path int true "タスクID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{taskId} [delete]
func (tc *taskController) DeleteTask(c echo.Context) error {
	userId, err := getUserIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}
	
	id := c.Param("taskId")
	taskId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid task ID"})
	}
	
	err = tc.tu.DeleteTask(userId, uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// getUserIdFromToken はJWTトークンからユーザーIDを取得する関数
func getUserIdFromToken(c echo.Context) (uint, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))
	return userId, nil
}
