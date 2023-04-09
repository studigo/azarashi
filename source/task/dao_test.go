package task

import (
	"testing"
)

// 検証用のDAOを生成する.
func genTestDAO(id int64, title string, description string, isClosed bool, pid int64, rid int64) *DAO {
	return &DAO{
		ID:          id,
		Title:       title,
		Description: description,
		IsClosed:    isClosed,
		ParentID:    pid,
		RootID:      rid,
	}
}

func TestDAO_toInner(t *testing.T) {
	tests := []struct {
		name string
		p    *DAO
		want *Inner
	}{
		{
			// DAOはDBの値かInnerから生成されるので正しい値であることが保証されている.
			// そのため正しくInnerに変換できている事だけチェックすれば良い.
			name: "DAOToInnerの内部で使用するサブ関数のチェック.",
			p:    genTestDAO(1, "pt", "pd", true, 1, 1),
			want: genTestInner(1, "pt", "pd", true, []*Inner{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.toInner(); !innerEqual(got, tt.want) {
				t.Errorf("DAO.toInner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDAOToInner(t *testing.T) {
	type args struct {
		dao []*DAO
	}
	tests := []struct {
		name string
		args args
		want []*Inner
	}{
		{
			name: "DAOの配列からInnerを生成できることを確認する.",
			args: args{
				dao: []*DAO{
					genTestDAO(1, "pt", "pd", false, 1, 1),
					genTestDAO(2, "c1t", "c1d", false, 1, 1),
					genTestDAO(3, "g1t", "g1d", true, 2, 1),
					genTestDAO(4, "g2t", "g2d", false, 2, 1),
					genTestDAO(5, "c2t", "c2d", true, 1, 1),
				},
			},

			want: []*Inner{
				genTestInner(1, "pt", "pd", false, []*Inner{
					genTestInner(2, "c1t", "c1d", false, []*Inner{
						genTestInner(3, "g1t", "g1d", true, []*Inner{}),
						genTestInner(4, "g2t", "g2d", false, []*Inner{}),
					}),
					genTestInner(5, "c2t", "c2d", true, []*Inner{}),
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DAOToInner(tt.args.dao)[0]
			want := tt.want[0]
			if !innerEqual(got, want) {
				t.Errorf("DAOToInner() = %v, want %v", got, want)
			}
		})
	}
}
