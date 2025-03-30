package usecase

import (
	"monelog/model"
	"monelog/repository"
	"monelog/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(request model.TaskRequest) (model.TaskResponse, error)
	UpdateTask(request model.TaskRequest, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks, err := tu.tr.GetAllTasks(userId)
	if err != nil {
		return nil, err
	}
	
	responses := make([]model.TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = task.ToResponse()
	}
	return responses, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task, err := tu.tr.GetTaskById(userId, taskId)
	if err != nil {
		return model.TaskResponse{}, err
	}
	return task.ToResponse(), nil
}

func (tu *taskUsecase) CreateTask(request model.TaskRequest) (model.TaskResponse, error) {
	if err := tu.tv.ValidateTaskRequest(request); err != nil {
		return model.TaskResponse{}, err
	}
	
	task := request.ToModel()
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	
	return task.ToResponse(), nil
}

func (tu *taskUsecase) UpdateTask(request model.TaskRequest, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tv.ValidateTaskRequest(request); err != nil {
		return model.TaskResponse{}, err
	}
	
	task := request.ToModel()
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	
	return task.ToResponse(), nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	return tu.tr.DeleteTask(userId, taskId)
}