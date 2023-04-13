package task

import (
	"testing"
)

func TestRequest_ToInner(t *testing.T) {
	tests := []struct {
		name    string
		p       *Request
		want    *Inner
		wantErr bool
	}{
		{
			name: "正常値.",
			p: &Request{Title: "t", Description: "d", Children: []*Request{
				{Title: "t", Description: "d", Children: []*Request{}},
			}},
			want: genTestInner(1, "t", "d", false, []*Inner{
				genTestInner(2, "ct", "cd", false, []*Inner{}),
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.ToInner()
			tt.want.id = got.id
			tt.want.children[0].id = got.children[0].id
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.ToInner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !innerEqual(got, tt.want) {
				t.Errorf("Request.ToInner() = %v, want %v", got, tt.want)
			}
		})
	}
}
