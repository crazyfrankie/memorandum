package dao

import (
	"memorandum/repository/db/model"
)

type TaskRepository interface {
	NewTask(task *model.Task) error
	DeleteTaskById(uid, id uint) error
	UpdateTask(uId uint, req *model.UpdateTaskReq) error
	SearchTask(uId uint, info string) (tasks []*model.Task, err error)
	FindTaskByUserId(uId, id uint) (r *model.Task, err error)
	FindTaskById(uId, id uint) (r *model.Task, err error)
	ListTask(start, limit int, userId uint) (r []*model.Task, total int64, err error)
}

type taskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

func (t *taskRepository) NewTask(task *model.Task) error {
	return DB.Create(&task).Error
}

func (t *taskRepository) DeleteTaskById(uid, ID uint) error {
	r, err := t.FindTaskByUserId(uid, ID)
	if err != nil {
		return err
	}
	return DB.Delete(&r).Error
}

func (t *taskRepository) UpdateTask(uId uint, req *model.UpdateTaskReq) error {
	task := new(model.Task)
	err := DB.Model(&model.Task{}).Where("id = ? AND uid=?", req.ID, uId).First(&task).Error
	if err != nil {
		return err
	}

	if req.Status != 0 {
		task.Status = req.Status
	}

	if req.Title != "" {
		task.Title = req.Title
	}

	if req.Content != "" {
		task.Content = req.Content
	}

	return DB.Save(task).Error
}

func (t *taskRepository) FindTaskByUserId(uId, id uint) (r *model.Task, err error) {
	err = DB.Model(&model.Task{}).Where("id = ? AND uid = ?", id, uId).First(&r).Error
	return
}

func (t *taskRepository) FindTaskById(uId, id uint) (r *model.Task, err error) {
	err = DB.Model(&model.Task{}).Where("id = ? AND uid = ?", id, uId).First(&r).Error

	return
}

func (t *taskRepository) ListTask(start, limit int, userId uint) (r []*model.Task, total int64, err error) {
	err = DB.Model(&model.Task{}).Preload("User").Where("uid = ?", userId).
		Count(&total).
		Limit(limit).Offset((start - 1) * limit).
		Find(&r).Error

	return
}

func (t *taskRepository) SearchTask(uId uint, info string) (tasks []*model.Task, err error) {
	err = DB.Where("uid=?", uId).Preload("User").First(&tasks).Error
	if err != nil {
		return
	}

	err = DB.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?",
		"%"+info+"%", "%"+info+"%").Find(&tasks).Error

	return
}
