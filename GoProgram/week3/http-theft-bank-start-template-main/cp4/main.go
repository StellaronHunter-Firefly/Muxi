package main

import (
	"fmt"
	"github.com/weiji6/hacker-support/encrypt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	//网址1
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/organization/code",
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

	m1, err := resp.GetHeader("Map-Fragments")
	if err != nil {
		fmt.Println(err)
	}
	//网址2
	req, err = httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		fmt.Println(err)
	}

	req.SetHeader("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiVEFOR2N5IiwiaWF0IjoxNzYzMTkyNTY1LCJuYmYiOjE3NjMxOTI1NjV9.qQTSCbDxd4yMIf4du_HnZdAiAEYacQr8URhqfcDeo70")

	resp, err = req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	m2, err := resp.GetHeader("Map-Fragments")
	if err != nil {
		fmt.Println(err)
	}
	//网址3
	encrypted, err := encrypt.AESEncryptOutInBase64([]byte("for {go func(){time.Sleep(1*time.Hour)}()}"), []byte("MuxiStudio203304"))
	if err != nil {
		fmt.Println(err)
	}
	req, err = httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate",
		string(encrypted),
		httptool.DEFAULT,
	)

	req.SetHeader("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiVEFOR2N5IiwiaWF0IjoxNzYzMTkyNTY1LCJuYmYiOjE3NjMxOTI1NjV9.qQTSCbDxd4yMIf4du_HnZdAiAEYacQr8URhqfcDeo70")

	resp, err = req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	m3, err := resp.GetHeader("Map-Fragments")
	if err != nil {
		fmt.Println(err)
	}
	//网址4
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

	m4, err := resp.GetHeader("Map-Fragments")
	if err != nil {
		fmt.Println(err)
	}
	//网址5
	req, err = httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		fmt.Println(err)
	}

	req.SetHeader("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiVEFOR2N5IiwiaWF0IjoxNzYzMTkyNTY1LCJuYmYiOjE3NjMxOTI1NjV9.qQTSCbDxd4yMIf4du_HnZdAiAEYacQr8URhqfcDeo70")

	resp, err = req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}

	m5, err := resp.GetHeader("Map-Fragments")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m1, m2, m3, m4, m5)
	//访问新网址
	req, err = httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/muxi/backend/computer/examination",
		"",
		httptool.DEFAULT,
	)
	if err != nil {
		fmt.Println(err)
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
Send request successfully! Please check your response body.
Send request successfully! Please check your response body.
Send request successfully! Please check your response body.
[muxi] [backend] [computer] [examination] []
Send request successfully! Please check your response body.
Message:
response body:
1.Message:
OK
2.Text:
TANGcy，真亏你能来到这里！在你眼前的就是最后的密码门了。
但是密码位数未知，看来只能通过全排列程序去暴力破解。

>示例如下：
============================================
输入：
3
1 2 3
输出：
[[1 2 3][1 3 2][2 1 3][2 3 1][3 1 2][3 2 1]]
============================================

>代码模板:

func permute(nums []int) [][]int {
    // insert your code

}

func main() {
    var n int
        fmt.Scanf("%d", &n)

        testSlice := make([]int, n)
    // 标准输入n个不重复的数字

    res := permute(testSlice)
    fmt.Println(res)
}

请在完成此程序后上传至 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination 来破解最后的密码门
3.ExtraInfo:
*/
