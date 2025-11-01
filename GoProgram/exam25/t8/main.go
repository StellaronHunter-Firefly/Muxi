package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostExample() {
	postData, _ := json.Marshal(map[string]interface{}{
		"task":   "test",
		"body":   "content",
		"userId": 1,
	})
	postReq, _ := http.NewRequest("POST", "http://weijiblog.top/start", bytes.NewBuffer(postData))
	client := &http.Client{}
	resp, _ := client.Do(postReq)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
