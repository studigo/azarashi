package task

import (
	"sort"
	"testing"

	"github.com/studigo/marinesnow/v2"
)

// 検証用のInnerを生成する.
func genTestInner(id int64, title string, description string, isClosed bool, children []*Inner) *Inner {
	inner := &Inner{
		id: marinesnow.Snowflake(id),
		title: varidateableString{
			value:     title,
			maxLength: TITLE_MAX_LENGTH,
			minLength: TITLE_MIN_LENGTH,
		},
		description: varidateableString{
			value:     description,
			maxLength: DESCRIPTION_MAX_LENGTH,
			minLength: DESCRIPTION_MIN_LENGTH,
		},
		isClosed: isClosed,
	}

	inner.parent = inner
	inner.children = children
	for _, v := range inner.children {
		v.parent = inner
	}

	return inner
}

// Inner同士を比較する.
func innerEqual(a *Inner, b *Inner) bool {
	var s func(i *Inner)
	s = func(i *Inner) {
		for _, v := range i.children {
			s(v)
		}
		sort.Slice(i.children,
			func(l, r int) bool {
				return i.children[l].ID() < i.children[r].ID()
			},
		)
	}

	var r func(a, b *Inner) bool
	r = func(a, b *Inner) bool {

		if len(a.children) != len(b.children) {
			return false
		}

		if a.id != b.id {
			return false
		}

		for i := 0; i < len(a.children); i++ {
			if !r(a.children[i], b.children[i]) {
				return false
			}
		}
		return true
	}

	s(a)
	s(b)
	return r(a, b)
}

func responseEqual(a, b *Response) bool {
	var s func(i *Response)
	s = func(i *Response) {
		for _, v := range i.Children {
			s(v)
		}
		sort.Slice(i.Children,
			func(l, r int) bool {
				return i.Children[l].ID < i.Children[r].ID
			},
		)
	}

	var r func(a, b *Response) bool
	r = func(a, b *Response) bool {

		if len(a.Children) != len(b.Children) {
			return false
		}

		if a.ID != b.ID {
			return false
		}

		for i := 0; i < len(a.Children); i++ {
			if !r(a.Children[i], b.Children[i]) {
				return false
			}
		}
		return true
	}

	s(a)
	s(b)
	return r(a, b)
}

