package main

//补充包
import (
	"fmt"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time // 记录上次消费的时间
	msgs := make(chan MSG)

	//这里源源不断的生产信息
	go func() {
		for i := 0; ; i++ {
			msgs <- MSG{
				ms: message{
					Topic:     "消费主题",
					Partition: 0,
					Offset:    0,
				},
				feedEvent: FeedEventDM{
					Type:    "grade",
					UserID:  i,
					Title:   "成绩提醒",
					Content: "您的成绩是xxx",
				},
			}
			//每次发送信息会停止0.01秒以模拟真实的场景
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//不断接受消息进行消费
	for msg := range msgs {
		// 添加新的值到events中
		consumeMSG = append(consumeMSG, msg)
		// 如果数量达到额定值就批量消费
		if len(consumeMSG) >= ConsumeNum {
			//进行异步消费
			go func(ms []MSG) {
				fn(ms)
			}(consumeMSG[:ConsumeNum:ConsumeNum]) //添加容量限制

			// go func() {
			// 	m := consumeMSG[:ConsumeNum]
			// 	fn(m)
			// }()
			// 闭包操作存在风险，可能访问到已修改的数据
			// 更新上次消费时间
			lastConsumeTime = time.Now()
			// 清除插入的数据
			consumeMSG = nil
		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
			// 如果距离上次消费已经超过5分钟且有未处理的消息
			if len(consumeMSG) > 0 {
				//进行异步消费
				go func(ms []MSG) {
					fn(ms)
				}(consumeMSG) //直接传递整个切片，避免切片操作
				// go func() {
				// 	m := consumeMSG[:ConsumeNum]
				// 	fn(m)
				// }()
				// 更新上次消费时间
				lastConsumeTime = time.Now()
				// 清空插入的数据
				consumeMSG = nil
				// consumeMSG = consumeMSG[ConsumeNum:]
			}
		}
	}
}

func fn(m []MSG) {
	fmt.Printf("本次消费了%d条消息\n", len(m))
}

/*
panic: runtime error: slice bounds out of range [:5] with capacity 3
原因出自切片越界了，在go语言中，当你对一个切片进行切片操作时，如果你尝试访问的范围超出了切片的当前长度或容量，就会引发这个运行时错误。
goroutine 18 [running]:
main.main.func2()
        D:/GoProgram/week2/PanicAnalyze/main.go:62 +0x3c
created by main.main in goroutine 1
        D:/GoProgram/week2/PanicAnalyze/main.go:61 +0x3c8
exit status 2
*/
