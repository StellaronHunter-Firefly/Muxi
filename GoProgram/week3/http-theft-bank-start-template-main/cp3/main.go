package main

import (
	"fmt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample",
		"",
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

	resp.Save("eye.png")

	req, err = httptool.NewRequest(
		httptool.POSTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate",
		"eye.png",
		httptool.FILE,
	)
	if err != nil {
		panic(err)
	}

	req.SetHeader("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiVEFOR2N5IiwiaWF0IjoxNzYzMTkyNTY1LCJuYmYiOjE3NjMxOTI1NjV9.qQTSCbDxd4yMIf4du_HnZdAiAEYacQr8URhqfcDeo70")

	resp, err = req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	resp.ShowBody()
}

/*
Send request successfully! Please check your response body.
Send request successfully! Please check your response body.
Message:
response body:
1.Message:
OK
2.Text:
你现在已经到第二扇门了，是虹膜识别安全门。
你需要向组织请求已准备好的虹膜样本，访问 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample 下载图片。
再将此图片上传至 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate 以破解此门，加油！
3.ExtraInfo:
*/
/*
Send request successfully! Please check your response body.
Send request successfully! Please check your response body.
Message:
response body:
1.Message:
OK
2.Text:
还剩最后一道门了！
我们需要银行结构图碎片，这些碎片就隐藏在前面某四个路由的响应头中，位于 map-fragments 字段。
将它们用"/"拼起来就是最后一道门的所在位置！注意response的信息。
3.ExtraInfo:
*/
