package task

import "testing"

func Test_varidateableString_Set(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		p       *varidateableString
		args    args
		wantErr bool
	}{
		{
			name: "正常値.",
			p: &varidateableString{
				minLength: 1,
				maxLength: 1,
				value:     "a",
			},
			args:    args{value: "b"},
			wantErr: false,
		},
		{
			name: "超過.",
			p: &varidateableString{
				minLength: 1,
				maxLength: 1,
				value:     "a",
			},
			args:    args{value: "bb"},
			wantErr: true,
		},
		{
			name: "不足.",
			p: &varidateableString{
				minLength: 1,
				maxLength: 1,
				value:     "a",
			},
			args:    args{value: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Set(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("varidateableString.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
