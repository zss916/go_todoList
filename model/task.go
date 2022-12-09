package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"to-do-list/cache"
)

// 任务模型
type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"` //外键关联
	Uid       uint   `gorm:"not null"`       //不能null
	Title     string `gorm:"index;not null"` //当作索引，查询快一点
	Status    int    `gorm:"default:0"`      //0是未完成，1已完成
	Content   string `gorm:"type:longtext"`  //longtext 长字段
	StartTime int64  //备忘录的开始时间
	EndTime   int64  `gorm:"default:0"` //备忘录的结束时间
}

func (Task *Task) View() uint64 {
	//增加点击数
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView
func (Task *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(Task.ID))                      //增加视频点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID))) //增加排行点击数
}
