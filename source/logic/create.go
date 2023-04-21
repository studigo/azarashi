package logic

import (
	"fmt"

	"github.com/studigo/azarashi/task"
)

// タスクを生成する.
func Create(request *task.Request) (*task.Inner, error) {

	inner, err := request.ToInner()
	if err != nil {
		return nil, fmt.Errorf("create failed: %s", err)
	}

	dao := inner.ToDAO()
	if db.Create(&dao).Error != nil {
		return nil, fmt.Errorf("create failed: %s", err)
	}

	return inner, nil
}

// 子タスクを追加する.
func AddChild(parent *task.Inner, child *task.Request) (*task.Inner, *task.Inner, error) {

	inner, err := child.ToInner()
	if err != nil {
		return nil, nil, fmt.Errorf("add child failed: %s", err)
	}

	parent.AddChild(inner)

	dao := inner.ToDAO()
	if db.Create(&dao).Error != nil {
		parent.DelChild(inner)
		return nil, nil, fmt.Errorf("add child failed: %s", err)
	}

	return parent, inner, nil
}
