package model

// 执行数据迁移(跑一下数据库迁移，如果不存在则新建表，存在判断字段是否一致不一致修改。)
func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}, &Task{})
	if err != nil {
		//panic(err)
		return
	}

	//uid 关联 User 的 id
	//DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")
}

//把task 表/user 表迁移到 数据库里面
