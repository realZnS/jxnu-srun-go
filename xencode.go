package main

import (
	"math"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ordat(msg string, idx int) int {
	if len(msg) > idx {
		return int(msg[idx])
	}
	return 0
}

func sencode(msg string, key bool) []int {
	l := len(msg)
	pwd := []int{}
	for i := 0; i < l; i += 4 {
		pwd = append(pwd, (ordat(msg, i) | ordat(msg, i+1)<<8 | ordat(msg, i+2)<<16 | ordat(msg, i+3)<<24))
	}
	if key {
		pwd = append(pwd, l)
	}
	return pwd
}

func lencode(msg []int, key bool) string {
	l := len(msg)
	ll := (l - 1) << 2
	if key {
		m := msg[l-1]
		if m < ll-3 || m > ll {
			return ""
		}
		ll = m
	}
	tmp := make([]string, l)
	for i := 0; i < l; i++ {
		var t strings.Builder
		t.WriteRune(rune(msg[i] & 0xff))
		t.WriteRune(rune(msg[i] >> 8 & 0xff))
		t.WriteRune(rune(msg[i] >> 16 & 0xff))
		t.WriteRune(rune(msg[i] >> 24 & 0xff))
		tmp[i] = t.String()
		// tmp[i] = string(msg[i]&0xff) + string(msg[i]>>8&0xff) + string(msg[i]>>16&0xff) + string(msg[i]>>24&0xff)
		if i > 0 {
			tmp[i] = tmp[i-1] + tmp[i]
		}
	}
	if key {
		return tmp[min(ll, l-1)]
	}
	return tmp[l-1]
}

func get_xencode(msg, key string) string {
	if msg == "" {
		return ""
	}
	pwd := sencode(msg, true)
	pwdk := sencode(key, false)
	for len(pwdk) < 4 {
		pwdk = append(pwdk, 0)
	}
	n := len(pwd) - 1
	z := pwd[n]
	y := pwd[0]
	c := 0x86014019 | 0x183639A0
	m := 0
	e := 0
	p := 0
	q := int(math.Floor(6 + 52/float64(n+1)))
	d := 0
	for 0 < q {
		d = d + c&-1 // (0x8CE0D9BF|0x731F2640)
		d = d & 0xffffffff
		e = d >> 2 & 3
		for p = 0; p < n; p++ {
			y = pwd[p+1]
			m = z>>5 ^ y<<2
			m = m + ((y>>3 ^ z<<4) ^ (d ^ y))
			m = m + (pwdk[(p&3)^e] ^ z)
			pwd[p] = pwd[p] + m&-1 // (0xEFB8D130|0x10472ECF)
			pwd[p] = pwd[p] & 0xffffffff
			z = pwd[p]
		}
		y = pwd[0]
		m = z>>5 ^ y<<2
		m = m + ((y>>3 ^ z<<4) ^ (d ^ y))
		m = m + (pwdk[(p&3)^e] ^ z)
		pwd[n] = pwd[n] + m&-1 // (0xBB390742|0x44C6F8BD)
		pwd[n] = pwd[n] & 0xffffffff
		z = pwd[n]
		q = q - 1
	}
	return lencode(pwd, false)
}