func TestInner_ID(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want int64
	}{
		{
			// ただのGetterなので1ケース通過でOKとする.
			name: "正しいIDを取得できる.",
			p:    genTestInner(47120414214, "t", "d", false, []*Inner{}),
			want: 47120414214,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ID(); got != tt.want {
				t.Errorf("Inner.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_Title(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want string
	}{
		{
			// ただのGetterなので1ケース通過でOKとする.
			name: "正しいタイトルを取得できる.",
			p:    genTestInner(47120414214, "t", "d", false, []*Inner{}),
			want: "t",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Title(); got != tt.want {
				t.Errorf("Inner.Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_Description(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want string
	}{
		{
			// ただのGetterなので1ケース通過でOKとする.
			name: "正しい説明文を取得できる.",
			p:    genTestInner(47120414214, "t", "d", false, []*Inner{}),
			want: "d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Description(); got != tt.want {
				t.Errorf("Inner.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_IsClosed(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want bool
	}{
		{
			// ただのGetterなので1ケース通過でOKとする.
			name: "正しい完了状態を取得できる.",
			p:    genTestInner(47120414214, "t", "d", true, []*Inner{}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsClosed(); got != tt.want {
				t.Errorf("Inner.IsClosed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_Parent(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want *Inner
	}{
		{
			// ただのGetterなので1ケース通過でOKとする.
			// 重要なのは内容だけなのでpointerの比較は行わない.
			name: "正しい親オブジェクトを取得できる.",
			p:    genTestInner(47120414214, "t", "d", false, []*Inner{}),
			want: genTestInner(47120414214, "t", "d", false, []*Inner{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Parent(); !innerEqual(got, tt.want) {
				t.Errorf("Inner.Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_Children(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want []*Inner
	}{
		{
			// ただのGetterなので1ケース通過でOKとする.
			// 重要なのは内容だけなのでpointerの比較は行わない.
			name: "正しい子オブジェクト一覧を取得できる.",
			p: genTestInner(47120414214, "t", "d", false, []*Inner{
				genTestInner(2, "c1t", "c1d", false, []*Inner{}),
				genTestInner(3, "c2t", "c2d", true, []*Inner{}),
			}),
			want: []*Inner{
				genTestInner(2, "c1t", "c1d", false, []*Inner{}),
				genTestInner(3, "c2t", "c2d", true, []*Inner{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Children()
			ok := true
			for i := 0; i < 2; i++ {
				ok = ok && innerEqual(got[i], tt.want[i])
			}
			if !ok {
				t.Errorf("Inner.Children() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_IsRoot(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want bool
	}{
		{
			name: "自身がrootである場合.",
			p: genTestInner(47120414214, "t", "d", false, []*Inner{
				genTestInner(2, "c1t", "c1d", false, []*Inner{}),
				genTestInner(3, "c2t", "c2d", true, []*Inner{}),
			}),
			want: true,
		},
		{
			name: "自身がrootでない場合.",
			p: genTestInner(47120414214, "t", "d", false, []*Inner{
				genTestInner(2, "c1t", "c1d", false, []*Inner{}),
				genTestInner(3, "c2t", "c2d", true, []*Inner{}),
			}).Children()[0],
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsRoot(); got != tt.want {
				t.Errorf("Inner.IsRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_ChangeTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		p       *Inner
		args    args
		wantErr bool
	}{
		{
			name:    "長さ0の場合.",
			p:       genTestInner(1, "t", "d", false, []*Inner{}),
			args:    args{title: ""},
			wantErr: true,
		},
		{
			name:    "長さ1の場合.",
			p:       genTestInner(1, "t", "d", false, []*Inner{}),
			args:    args{title: "a"},
			wantErr: false,
		},
		{
			name: "長さ100の場合.",
			p:    genTestInner(1, "t", "d", false, []*Inner{}),
			args: func() args {
				ret := args{title: ""}
				for i := 0; i < 100; i++ {
					ret.title += "a"
				}
				return ret
			}(),
			wantErr: false,
		},
		{
			name: "長さ101の場合.",
			p:    genTestInner(1, "t", "d", false, []*Inner{}),
			args: func() args {
				ret := args{title: ""}
				for i := 0; i < 101; i++ {
					ret.title += "a"
				}
				return ret
			}(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.ChangeTitle(tt.args.title); (err != nil) != tt.wantErr {
				t.Errorf("Inner.ChangeTitle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInner_ChangeDescription(t *testing.T) {
	type args struct {
		description string
	}
	tests := []struct {
		name    string
		p       *Inner
		args    args
		wantErr bool
	}{
		{
			name:    "長さ0の場合.",
			p:       genTestInner(1, "t", "d", false, []*Inner{}),
			args:    args{description: ""},
			wantErr: false,
		},
		{
			name: "長さ1000の場合.",
			p:    genTestInner(1, "t", "d", false, []*Inner{}),
			args: func() args {
				ret := args{description: ""}
				for i := 0; i < 1000; i++ {
					ret.description += "a"
				}
				return ret
			}(),
			wantErr: false,
		},
		{
			name: "長さ1001の場合.",
			p:    genTestInner(1, "t", "d", false, []*Inner{}),
			args: func() args {
				ret := args{description: ""}
				for i := 0; i < 1001; i++ {
					ret.description += "a"
				}
				return ret
			}(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.ChangeDescription(tt.args.description); (err != nil) != tt.wantErr {
				t.Errorf("Inner.ChangeDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInner_Has(t *testing.T) {
	type args struct {
		child *Inner
	}
	tests := []struct {
		name string
		p    *Inner
		args args
		want bool
	}{
		{
			name: "指定した子が無い場合.",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c", "c", false, []*Inner{}),
			}),
			args: args{child: genTestInner(3, "a", "a", false, []*Inner{})},
			want: false,
		},
		{
			name: "指定した子がある場合.",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c", "c", false, []*Inner{}),
			}),
			args: args{child: genTestInner(2, "c", "c", false, []*Inner{})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Has(tt.args.child); got != tt.want {
				t.Errorf("Inner.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_AddChild(t *testing.T) {
	type args struct {
		child *Inner
	}
	tests := []struct {
		name string
		p    *Inner
		args args
		want *Inner
	}{
		{
			name: "正しく子を追加できる.",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c", "c", false, []*Inner{}),
			}),
			args: args{genTestInner(3, "d", "d", false, []*Inner{})},
			want: genTestInner(3, "d", "d", false, []*Inner{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.AddChild(tt.args.child)
			got := tt.p.Children()[1]
			if !innerEqual(got, tt.want) {
				t.Errorf("Inner.AddChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_DelChild(t *testing.T) {
	type args struct {
		child *Inner
	}
	tests := []struct {
		name string
		p    *Inner
		args args
		want *Inner
	}{
		{
			name: "正しく子を削除できる(他の子が消えていない).",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c", "c", false, []*Inner{}),
				genTestInner(3, "d", "d", false, []*Inner{}),
			}),
			args: args{child: genTestInner(2, "c", "c", false, []*Inner{})},
			want: genTestInner(3, "d", "d", false, []*Inner{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.DelChild(tt.args.child)
			got := tt.p.children[0]
			if !innerEqual(got, tt.want) && len(tt.p.children) == 1 {
				t.Errorf("Inner.DelChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_ChangeParent(t *testing.T) {
	type args struct {
		newParent *Inner
	}
	tests := []struct {
		name string
		p    *Inner
		args args
		want *Inner
	}{
		{
			name: "正しく親を変更でき, 元の親に親子情報が残らない.",
			p: genTestInner(1, "p1", "p1", false, []*Inner{
				genTestInner(2, "c", "c", false, []*Inner{}),
			}),
			args: args{newParent: genTestInner(3, "p2", "p2", false, []*Inner{})},
			want: genTestInner(2, "c", "c", false, []*Inner{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.children[0].ChangeParent(tt.args.newParent)
			l1 := len(tt.p.children)
			l2 := len(tt.args.newParent.children)
			got := tt.args.newParent.children[0]
			if l1 != 0 || l2 != 1 || !innerEqual(got, tt.want) {
				t.Errorf("Inner.DelChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_Root(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want *Inner
	}{
		{
			name: "ひ孫要素からrootを取得できる.",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c1", "d", false, []*Inner{
					genTestInner(3, "g1", "d", false, []*Inner{
						genTestInner(4, "gg1", "d", false, []*Inner{}),
					}),
				}),
			}),
			want: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c1", "d", false, []*Inner{
					genTestInner(3, "g1", "d", false, []*Inner{
						genTestInner(4, "gg1", "d", false, []*Inner{}),
					}),
				}),
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.children[0].children[0].children[0].Root()
			if !innerEqual(got, tt.want) {
				t.Errorf("Inner.Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_Locked(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want bool
	}{
		{
			name: "ロックされていない場合.",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c1", "c1", true, []*Inner{}),
				genTestInner(3, "c2", "c2", true, []*Inner{}),
			}),
			want: false,
		},
		{
			name: "ロックされている場合.",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c1", "c1", true, []*Inner{}),
				genTestInner(3, "c2", "c2", false, []*Inner{}),
			}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Locked(); got != tt.want {
				t.Errorf("Inner.Locked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_Close(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
	}{
		{
			name: "未完了の場合.",
			p:    genTestInner(1, "t", "d", false, []*Inner{}),
		},
		{
			name: "完了済みの場合.",
			p:    genTestInner(1, "t", "d", true, []*Inner{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Close()
			if tt.p.isClosed == false {
				t.Errorf("Inner.Close()")
			}
		})
	}
}

func TestInner_Open(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
	}{
		{
			name: "未完了の場合.",
			p:    genTestInner(1, "t", "d", false, []*Inner{}),
		},
		{
			name: "完了済みの場合.",
			p:    genTestInner(1, "t", "d", true, []*Inner{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Open()
			if tt.p.isClosed == true {
				t.Errorf("Inner.Open()")
			}
		})
	}
}

func TestInner_ToResponse(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want *Response
	}{
		{
			name: "正しくResponseに変換できる.",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c1", "d", true, []*Inner{
					genTestInner(3, "g1", "d", true, []*Inner{
						genTestInner(4, "gg1", "d", true, []*Inner{}),
					}),
					genTestInner(5, "g2", "d", false, []*Inner{}),
					genTestInner(6, "g3", "d", true, []*Inner{}),
				}),
				genTestInner(7, "c2", "d", false, []*Inner{}),
			}),
			want: &Response{ID: 1, Title: "t", Description: "d", IsClosed: false, Children: []*Response{
				{ID: 2, Title: "c1", Description: "d", IsClosed: true, Children: []*Response{
					{ID: 3, Title: "g1", Description: "d", IsClosed: true, Children: []*Response{
						{ID: 4, Title: "gg1", Description: "d", IsClosed: true, Children: []*Response{}},
					}},
					{ID: 5, Title: "g2", Description: "d", IsClosed: false, Children: []*Response{}},
					{ID: 6, Title: "g3", Description: "d", IsClosed: true, Children: []*Response{}},
				}},
				{ID: 7, Title: "c2", Description: "d", IsClosed: false, Children: []*Response{}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ToResponse(); !responseEqual(got, tt.want) {
				t.Errorf("Inner.ToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInner_ToDAO(t *testing.T) {
	tests := []struct {
		name string
		p    *Inner
		want []*DAO
	}{
		{
			name: "正しくDAOの配列に変換できる.",
			p: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "c1", "d", true, []*Inner{
					genTestInner(3, "g1", "d", true, []*Inner{
						genTestInner(4, "gg1", "d", true, []*Inner{}),
					}),
					genTestInner(5, "g2", "d", false, []*Inner{}),
					genTestInner(6, "g3", "d", true, []*Inner{}),
				}),
				genTestInner(7, "c2", "d", false, []*Inner{}),
			}),
			want: []*DAO{
				{ID: 1, Title: "t", Description: "d", IsClosed: false, ParentID: 1, RootID: 1},
				{ID: 2, Title: "c1", Description: "d", IsClosed: true, ParentID: 1, RootID: 1},
				{ID: 3, Title: "g1", Description: "d", IsClosed: true, ParentID: 2, RootID: 1},
				{ID: 4, Title: "gg1", Description: "d", IsClosed: true, ParentID: 3, RootID: 1},
				{ID: 5, Title: "g2", Description: "d", IsClosed: false, ParentID: 2, RootID: 1},
				{ID: 6, Title: "g3", Description: "d", IsClosed: true, ParentID: 2, RootID: 1},
				{ID: 7, Title: "c2", Description: "d", IsClosed: false, ParentID: 1, RootID: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.ToDAO()
			sort.Slice(got, func(a, b int) bool { return got[a].ID < got[b].ID })
			sort.Slice(tt.want, func(a, b int) bool { return tt.want[a].ID < tt.want[b].ID })
			equal := func(a, b []*DAO) bool {
				ok := true
				for i := 0; i < 7; i++ {
					id := got[i].ID == tt.want[i].ID
					title := got[i].Title == tt.want[i].Title
					description := got[i].Description == tt.want[i].Description
					isClosed := got[i].IsClosed == tt.want[i].IsClosed
					parentID := got[i].ParentID == tt.want[i].ParentID
					rootID := got[i].RootID == tt.want[i].RootID
					ok = ok && id && title && description && isClosed && parentID && rootID
				}
				return ok
			}
			if !equal(got, tt.want) {
				t.Errorf("Inner.ToDAO() = %v, want %v", got, tt.want)
			}
		})
	}
}
