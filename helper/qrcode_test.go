package helper

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/xid"
	"github.com/skip2/go-qrcode"
)

func TestPrintQRCode(t *testing.T) {
	type args struct {
		file string
	}

	file := fmt.Sprintf("%s.png", xid.New().String())
	_ = qrcode.WriteFile("test", qrcode.Medium, 50, file)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "01",
			args:    args{file: file},
			wantErr: false,
		},
		{
			name:    "02",
			args:    args{file: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if err := PrintQRCode(tt.args.file); (err != nil) != tt.wantErr {
					t.Errorf("PrintQRCode() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
	_ = os.Remove(file)
}

func BenchmarkPrintQRCode(t *testing.B) {
	type args struct {
		file string
	}
	file := fmt.Sprintf("%s.png", xid.New().String())
	_ = qrcode.WriteFile("test", qrcode.Medium, 50, file)
	tt := struct {
		name    string
		args    args
		wantErr bool
	}{
		name:    "01",
		args:    args{file: file},
		wantErr: false,

	}
	for i := 0; i < t.N; i++ {
		if err := PrintQRCode(tt.args.file); (err != nil) != tt.wantErr {
			t.Errorf("PrintQRCode() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
	_ = os.Remove(file)
}
