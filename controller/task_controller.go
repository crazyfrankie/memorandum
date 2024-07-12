package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"memorandum/consts"
	"memorandum/pkg/util"
	"memorandum/repository/db/dao"
	"memorandum/repository/db/model"
	"memorandum/service"
)

var taskService service.TaskService = service.NewTaskService(dao.NewTaskRepository(), dao.NewUserRepository())

// CreateTaskHandler 创建任务
// @Summary 创建任务
// @Description 创建任务
// @Tags 任务
// @Accept json
// @Produce json
// @Param task body model.CreateTaskReq true "创建任务请求体"
// @Success 200 {object} model.Task
// @Failure 400 {object} ctl.ErrResponse
// @Failure 500 {object} ctl.ErrResponse
// @Router /v1/tasks [post]
func CreateTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.CreateTaskReq
		if err := c.ShouldBind(&task); err != nil {
			util.LogrusObj.Info(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		if err := taskService.TaskCreate(c.Request.Context(), &task); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		}

		c.JSON(http.StatusOK, task)
	}
}

// ShowTaskHandler 获取单个任务
// @Summary 获取单个任务
// @Description 获取单个任务
// @Tags 任务
// @Produce json
// @Param id path int true "任务ID"
// @Success 200 {object} model.Task
// @Failure 400 {object} ctl.ErrResponse
// @Failure 500 {object} ctl.ErrResponse
// @Router /v1/tasks/{id} [get]
func ShowTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.ShowTaskReq
		if err := c.ShouldBindQuery(&task); err != nil {
			util.LogrusObj.Info(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		resp, err := taskService.ShowTask(c.Request.Context(), &task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

// ListHandler 获取分页任务列表
// @Summary 获取分页任务列表
// @Description 获取分页任务列表
// @Tags 任务
// @Produce json
// @Param start query int true "起始页"
// @Param limit query int true "每页限制数"
// @Success 200 {object} ctl.DataResponse
// @Failure 400 {object} ctl.ErrResponse
// @Failure 500 {object} ctl.ErrResponse
// @Router /v1/tasks [get]
func ListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ListTasksReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Info(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		if req.Limit == 0 {
			req.Limit = consts.BasePageLimit
		}
		resp, total, err := taskService.ListTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, DataResponse(resp, total))
	}
}

// DeleteHandler 删除任务
// @Summary 删除任务
// @Description 删除任务
// @Tags 任务
// @Produce json
// @Param id path int true "任务ID"
// @Success 200 {object} string
// @Failure 400 {object} ctl.ErrResponse
// @Failure 500 {object} ctl.ErrResponse
// @Router /v1/tasks/{id} [delete]
func DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.DeleteTaskReq
		if err := c.ShouldBindQuery(&req); err != nil {
			util.LogrusObj.Info(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		_, err := taskService.DeleteTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, "delete successfully")
	}
}

// TaskUpdate DeleteHandler 删除任务
// @Summary 删除任务
// @Description 删除任务
// @Tags 任务
// @Produce json
// @Param id path int true "任务ID"
// @Success 200 {object} ctl.DataResponse
// @Failure 400 {object} ctl.ErrResponse
// @Failure 500 {object} ctl.ErrResponse
// @Router /v1/tasks/{id} [delete]
func TaskUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.UpdateTaskReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Info(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		resp, err := taskService.UpdateTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, DataResponse("update successfully", resp))
	}
}

// TaskSearch 搜索关键词任务
// @Summary 搜索关键词任务
// @Description 搜索关键词任务
// @Tags 任务
// @Produce json
// @Param info query string true "搜索关键词"
// @Success 200 {object} []model.Task
// @Failure 400 {object} ctl.ErrResponse
// @Failure 500 {object} ctl.ErrResponse
// @Router /v1/tasks/search [get]
func TaskSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.SearchTaskReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Info(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		resp, err := taskService.SearchTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
