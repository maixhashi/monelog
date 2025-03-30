package validator

import (
	"fmt"
	"monelog/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	ValidateTaskRequest(task model.TaskRequest) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) ValidateTaskRequest(task model.TaskRequest) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, model.TaskTitleMaxLength).Error(
				fmt.Sprintf("limited max %d char", model.TaskTitleMaxLength),
			),
		),
	)
}