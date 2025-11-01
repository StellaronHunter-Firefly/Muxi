package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	const loginUrl = "https://demo.muxixyz.com/library/login"

	client := &http.Client{}

	type UserMessage struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = username[:len(username)-2]

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = password[:len(password)-2]
	usermessage := UserMessage{
		Username: username,
		Password: password,
	}

	jsonData, err := json.Marshal(usermessage)
	if err != nil {
		fmt.Printf("序列化JSON失败:%v", err)
		return
	}
	request, err := http.NewRequest("POST", loginUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("构建请求时失败:%v", err)
		return
	}

	request.Header.Set("content-type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("发送请求时失败:%v", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("读取响应出错:%v", err)
		return
	}

	type Data struct {
		Cookie []string `json:"cookie"`
	}
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Printf("解析失败:%v", err)
		return
	}

	fmt.Printf("解析到的json:%v\n", resp)
	fmt.Printf("cookie:%v", resp.Data.Cookie)
	//cookie:
}
