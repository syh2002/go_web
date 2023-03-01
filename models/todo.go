package models

import (
	"gin/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
// 数据库表建立
func CreateATOdo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}
// 查
func GetTodoList() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil,  err
	}
	return
}
// 查一个
func GetATodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id=?",id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}
// 更新
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}
// 删除
func DeleteATode(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}