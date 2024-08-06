package controllers

import (
	"strconv"
	data "task-management-api/data"
	dtos "task-management-api/dtos"

	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	taskDtos := data.GetTask()
	ctx.JSON(200, taskDtos)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	taskDto, err := data.GetTaskById(taskId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, taskDto)
}

func CreateTask(ctx *gin.Context) {
	var createTaskDto dtos.CreateTaskDto
	err := ctx.ShouldBindJSON(&createTaskDto)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := data.CreateTask(createTaskDto)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{
		"id":      id,
		"message": "Task created successfully"})
}

func UpdateTask(ctx *gin.Context) {
	var updateTaskDto dtos.UpdateTaskDto
	err := ctx.ShouldBindJSON(&updateTaskDto)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = data.UpdateTask(updateTaskDto)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = data.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Task deleted successfully"})
}
