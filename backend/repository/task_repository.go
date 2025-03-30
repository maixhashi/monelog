package repository

import (
	"fmt"
	"monelog/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(userId uint) ([]model.Task, error)
	GetTaskById(userId uint, taskId uint) (model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(userId uint) ([]model.Task, error) {
	var tasks []model.Task
	if err := tr.db.Where("user_id=?", userId).Order("created_at").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *taskRepository) GetTaskById(userId uint, taskId uint) (model.Task, error) {
	var task model.Task
	result := tr.db.Where("user_id=?", userId).First(&task, taskId)
	if result.Error != nil {
		return model.Task{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Task{}, fmt.Errorf("task not found")
	}
	return task, nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	return tr.db.Create(task).Error
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(&model.Task{}).Clauses(clause.Returning{}).
		Where("id=? AND user_id=?", taskId, userId).
		Updates(map[string]interface{}{
			"title": task.Title,
		}).First(task)
	
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}