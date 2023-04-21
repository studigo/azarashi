package task

import (
	"fmt"

	"github.com/studigo/marinesnow/v2"
)

// タイトルの文字数制限.
const (
	TITLE_MIN_LENGTH = 1
	TITLE_MAX_LENGTH = 100
)

// 説明文の文字数制限.
const (
	DESCRIPTION_MIN_LENGTH = 0
	DESCRIPTION_MAX_LENGTH = 1000
)

// 内部処理で使用するタスクを表す構造体.
type Inner struct {
	id          marinesnow.Snowflake
	title       varidateableString
	description varidateableString
	isClosed    bool
	parent      *Inner
	children    []*Inner
}

func (p *Inner) ID() int64           { return int64(p.id) }
func (p *Inner) Title() string       { return p.title.value }
func (p *Inner) Description() string { return p.description.value }
func (p *Inner) IsClosed() bool      { return p.isClosed }
func (p *Inner) Parent() *Inner      { return p.parent }
func (p *Inner) Children() []*Inner  { return p.children }
func (p *Inner) IsRoot() bool        { return p.id == p.parent.id }

// タイトルを変更する.
func (p *Inner) ChangeTitle(title string) error {
	if err := p.title.Set(title); err != nil {
		return fmt.Errorf("title change failed: %s", err)
	}
	return nil
}

// 説明文を変更する.
func (p *Inner) ChangeDescription(description string) error {
	if err := p.description.Set(description); err != nil {
		return fmt.Errorf("description change failed: %s", err)
	}
	return nil
}

// 指定した子があるか判定する.
func (p *Inner) Has(child *Inner) bool {
	for _, v := range p.children {
		if v.id == child.id {
			return true
		}
	}
	return false
}

// 子を追加する(重複チェックあり).
func (p *Inner) AddChild(child *Inner) {
	if !p.Has(child) {
		child.parent = p
		p.children = append(p.children, child)
	}
}

// 子を削除する.
func (p *Inner) DelChild(child *Inner) {
	for i, v := range p.children {
		if v.id == child.id {
			child.parent = child
			p.children = append(p.children[:i], p.children[i+1:]...)
			return
		}
	}
}

// 親を変更する.
func (p *Inner) ChangeParent(newParent *Inner) {
	p.parent.DelChild(p)
	newParent.AddChild(p)
	p.parent = newParent
}

// rootを取得する.
func (p *Inner) Root() *Inner {
	if p.IsRoot() {
		return p
	}
	return p.parent.Root()
}

// 子タスクが未完了ならtrue.
func (p *Inner) Locked() bool {
	for _, v := range p.children {
		if !v.isClosed {
			return true
		}
	}
	return false
}

// タスクを閉じる.
func (p *Inner) Close() {
	p.isClosed = true
}

// タスクを開く.
func (p *Inner) Open() {
	p.isClosed = false
}

// Inner を生成する.
func NewInner(title string, description string) (*Inner, error) {
	res := new(Inner)
	res.title.minLength = TITLE_MIN_LENGTH
	res.title.maxLength = TITLE_MAX_LENGTH
	res.description.minLength = DESCRIPTION_MIN_LENGTH
	res.description.maxLength = DESCRIPTION_MAX_LENGTH

	if err := res.title.Set(title); err != nil {
		return nil, fmt.Errorf("generation failed: %s", err)
	}

	if err := res.description.Set(description); err != nil {
		return nil, fmt.Errorf("generation failed: %s", err)
	}

	sf, err := marinesnow.Generate()
	if err != nil {
		return nil, fmt.Errorf("generation failed: %s", err)
	}

	res.id = sf
	res.isClosed = false
	res.parent = res
	res.children = []*Inner{}

	return res, nil
}

// Responseに変換する.
func (p *Inner) ToResponse() *Response {
	res := new(Response)
	res.ID = p.ID()
	res.Title = p.Title()
	res.Description = p.Description()
	res.IsClosed = p.IsClosed()

	for _, v := range p.children {
		c := v.ToResponse()
		res.Children = append(res.Children, c)
	}

	return res
}

// DAOに変換する.
func (p *Inner) ToDAO() []*DAO {
	res := []*DAO{}

	var dfs func(i *Inner)
	dfs = func(i *Inner) {

		res = append(res, &DAO{
			ID:          i.ID(),
			Title:       i.Title(),
			Description: i.Description(),
			IsClosed:    i.IsClosed(),
			ParentID:    i.parent.ID(),
			RootID:      i.Root().ID(),
		})

		for _, v := range i.children {
			dfs(v)
		}
	}
	dfs(p)

	return res
}
