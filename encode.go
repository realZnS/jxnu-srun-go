package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"golang.org/x/text/encoding/charmap"
)

func get_base64(s string) string {
	latin1, _ := charmap.ISO8859_1.NewEncoder().Bytes([]byte(s))
	const ALPHA = "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA"
	s64 := base64.NewEncoding(ALPHA).EncodeToString([]byte(latin1))
	return string(s64)
}

func get_md5(password string, token string) string {
	h := hmac.New(md5.New, []byte(token))
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func get_sha1(value string) string {
	h := sha1.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%x", h.Sum(nil))
}
