package helper

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

// GoogleAuthenticator Google Authenticator
type GoogleAuthenticator struct{}

// hmacSha1 hmac with sha1
func (ga *GoogleAuthenticator) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

// GenerateSecret generate new secret
func (ga *GoogleAuthenticator) GenerateSecret() (string, error) {
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.BigEndian, time.Now().UnixNano()/1000/30)
	secret := strings.ToUpper(base32.StdEncoding.EncodeToString(ga.hmacSha1(buf.Bytes(), nil)))
	return secret, nil
}

// GetCode get code by secret
func (ga *GoogleAuthenticator) GetCode(secret string) (code string, err error) {
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {
		return "", err
	}
	var b []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		b = append(b, byte((time.Now().Unix()/30>>shift)&mask))
	}
	hash := ga.hmacSha1(key, b)
	offset := hash[len(hash)-1] & 0x0F
	bts := hash[offset : offset+4]
	bts[0] = bts[0] & 0x7F
	number := ((uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])) % 1000000
	return fmt.Sprintf("%d", number), nil
}

// GetQRCode get QR code content
func (ga *GoogleAuthenticator) GetQRCode(user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
}

// NewGoogleAuthenticator new Google Authenticator
func NewGoogleAuthenticator() *GoogleAuthenticator {
	return &GoogleAuthenticator{}
}
