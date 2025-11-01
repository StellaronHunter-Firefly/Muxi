// 完善代码
package main

type Sunjunnan struct {
	Name                  string
	FaceSimilarity        int
	VoiceSimilarity       int
	PersonalitySimilarity int
}

func GetData(sunjunnans []SunJunnan) {
	//补全获取数据逻辑

}

func Sort(sunjunnans []SunJunnan) {
	//补全排序逻辑

}

func OutputData(sunjunnans []SunJunnan) {
	//补全输出逻辑

}

func main() {
	var sunjunnans []SunJunnan

	GetData(nums, name)
	Sort(nums, name)
	OutputData(nums, name)
}

// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// // Sunjunnan 定义一个结构体来存储每个SJN的信息
// type Sunjunnan struct {
// 	Name           string // SJN的名字
// 	FaceSimilarity int    // 外貌相似度
// 	VoiceSimilarity int  // 声音相似度
// 	PersonalitySimilarity int // 性格相似度
// }

// // GetData 获取用户输入的数据并填充到sunjunnans切片中
// func GetData(sunjunnans *[]Sunjunnan) {
// 	var n int
// 	fmt.Print("请输入sjn的个数: ")
// 	fmt.Scan(&n)

// 	for i := 0; i < n; i++ {
// 		var name string
// 		var faceSim, voiceSim, personalitySim int
// 		fmt.Printf("请输入第%d个sjn的信息 (名字 外貌相似度 声音相似度 性格相似度): ", i+1)
// 		fmt.Scan(&name, &faceSim, &voiceSim, &personalitySim)

// 		// 创建一个新的Sunjunnan实例并添加到切片中
// 		*subjunnans = append(*subjunnans, Sunjunnan{
// 			Name:                name,
// 			FaceSimilarity:      faceSim,
// 			VoiceSimilarity:     voiceSim,
// 			PersonalitySimilarity: personalitySim,
// 		})
// 	}
// }

// // Sort 根据外貌相似度、声音相似度和性格相似度对sunjunnans进行排序
// func Sort(sunjunnans []Sunjunnan) {
// 	sort.Slice(sunjunnans, func(i, j int) bool {
// 		// 首先按外貌相似度降序排列
// 		if sunjunnans[i].FaceSimilarity != sunjunnans[j].FaceSimilarity {
// 			return sunjunnans[i].FaceSimilarity > sunjunnans[j].FaceSimilarity
// 		}
// 		// 如果外貌相似度相同，则按声音相似度降序排列
// 		if sunjunnans[i].VoiceSimilarity != sunjunnans[j].VoiceSimilarity {
// 			return sunjunnans[i].VoiceSimilarity > sunjunnans[j].VoiceSimilarity
// 		}
// 		// 如果外貌相似度和声音相似度都相同，则按性格相似度降序排列
// 		return sunjunnans[i].PersonalitySimilarity > sunjunnans[j].PersonalitySimilarity
// 	})
// }

// // OutputData 输出排序后的结果
// func OutputData(sunjunnans []Sunjunnan) {
// 	fmt.Println("排序后的sjn们：")
// 	for _, sjn := range subjunnans {
// 		fmt.Printf("%s: 外貌相似度: %d, 声音相似度: %d, 性格相似度: %d\n",
// 			sjn.Name, sjn.FaceSimilarity, sjn.VoiceSimilarity, sjn.PersonalitySimilarity)
// 	}
// }

// func main() {
// 	var sunjunnans []Sunjunnan

// 	// 获取用户输入的数据
// 	GetData(&subjunnans)

// 	// 对数据进行排序
// 	Sort(subjunnans)

