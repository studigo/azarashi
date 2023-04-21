package task

import (
	"fmt"
)

// タスク生成時に使用する構造体.
type Request struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Children    []*Request `json:"children"`
}

// Innerに変換する.
func (p *Request) ToInner() (*Inner, error) {

	res, err := NewInner(p.Title, p.Description)
	if err != nil {
		return nil, fmt.Errorf("convert failed: %s", err)
	}

	for _, v := range p.Children {
		c, err := v.ToInner()
		if err != nil {
			return nil, fmt.Errorf("convert failed: %s", err)
		}
		res.children = append(res.children, c)
		c.parent = res
	}

	return res, nil
}
