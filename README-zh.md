# Gochan

## 背景
一般情况下，我们可以通过定义一个带缓冲的 channel 变量接收某种事件，然后通过一个专用的 goroutine 消费执行这个 channel 中的事件。

但是如果相关事件很多的时候，一个 goroutine 不够用了怎么办呢？或许我们会想到多创建几个专用的 goroutine 来并发地消费执行这个 channel 中的事件；如果 channel 中各个事件之间是独立的，是可行的，但是如果某些事件之间具有某种顺序上的约束，那么就需要对事件进行特定的分类。

比如，一个订单的支付与货物发货，两个事件是需要保序的；但是不同的订单之间又是可以并发执行；其实就是实现一个微型的按特定主题分类的 pub-sub（发布-订阅）系统。以订单为例，可以根据订单单号，把相同单号的事件推送到同一个队列（channel），一个特定的执行器（goroutine）来消费执行这个队列中的事件，如此平行扩展多个类似的组合，实现并发。


## 示意

### 平常的设计
```bash

event ->
        |
event -> buffer-channel -> goroutine
        |
event ->

```

### 状态无依赖的并发设计

事件之间完全没有状态依赖，因此可以简单扩展 goroutine 进行加快事件执行速度。

```bash

event ->                 ->goroutine
        |               |
event -> buffer-channel -> goroutine
        |               |
event ->                 ->goroutine

```

### 状态存在依赖的并发设计

引入一层分发器（dispatcher），根据某个特性（比如 uuid）把事件分发到相应的队列（buffer-channel）中。

```bash

event ->              -> buffer-channel -> goroutine
        |            |
event --> dispatcher -> buffer-channel -> goroutine
        |            |
event ->              -> buffer-channel -> goroutine

```