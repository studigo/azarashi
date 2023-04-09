package task

import (
	"github.com/studigo/marinesnow/v2"
)

// タスクをDBに登録する際に使用する構造体.
type DAO struct {
	ID          int64
	Title       string
	Description string
	IsClosed    bool
	ParentID    int64
	RootID      int64
}

// テーブル名を指定.
func (*DAO) TableName() string {
	return "tasks"
}

// DAOからInnerを復元する(親子関係は復元しない).
func (p *DAO) toInner() *Inner {
	res := new(Inner)
	res.title.minLength = TITLE_MIN_LENGTH
	res.title.maxLength = TITLE_MAX_LENGTH
	res.description.minLength = DESCRIPTION_MIN_LENGTH
	res.description.maxLength = DESCRIPTION_MAX_LENGTH
	res.id = marinesnow.Snowflake(p.ID)
	res.title.Set(p.Title)
	res.description.Set(p.Description)
	res.isClosed = p.IsClosed
	res.parent = res
	res.children = []*Inner{}
	return res
}

// DAOからInnerを復元する.
func DAOToInner(dao []*DAO) []*Inner {
	res := []*Inner{}
	tmp := map[int64][]*Inner{}

	// 親子関係以外を復元.
	for _, v := range dao {
		inner := v.toInner()
		if v.ParentID == v.ID {
			res = append(res, inner)
		} else {
			tmp[v.ParentID] = append(tmp[v.ParentID], inner)
		}
	}

	// root以外の親子関係を復元.
	for _, v := range tmp {
		for _, inner := range v {
			if _, ok := tmp[inner.ID()]; ok {
				inner.children = tmp[inner.ID()]
			}
		}
	}

	// rootの親子関係を復元.
	for _, v := range res {
		if _, ok := tmp[v.ID()]; ok {
			v.children = tmp[v.ID()]
		}
	}

	return res
}
