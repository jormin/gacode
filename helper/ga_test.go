package helper

import (
	"reflect"
	"testing"
)

func TestGoogleAuthenticator_GenerateSecret(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "01",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ga := &GoogleAuthenticator{}
				_, err := ga.GenerateSecret()
				if (err != nil) != tt.wantErr {
					t.Errorf("GenerateSecret() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func TestGoogleAuthenticator_GetCode(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "01",
			args: args{
				secret: "HIUKS7E5ZDQXM2HLQH5USZ7HZUQASSDQ",
			},
			wantErr: false,
		},
		{
			name: "02",
			args: args{
				secret: "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ga := &GoogleAuthenticator{}
				_, err := ga.GetCode(tt.args.secret)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetCode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func TestGoogleAuthenticator_GetQRCode(t *testing.T) {
	type args struct {
		user   string
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				user:   "abc",
				secret: "123",
			},
			want: "otpauth://totp/abc?secret=123",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ga := &GoogleAuthenticator{}
				if got := ga.GetQRCode(tt.args.user, tt.args.secret); got != tt.want {
					t.Errorf("GetQRCode() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestGoogleAuthenticator_hmacSha1(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "01",
			args: args{
				key:  []byte("123"),
				data: []byte("abc"),
			},
			want: []byte{84, 11, 12, 83, 212, 146, 88, 55, 189, 146, 179, 247, 26, 190, 122, 157, 112, 182, 118, 196},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ga := &GoogleAuthenticator{}
				if got := ga.hmacSha1(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("hmacSha1() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNewGoogleAuthenticator(t *testing.T) {
	tests := []struct {
		name string
		want *GoogleAuthenticator
	}{
		{
			name: "01",
			want: &GoogleAuthenticator{},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewGoogleAuthenticator(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewGoogleAuthenticator() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func BenchmarkGoogleAuthenticator_GenerateSecret(b *testing.B) {
	tt := struct {
		name    string
		wantErr bool
	}{
		name:    "01",
		wantErr: false,
	}
	for i := 0; i < b.N; i++ {
		ga := &GoogleAuthenticator{}
		_, err := ga.GenerateSecret()
		if (err != nil) != tt.wantErr {
			b.Errorf("GenerateSecret() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

func BenchmarkGoogleAuthenticator_GetCode(b *testing.B) {
	type args struct {
		secret string
	}
	tt := struct {
		name     string
		args     args
		wantCode string
		wantErr  bool
	}{
		name:
		"01",
		args: args{
			secret: "HIUKS7E5ZDQXM2HLQH5USZ7HZUQASSDQ",
		},
		wantErr: false,
	}
	for i := 0; i < b.N; i++ {
		ga := &GoogleAuthenticator{}
		_, err := ga.GetCode(tt.args.secret)
		if (err != nil) != tt.wantErr {
			b.Errorf("GetCode() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

func BenchmarkGoogleAuthenticator_hmacSha1(b *testing.B) {
	type args struct {
		key  []byte
		data []byte
	}
	tt := struct {
		name string
		args args
		want []byte
	}{
		name: "01",
		args: args{
			key:  []byte("123"),
			data: []byte("abc"),
		},
		want: []byte{84, 11, 12, 83, 212, 146, 88, 55, 189, 146, 179, 247, 26, 190, 122, 157, 112, 182, 118, 196},
	}
	for i := 0; i < b.N; i++ {
		ga := &GoogleAuthenticator{}
		if got := ga.hmacSha1(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
			b.Errorf("hmacSha1() = %v, want %v", got, tt.want)
		}
	}
}
