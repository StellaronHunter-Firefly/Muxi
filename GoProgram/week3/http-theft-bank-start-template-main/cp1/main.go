package main

import (
	"fmt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/organization/code",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	passport, err := resp.GetHeader("Passport")
	if err != nil {
		fmt.Println(err)
	}

	req, err = httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/organization/secret_key",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		panic(err)
	}

	req.SetHeader("Passport", passport[0])

	resp, err = req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	resp.ShowBody()
}

//Passport : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiVEFOR2N5IiwiaWF0IjoxNzYzMTkyNTY1LCJuYmYiOjE3NjMxOTI1NjV9.qQTSCbDxd4yMIf4du_HnZdAiAEYacQr8URhqfcDeo70
/*
Send request successfully! Please check your response body.
Send request successfully! Please check your response body.
Message:
response body:
1.Message:
OK
2.Text:
恭喜你拿到了 passport，现在你可以着手准备骇入银行。
银行的第一道门是代码安全门，我们计划将错误代码写入此门来破解它。
但是这个门具有识别明文代码的功能，所以我们还需要一个密钥加密我们的错误代码，再写入至此门。
不需要担心，两者我们都为你提供了，你只需要解析 response 中的密文（在 ExtraInfo 中）来得到它们。
你现在的任务：
解析密文，获取 error_code 和 secret_key
使用 secret_key 加密 error_code
然后将加密过的 error_code 放入 请求body 并以 "正确的请求方法" 发送至 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate , 同时注意response的信息。
3.ExtraInfo:
c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ==
*/
