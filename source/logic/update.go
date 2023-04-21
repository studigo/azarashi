package logic

import (
	"fmt"

	"github.com/studigo/azarashi/task"
	"gorm.io/gorm"
)

// タスクを完了状態にする(子タスクが完了していなければエラー).
func Close(target *task.Inner) error {

	if target.Locked() {
		return fmt.Errorf("close failed: target is locked")
	}

	err := db.Model(&task.DAO{}).Where("id = ?", target.ID()).Update("is_closed", true).Error
	if err != nil {
		return fmt.Errorf("close failed: %s", err)
	}

	target.Close()
	return nil
}

// タスクを未完了状態にする(親タスクも再帰的に未完了にする).
func Open(target *task.Inner) error {

	openTarget := []int64{}

	var rec func(i *task.Inner, open bool) *task.Inner
	rec = func(i *task.Inner, open bool) *task.Inner {
		openTarget = append(openTarget, i.ID())
		if open {
			i.Open()
		}
		if i.IsRoot() {
			return nil
		}
		return rec(i.Parent(), open)
	}
	rec(target, false)
	fmt.Println(openTarget)
	err := db.Model(&task.DAO{}).Where("id IN ?", openTarget).Update("is_closed", false).Error
	if err != nil {
		return fmt.Errorf("open failed: %s", err)
	}

	rec(target, true)

	return nil
}

// タスクのタイトルを変更する.
func ChangeTitle(target *task.Inner, title string) error {

	old := target.Title()
	if err := target.ChangeTitle(title); err != nil {
		return err
	}

	err := db.Model(&task.DAO{}).Where("id = ?", target.ID()).Update("title", title).Error
	if err != nil {
		target.ChangeTitle(old)
		return fmt.Errorf("title change failed: %s", err)
	}
	return nil
}

// 説明文を変更する.
func ChangeDescription(target *task.Inner, description string) error {
	old := target.Description()
	if err := target.ChangeDescription(description); err != nil {
		return err
	}

	res := db.Model(&task.DAO{}).Where("id = ?", target.ID()).Update("description", description)
	if res.Error != nil {
		target.ChangeDescription(old)
		return fmt.Errorf("description change failed: %s", res.Error)
	}
	return nil
}

// 親を変更する.
func ChangeParent(target *task.Inner, newParent *task.Inner) error {

	old := target.Parent()
	target.ChangeParent(newParent)

	err := db.Transaction(func(tx *gorm.DB) error {
		dao := target.ToDAO()
		ids := []int64{}
		for _, v := range dao {
			ids = append(ids, v.ID)
		}
		if err := db.Model(&task.DAO{}).Where("id IN ?", ids).Updates(
			map[string]interface{}{
				"parent_id": newParent.ID(),
				"root_id":   newParent.Root().ID(),
			},
		).Error; err != nil {
			return fmt.Errorf("parent change failed: %s", err)
		}
		return nil
	})

	if err != nil {
		target.ChangeParent(old)
		return err
	}

	return nil
}
