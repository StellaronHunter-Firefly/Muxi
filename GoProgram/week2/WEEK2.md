# Go深度学习

## map和slice内存扩容的区别

### Slice

- 当容量小于1024时，新容量为原容量的2倍。
    
- 当容量大于等于1024时，新容量为原容量的1.25倍。
    

### Map

- 等量扩容：不扩大容量，重新排列键值对，提高桶利用率。
    
- 渐进式扩容：每次操作尝试搬迁桶，逐步完成扩容。
    

## 结构体的比较

- 可比较的结构体：所有字段都是可比较的类型时，可以直接使用 `==` 或 `!=` 进行比较。
    
- 不可比较的结构体：包含不可比较字段（如切片、映射、函数等）时，需要编写自定义的比较函数来进行比较。
    

# 闭包计数器

见文件夹counter

# Panic问题分析与解决

## panic内容与原因

panic: runtime error: slice bounds out of range [:5] with capacity 3

原因出自切片越界了，当对切片进行切片操作时，如果尝试访问的范围超出了切片的当前容量，就会引发该错误。

## 更改

```Go
go func(ms []MSG) {
    fn(ms)
}(consumeMSG[:ConsumeNum:ConsumeNum]) //添加容量限制
// go func() {
//    m := consumeMSG[:ConsumeNum]
//    fn(m)
// }()
// 闭包操作存在风险，可能访问到已修改的数据
            
```

```Go
go func(ms []MSG) {
    fn(ms)
}(consumeMSG) //直接传递整个切片，避免切片操作
// go func() {
//    m := consumeMSG[:ConsumeNum]
//    fn(m)
// }()
```

## 两种代码的区别和使用闭包的风险

为何是

```Plain
go func(ms []MSG) {
    fn(ms)
}(consumeMSG[:ConsumeNum:ConsumeNum]) 
```

而不是

```Plain
go func() {
    m := consumeMSG[:ConsumeNum:ConsumeNum]
    fn(m)
}()
```

### 关键区别

#### 执行时机不同

##### 方式一（参数传递）：

- 切片操作 consumeMSG[:ConsumeNum:ConsumeNum]在goroutine启动前立即执行
    
- 参数值在调用时就已经确定
    

##### 方式二（闭包捕获）：

- 切片操作 consumeMSG[:ConsumeNum:ConsumeNum]在goroutine内部执行
    
- 执行时机有延迟，可能在其他goroutine修改数据之后
    

#### 数据竞争风险

##### 方式一（参数传递）- 更安全：

```Go
// 切片在goroutine启动前创建，传递的是副本
batch := consumeMSG[:ConsumeNum:ConsumeNum]  // 立即执行
go func(ms []MSG) {
    // 这里使用的是创建时的切片快照
    fn(ms)  // 安全的
}(batch)
```

##### 方式二（闭包捕获）- 有风险：

```Go
go func() {
    // 这里才执行切片操作，此时consumeMSG可能已被其他goroutine修改！
    m := consumeMSG[:ConsumeNum:ConsumeNum]  // 延迟执行
    fn(m)  // 可能访问到已修改的数据
}()
```

#### 具体风险场景

假设执行顺序：

时间线：

t1: 主goroutine创建goroutine（方式二）

t2: 主goroutine执行 consumeMSG = consumeMSG[ConsumeNum:]

t3: 新goroutine才执行 m := consumeMSG[:ConsumeNum:ConsumeNum] ← 可能panic！

方式一避免了这个风险，因为切片操作在修改`consumeMSG`之前就完成了。  

# 随机休眠

见文件夹RandomSleep

# 交替打印

见文件夹printer