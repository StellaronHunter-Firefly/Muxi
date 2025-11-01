package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	password := "muxi2025"
	paper := "uNTaNxYKd72BmdUi001LlzM64rw64IG1EohE4qCzR+CGk7iJdpnEtIByqHbieWHtduDyyUDN3arNalT9+RAv210fq+6XbdXz0D0st4L4IQZDoLpKVZ1D1W4tadwi1J/rJIdbsnvUq6IoSxHYXdCW3w=="
	fmt.Println(DecodeEncryptedBase64(paper, password))
}

func DecodeEncryptedBase64(paper string, password string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(strings.TrimSpace(paper))
	if err != nil {
		return "", fmt.Errorf("base64 解码失败: %w", err)
	}
	key := sha256.Sum256([]byte(password))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(raw) < nonceSize {
		return "", fmt.Errorf("数据长度不足: %d < %d", len(raw), nonceSize)
	}
	nonce := raw[:nonceSize]
	ciphertext := raw[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("GCM 解密失败: %w", err)
	}
	return string(plaintext), nil
}

//cd D:\GoProgram\exam25\t2
