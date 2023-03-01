package main

import (
	"gin/dao"
	"gin/models"
	"gin/rounters"
)

func main() {
	// 创建数据库
	// bubble
	//	连接数据库
	err := dao.InitMysql()

	if err != nil {
		panic(err)
	}

	defer dao.Close()	// 关闭数据库
	// 绑定模型
	dao.DB.AutoMigrate(&models.Todo{})
		
	r := rounters.SetupRounter()
	
	r.Run()
}
