package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	tasks := []entity.Task{}
	err := r.db.Where("user_id = ?", id).Find(&tasks).Error
	return tasks, err // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	res := r.db.Create(&task).Error
	return task.ID, res // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	tasks := entity.Task{}
	err := r.db.Where("id = ?", id).First(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	tasks := []entity.Task{}
	err := r.db.Where("category_id = ?", catId).Find(&tasks).Error
	return tasks, err // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	err := r.db.Where("id = ?", task.ID).Updates(&task).Error
	return err // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&entity.Task{}).Error
	return err // TODO: replace this
}
