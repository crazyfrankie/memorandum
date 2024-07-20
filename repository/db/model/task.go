package model

import (
	"memorandum/repository/cache"
	"strconv"
	"time"
)

type Task struct {
	ID        uint `gorm:"primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	User      User       `gorm:"foreignKey:Uid;references:ID"`
	Uid       uint       `gorm:"not null"`
	Title     string     `json:"title" gorm:"index;not null"`
	Status    uint8      `gorm:"default:0"`
	Content   string     `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

type CreateTaskReq struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

type DeleteTaskReq struct {
	Id uint `json:"id" form:"id"`
}

type UpdateTaskReq struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  uint8  `form:"status" json:"status"` // 0 待办   1已完成
}

type ShowTaskReq struct {
	Id uint `json:"id" form:"id"`
}

type ListTasksReq struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

type SearchTaskReq struct {
	Info string `form:"info" json:"info"`
}

type TaskResp struct {
	ID        uint   `json:"id" example:"1"`       // 任务ID
	Title     string `json:"title" example:"吃饭"`   // 题目
	Content   string `json:"content" example:"睡觉"` // 内容
	Status    uint8  `json:"status" example:"0"`   // 状态(0未完成，1已完成)
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func (t *Task) View() uint64 {
	// 增加点击数
	countStr, _ := cache.RDB.Get(cache.TaskViewKey(t.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (t *Task) AddView() {
	cache.RDB.Incr(cache.TaskViewKey(t.ID))
	cache.RDB.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(t.ID)))
}
