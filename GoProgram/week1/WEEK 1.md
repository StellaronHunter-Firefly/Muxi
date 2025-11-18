# Markdown尝试
## 分级标题
### 标题1
#### 副标题1
### 标题2
#### 副标题2
## 代码块尝试
```
package main

import (
    "fmt"
)

func main() {
    fmt.Fprintln("Hello,World!")
}
```
## 数字/圆点列表
1. 1
2. 2
3. 3

- 1
- 2
- 3

## 字体加粗，删除线
**这**是**一**段~~文字~~
# 学习go的数组、切片
- 给定s := make([]byte, 5)，那么len(s) 和 cap(s) 分别是多少？令s = s[2:4]，len(s) 和 cap(s) 又分别是多少？
- 比较字符串"hello，世界"的长度和for range该字符串的循环次数
```
package main

import (
    "fmt"
)

func main() {
    s := make([]byte, 5)
    fmt.Printf("len:%v, cap:%v\n", len(s), cap(s))
    slices1 := "hello，世界"
    var i int
    for k := range slices1 {
        i = k
    }

    fmt.Printf("长度：%v, 循环次数：%v\n", len(slices1), i)
    if len(slices1) > i {
        fmt.Println("长度大于循环次数")
    } else if len(slices1) < i {
        fmt.Println("长度大于循环次数")
    } else {
        fmt.Println("长度等于循环次数")
    }
}

```