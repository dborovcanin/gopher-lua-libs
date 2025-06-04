// Package crypto implements golang package crypto functionality for lua.
package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// MD5 lua crypto.md5(string) return string
func MD5(L *lua.LState) int {
	str := L.CheckString(1)
	hash := md5.Sum([]byte(str))
	L.Push(lua.LString(fmt.Sprintf("%x", hash)))
	return 1
}

// SHA256 lua crypto.sha256(string) return string
func SHA256(L *lua.LState) int {
	str := L.CheckString(1)
	hash := sha256.Sum256([]byte(str))
	L.Push(lua.LString(fmt.Sprintf("%x", hash)))
	return 1
}

func AESEncrypt(l *lua.LState) int {
	key, iv, data, err := decodeParams(l)
	if err != nil {
		l.Push(lua.LString(fmt.Sprintf("failed to decode params: %v", err)))
		return 2
	}

	enc, err := encrypt(key, iv, data)
	if err != nil {
		l.Push(lua.LString(fmt.Sprintf("failed to encrypt: %v", err)))
		return 2
	}
	l.Push(lua.LString(hex.EncodeToString(enc)))

	return 1
}

func AESDecrypt(l *lua.LState) int {
	key, iv, data, err := decodeParams(l)
	if err != nil {
		l.Push(lua.LString(fmt.Sprintf("failed to decode params: %v", err)))
		return 2
	}

	dec, err := decrypt(key, iv, data)
	if err != nil {
		l.Push(lua.LString(fmt.Sprintf("failed to decrypt: %v", err)))
		return 2
	}

	l.Push(lua.LString(hex.EncodeToString(dec)))

	return 1
}

func decodeParams(l *lua.LState) (key, iv, data []byte, err error) {
	keyStr := l.ToString(1)
	ivStr := l.ToString(2)
	dataStr := l.ToString(3)

	key, err = hex.DecodeString(keyStr)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode key: %v", err)
	}

	iv, err = hex.DecodeString(ivStr)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode IV: %v", err)
	}

	data, err = hex.DecodeString(dataStr)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode data: %v", err)
	}

	return key, iv, data, nil
}
