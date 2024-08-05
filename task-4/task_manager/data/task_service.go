package data

import (
	"errors"
	"time"

	dtos "task-management-api/dtos"
	model "task-management-api/models"
)

var Tasks = []model.Task{
	{
		ID:          1,
		Title:       "Finish Report",
		Description: "Work on the detailed report for the project.",
		Deadline:    time.Now().AddDate(0, 0, 3),
		Status:      model.InProgress,
	},
	{
		ID:          2,
		Title:       "Code Review",
		Description: "Review and refactor existing codebase.",
		Deadline:    time.Now().AddDate(0, 0, 5),
		Status:      model.InProgress,
	},
	{
		ID:          3,
		Title:       "Bug Fixing",
		Description: "Identify and fix bugs reported by QA.",
		Deadline:    time.Now().AddDate(0, 0, 2),
		Status:      model.InProgress,
	},
	{
		ID:          4,
		Title:       "Meeting Prep",
		Description: "Prepare agenda and materials for the team meeting.",
		Deadline:    time.Now().AddDate(0, 0, 4),
		Status:      model.InProgress,
	},
	{
		ID:          5,
		Title:       "Deploy New Feature",
		Description: "Implement and deploy new feature according to specifications.",
		Deadline:    time.Now().AddDate(0, 0, 7),
		Status:      model.InProgress,
	},
}

const layout = "2006-01-02T15:04:05Z07:00"

func GetTask() []dtos.TaskDto {

	data := []dtos.TaskDto{}

	for _, task := range Tasks {
		taskDto := dtos.TaskDto{
			ID:          int(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Deadline:    task.Deadline.String(),
			Status:      int(task.Status),
		}
		data = append(data, taskDto)
	}
	return data

}

func GetTaskById(id int) (dtos.TaskDto, error) {

	data := dtos.TaskDto{}

	for _, task := range Tasks {
		if int(task.ID) == id {
			data := dtos.TaskDto{
				ID:          int(task.ID),
				Title:       task.Title,
				Description: task.Description,
				Deadline:    task.Deadline.String(),
				Status:      int(task.Status),
			}
			return data, nil
		}
	}
	return data, errors.New("no such id")

}

func CreateTask(createTaskDto dtos.CreateTaskDto) (uint16, error) {
	deadline, err := time.Parse(layout, createTaskDto.Deadline)
	if err != nil {
		return 0, err
	}
	newId := uint16(len(Tasks) + 1)
	newTask := model.Task{
		ID:          newId,
		Title:       createTaskDto.Title,
		Description: createTaskDto.Description,
		Deadline:    deadline,
		Status:      model.Status(createTaskDto.Status),
	}
	Tasks = append(Tasks, newTask)
	return newId, nil
}

func UpdateTask(updateTaskDto dtos.UpdateTaskDto) error {
	for i, task := range Tasks {
		if int(task.ID) == updateTaskDto.ID {
			deadline, err := time.Parse(layout, updateTaskDto.Deadline)
			if err != nil {
				return err
			}
			Tasks[i].Title = updateTaskDto.Title
			Tasks[i].Description = updateTaskDto.Description
			Tasks[i].Deadline = deadline
			Tasks[i].Status = model.Status(updateTaskDto.Status)
			return nil
		}
	}
	return errors.New("no such id")
}

func DeleteTask(id int) error {
	for i, task := range Tasks {
		if int(task.ID) == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("no such id")
}
