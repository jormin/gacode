package helper

import (
	"fmt"
	"os"
	"os/user"
	"reflect"
	"testing"

	"github.com/jormin/gacode/entity"
)

func TestGetDatafilePath(t *testing.T) {
	u, _ := user.Current()
	tests := []struct {
		name string
		want string
	}{
		{
			name: "01",
			want: fmt.Sprintf("%s/.gacode", u.HomeDir),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := GetDatafilePath(); got != tt.want {
					t.Errorf("GetDatafilePath() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNewData(t *testing.T) {
	tests := []struct {
		name string
		want *entity.Data
	}{
		{
			name: "01",
			want: &entity.Data{
				Accounts: []*entity.Account{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewData(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewData() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestReadData(t *testing.T) {
	u, _ := user.Current()
	file := fmt.Sprintf("%s/.gacode", u.HomeDir)
	_ = os.Remove(file)
	tests := []struct {
		name    string
		wantErr bool
		deal    func()
	}{
		{
			name:    "01",
			wantErr: false,
			deal:    nil,
		},
		{
			name:    "02",
			wantErr: false,
			deal: func() {
				_ = os.WriteFile(file, []byte(`{"accounts":[]}`), 0777)
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if tt.deal != nil {
					tt.deal()
				}
				_, err := ReadData()
				if (err != nil) != tt.wantErr {
					t.Errorf("ReadData() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
	_ = os.Remove(file)
}

func TestWriteData(t *testing.T) {
	type args struct {
		data *entity.Data
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "01",
			args:    args{data: NewData()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if err := WriteData(tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("WriteData() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
