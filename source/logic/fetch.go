package logic

import (
	"fmt"

	"github.com/studigo/azarashi/task"
)

// タスクを取得する.
func Fetch(id int64) (*task.Inner, error) {

	dao := []*task.DAO{}
	if err := db.Where(&task.DAO{RootID: id}).Find(&dao).Error; err != nil {
		return nil, fmt.Errorf("fetch failed: %s", err)
	}

	if len(dao) == 0 {
		return nil, fmt.Errorf("fetch failed: resource not found")
	}

	res := task.DAOToInner(dao)[0]
	return res, nil
}
