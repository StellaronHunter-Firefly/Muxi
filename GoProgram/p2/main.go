package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func GetExample() {
// 	req, _ := http.NewRequest("GET", "http://weijiblog.top/download", nil)
// 	req.Header.Set("Accept", "application/json")

// 	client := &http.Client{}
// 	resp, _ := client.Do(req)
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// }

// func PostExample() {
// 	postData, _ := json.Marshal(map[string]interface{}{
// 		"title":  "test",
// 		"body":   "content",
// 		"userId": 1,
// 	})
// 	postReq, _ := http.NewRequest("POST", "http://weijiblog.top/start", bytes.NewBuffer(postData))
// 	client := &http.Client{}
// 	resp, _ := client.Do(postReq)
// 	defer resp.Body.Close()
// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// }
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://weijiblog.top/download"
	payload := strings.NewReader("name=Tom&age=18")
	resp, err := http.Post(url, "http://weijiblog.top/download", payload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// func main() {
// 	GetExample()
// 	//PostExample()
// }

//cd D:\GoProgram\p2
