package main

import (
	"fmt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.POSTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
		"./code/main.go",
		httptool.FILE,
	)
	if err != nil {
		panic(err)
	}

	req.SetHeader("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiVEFOR2N5IiwiaWF0IjoxNzYzMTkyNTY1LCJuYmYiOjE3NjMxOTI1NjV9.qQTSCbDxd4yMIf4du_HnZdAiAEYacQr8URhqfcDeo70")

	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	resp.ShowBody()
}

/*
Send request successfully! Please check your response body.
Message:
response body:
1.Message:
OK
2.Text:
END
我就知道你能成功！Backend组织欢迎你！
3.ExtraInfo:
*/
