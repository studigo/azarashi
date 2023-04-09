package task

import (
	"reflect"
	"testing"
)

func TestRequest_ToInner(t *testing.T) {
	tests := []struct {
		name    string
		p       *Request
		want    *Inner
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.ToInner()
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.ToInner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.ToInner() = %v, want %v", got, tt.want)
			}
		})
	}
}
