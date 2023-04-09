package task

import "testing"

func Test_varidateableString_MinLength(t *testing.T) {
	tests := []struct {
		name string
		p    *varidateableString
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.MinLength(); got != tt.want {
				t.Errorf("varidateableString.MinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_varidateableString_MaxLength(t *testing.T) {
	tests := []struct {
		name string
		p    *varidateableString
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.MaxLength(); got != tt.want {
				t.Errorf("varidateableString.MaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_varidateableString_Value(t *testing.T) {
	tests := []struct {
		name string
		p    *varidateableString
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Value(); got != tt.want {
				t.Errorf("varidateableString.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Set(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("varidateableString.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
