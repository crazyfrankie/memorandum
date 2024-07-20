package util

import (
	"memorandum/repository/cache"
	"memorandum/repository/db/model"
	"strconv"
)

type task model.Task

func (t *task) View() uint64 {
	// 增加点击数
	countStr, _ := cache.RDB.Get(cache.TaskViewKey(t.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (t *task) AddView() {
	cache.RDB.Incr(cache.TaskViewKey(t.ID))
	cache.RDB.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(t.ID)))
}
