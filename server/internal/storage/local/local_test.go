package local

import (
	"blog/internal/config"
	"path"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		setting *config.App
	}
	tests := []struct {
		name    string
		args    args
		want    *Local
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.setting)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	str := "uploads/2022/04/06/asdfadsfas.png.jpg"
	t.Log(path.Base(str))
}