// 	// 输出排序后的结果
// 	OutputData(subjunnans)
// }
/*当然，strconv 包是 Go 语言标准库中的一个重要工具包，主要用于字符串和基本数据类型（如整数、浮点数、布尔值等）之间的转换。它提供了多种函数来处理这些转换操作，确保在不同数据类型之间进行安全和高效的转换。
strconv 包的主要功能
字符串转换为数字：
Atoi(s string) (int, error)：将字符串转换为 int 类型。
ParseInt(s string, base int, bitSize int) (i int64, err error)：将字符串转换为指定位宽的整数。
ParseUint(s string, base int, bitSize int) (uint64, error)：将字符串转换为无符号整数。
ParseFloat(s string, bitSize int) (float64, error)：将字符串转换为浮点数。
数字转换为字符串：
Itoa(i int) string：将 int 类型转换为字符串。
FormatInt(i int64, base int) string：将整数转换为指定进制的字符串。
FormatUint(i uint64, base int) string：将无符号整数转换为指定进制的字符串。
FormatFloat(f float64, fmt byte, prec, bitSize int) string：将浮点数转换为字符串。
其他转换：
Quote(s string) string 和 Unquote(s string) (string, error)：用于转义和取消转义字符串。
AppendQuote(dst []byte, s string) []byte 和 AppendUnquote(dst []byte, s string) ([]byte, error)：类似于 Quote 和 Unquote，但返回字节切片而不是字符串。
IsPrint(r rune) bool：检查一个字符是否为可打印字符。
CanBackquote(s string) bool：检查一个字符串是否可以用反引号括起来而不改变其内容。
使用方法示例
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 示例字符串
	strNum := "12345"
	strFloat := "123.45"
	strBool := "true"

	// 字符串转换为整数
	intNum, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println("Error converting to int:", err)
	} else {
		fmt.Printf("String '%s' converted to int: %d\n", strNum, intNum)
	}

	// 字符串转换为浮点数
	floatNum, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Error converting to float:", err)
	} else {
		fmt.Printf("String '%s' converted to float: %.2f\n", strFloat, floatNum)
	}

	// 字符串转换为布尔值
	boolVal, err := strconv.ParseBool(strBool)
	if err != nil {
		fmt.Println("Error converting to bool:", err)
	} else {
		fmt.Printf("String '%s' converted to bool: %t\n", strBool, boolVal)
	}

	// 整数转换为字符串
	intToString := strconv.Itoa(intNum)
	fmt.Printf("Integer %d converted to string: '%s'\n", intNum, intToString)

	// 浮点数转换为字符串
	floatToString := strconv.FormatFloat(floatNum, 'f', 2, 64)
	fmt.Printf("Float %.2f converted to string: '%s'\n", floatNum, floatToString)

	// 布尔值转换为字符串
	boolToString := strconv.FormatBool(boolVal)
	fmt.Printf("Boolean %t converted to string: '%s'\n", boolVal, boolToString)

	// 其他转换示例
	hexStr := strconv.FormatInt(int64(intNum), 16)
	fmt.Printf("Integer %d converted to hex string: '%s'\n", intNum, hexStr)

	binStr := strconv.FormatUint(uint64(intNum), 2)
	fmt.Printf("Integer %d converted to binary string: '%s'\n", intNum, binStr)
}


代码说明
字符串转换为整数：
使用 strconv.Atoi 将字符串 "12345" 转换为整数 12345。
使用 strconv.ParseInt 将字符串 "12345" 转换为 int64 类型的整数 12345。
字符串转换为浮点数：
使用 strconv.ParseFloat 将字符串 "123.45" 转换为 float64 类型的浮点数 123.45。
字符串转换为布尔值：
使用 strconv.ParseBool 将字符串 "true" 转换为布尔值 true。
整数转换为字符串：
使用 strconv.Itoa 将整数 12345 转换为字符串 "12345"。
浮点数转换为字符串：
使用 strconv.FormatFloat 将浮点数 123.45 转换为字符串 "123.45"，保留两位小数。
布尔值转换为字符串：
使用 strconv.FormatBool 将布尔值 true 转换为字符串 "true"。
其他转换示例：
使用 strconv.FormatInt 将整数 12345 转换为十六进制字符串 "3039"。
使用 strconv.FormatUint 将无符号整数 12345 转换为二进制字符串 "11000000111001"。
通过这些示例，你可以看到 strconv 包提供了丰富的功能来进行不同类型之间的转换，确保在不同的场景下都能方便地进行数据处理。
*/
