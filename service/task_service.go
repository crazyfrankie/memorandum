package service

import (
	"context"
	"time"

	"memorandum/pkg/ctl"
	"memorandum/pkg/util"
	"memorandum/repository/db/dao"
	"memorandum/repository/db/model"
)

type TaskService interface {
	TaskCreate(ctx context.Context, req *model.CreateTaskReq) (err error)
	DeleteTask(ctx context.Context, req *model.DeleteTaskReq) (resp interface{}, err error)
	UpdateTask(ctx context.Context, req *model.UpdateTaskReq) (resp interface{}, err error)
	SearchTask(ctx context.Context, req *model.SearchTaskReq) (resp interface{}, err error)
	ShowTask(ctx context.Context, req *model.ShowTaskReq) (resp interface{}, err error)
	ListTask(ctx context.Context, req *model.ListTasksReq) (resp interface{}, total int64, err error)
}

type taskService struct {
	taskRepo dao.TaskRepository
	userRepo dao.UserRepository
}

func NewTaskService(taskRepo dao.TaskRepository, userRepo dao.UserRepository) TaskService {
	return &taskService{taskRepo: taskRepo, userRepo: userRepo}
}

func (t *taskService) TaskCreate(ctx context.Context, req *model.CreateTaskReq) (err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return err
	}

	user, err := t.userRepo.FindByID(u.ID)
	if err != nil {
		util.LogrusObj.Info(err)
		return err
	}

	task := &model.Task{
		User:      *user,
		Uid:       user.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}

	err = t.taskRepo.NewTask(task)
	if err != nil {
		util.LogrusObj.Info(err)
		return err
	}

	return nil
}

func (t *taskService) DeleteTask(ctx context.Context, req *model.DeleteTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	err = t.taskRepo.DeleteTaskById(u.ID, req.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	return resp, nil
}

func (t *taskService) UpdateTask(ctx context.Context, req *model.UpdateTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	err = t.taskRepo.UpdateTask(u.ID, req)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	return req, nil
}

func (t *taskService) SearchTask(ctx context.Context, req *model.SearchTaskReq) (resp interface{}, err error) {
	user, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	tasks, err := t.taskRepo.SearchTask(user.ID, req.Info)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	taskRespList := make([]*model.TaskResp, 0)
	for _, task := range tasks {
		taskRespList = append(taskRespList, &model.TaskResp{
			ID:        task.ID,
			Title:     task.Title,
			Content:   task.Content,
			Status:    task.Status,
			View:      task.View(),
			CreatedAt: task.CreatedAt.Unix(),
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
		})
	}

	return taskRespList, nil
}

func (t *taskService) ListTask(ctx context.Context, req *model.ListTasksReq) (resp interface{}, total int64, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	var tasks []*model.Task
	tasks, total, err = t.taskRepo.ListTask(req.Start, req.Limit, u.ID)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	return tasks, total, nil
}

func (t *taskService) ShowTask(ctx context.Context, req *model.ShowTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	task, err := t.taskRepo.FindTaskById(u.ID, req.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	return task, nil
}
