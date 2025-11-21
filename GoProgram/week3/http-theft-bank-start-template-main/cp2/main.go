package main

import (
	"fmt"
	"github.com/weiji6/hacker-support/encrypt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	ori, err := encrypt.Base64Decode("c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ==")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ori))
	encrypted, err := encrypt.AESEncryptOutInBase64([]byte("for {go func(){time.Sleep(1*time.Hour)}()}"), []byte("MuxiStudio203304"))
	if err != nil {
		fmt.Println(err)
	}
	req, err := httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate",
		string(encrypted),
		httptool.DEFAULT,
	)

	if err != nil {
		fmt.Println(err)
	}
	req.SetHeader("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiVEFOR2N5IiwiaWF0IjoxNzYzMTkyNTY1LCJuYmYiOjE3NjMxOTI1NjV9.qQTSCbDxd4yMIf4du_HnZdAiAEYacQr8URhqfcDeo70")

	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	resp.ShowBody()
}

/*
secret_key:MuxiStudio203304, error_code:for {go func(){time.Sleep(1*time.Hour)}()}
http 404. we can not find the path, did you input the right information? the wrong message is:
secret_key:MuxiStudio203304, error_code:for {go func(){time.Sleep(1*time.Hour)}()}
Send request successfully! Please check your response body.
Message:
response body:
1.Message:
OK
2.Text:
干的漂亮！你已经突破了第一扇门，请继续访问 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate 。
3.ExtraInfo:
*/
