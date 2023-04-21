package logic

import (
	"fmt"

	"github.com/studigo/azarashi/task"
)

// タスクを削除する(子もすべて再帰的に消える).
func Delete(target *task.Inner) error {

	var deltarget = []int64{}

	var dfs func(p *task.Inner)
	dfs = func(p *task.Inner) {
		deltarget = append(deltarget, p.ID())
		for _, v := range p.Children() {
			dfs(v)
		}
	}
	dfs(target)
	if err := db.Delete(&task.DAO{}, deltarget).Error; err != nil {
		return fmt.Errorf("delete failed: %s", err)
	}

	target.Parent().DelChild(target)
	return nil
}
