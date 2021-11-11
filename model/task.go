package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	"to-do-list/cache"
)

//任务模型
type Task struct {
	gorm.Model
	User 		  User   `gorm:"ForeignKey:UID"`
	Title         string `gorm:"index;not null"`
	Status        int    `gorm:"default:'0'"`
	Content       string `gorm:"type:longtext"`
	StartTime 	  time.Time
	EndTime 	  time.Time `gorm:"default:'0'"`
}

func (Task *Task) View() uint64 {
	//增加点击数
 	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//AddView
func (Task *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(Task.ID))	//增加视频点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID)))	//增加排行点击数
}

