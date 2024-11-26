# 第六章并发编程：6.1上下文context

## 基础概念
context实际上是上下文的意思，上下文 context.Context Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体。上下文与 Goroutine 有比较密切的关系，是 Go 语言中独特的设计，在其他编程语言中我们很少见到类似的概念。
## 设计原理
在 Goroutine 构成的树形结构中对信号进行同步以减少计算资源的浪费是 context.Context 的最大作用。Go 服务的每一个请求都是通过单独的 Goroutine 处理的2，HTTP/RPC 请求的处理器会启动新的 Goroutine 访问数据库和其他服务。
## 方法
#### context定义了4个需要实现的方法
| 方法名      | 含义                                                                     | 使用示例 |
|----------|------------------------------------------------------------------------|------|
| Deadline | 返回 context.Context 被取消的时间，也就是完成工作的截止日期                                 ||
| Done     | 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 Done 方法会返回同一个 Channel ||
| Err      | 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值             ||
| Value    | 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果   ||
#### context其他方法:能返回私有结构体的方法
|方法名|含义|示例|
|---|---|---|
|context.Background|||
|context.TODO|||
|context.WithDeadline|||
|context.WithValue|||
