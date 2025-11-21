package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	bankGateURL = "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate"
	initialCode = "TANGcy"
	timeout     = 10 * time.Second
)

func main() {

	// 根据ExtraInfo解析的结果
	secretKey := "MuxiStudio203304"
	errorCode := "for {go func(){time.Sleep(1*time.Hour)}()}"
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiVEFOR2N5IiwiaWF0IjoxNzYzMTgxODczLCJuYmYiOjE3NjMxODE4NzN9.ZOcl_DidblW-EAWmRHGZS9h6FN9puPpzu4LO38pRF80"

	fmt.Printf("密钥: %s\n", secretKey)
	fmt.Printf("错误代码: %s\n", errorCode)

	// 1. 使用secret_key加密error_code
	encryptedCode, err := aesEncrypt(errorCode, secretKey)
	if err != nil {
		fmt.Printf("加密失败: %v\n", err)
		return
	}
	encryptedHex := fmt.Sprintf("%x", encryptedCode)

	// 2. 将加密过的error_code放入请求body
	result, err := sendEncryptedErrorCode(encryptedHex, passport)
	if err != nil {
		fmt.Printf("发送失败: %v\n", err)
		return
	}

	fmt.Printf("\n银行响应结果\n%s\n", result)
}

func aesEncrypt(plainText, key string) ([]byte, error) {
	keyBytes := []byte(key)
	if len(keyBytes) > 16 {
		keyBytes = keyBytes[:16]
	} else if len(keyBytes) < 16 {
		padding := make([]byte, 16-len(keyBytes))
		keyBytes = append(keyBytes, padding...)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	plainBytes := []byte(plainText)
	plainBytes = pkcs7Pad(plainBytes, block.BlockSize())

	cipherText := make([]byte, len(plainBytes))
	mode := cipher.NewCBCEncrypter(block, keyBytes[:block.BlockSize()])
	mode.CryptBlocks(cipherText, plainBytes)

	return cipherText, nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func sendEncryptedErrorCode(encryptedCode, passport string) (string, error) {
	client := &http.Client{Timeout: timeout}

	// 创建请求体 - 只包含加密后的error_code
	requestBody := map[string]string{
		"encrypted_code": encryptedCode,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("创建请求体失败: %v", err)
	}

	// 创建PUT请求
	req, err := http.NewRequest("PUT", bankGateURL, strings.NewReader(string(jsonBody)))
	if err != nil {
		return "", fmt.Errorf("创建PUT请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("code", initialCode)
	req.Header.Set("passport", passport)

	fmt.Println("请求头:")
	for k, v := range req.Header {
		fmt.Printf("  %s: %v\n", k, v)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("PUT请求失败: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("\n响应状态: %d %s\n", resp.StatusCode, resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	responseStr := string(body)
	fmt.Printf("响应内容: %s\n", responseStr)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	return responseStr, nil
}
