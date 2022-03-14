## Installation logging

cron定时任务包
基于github.com/robfig/cron/v3的封装
Env：

golang >= 1.13.0

Install:

```
// 安装
go get -u github.com/DavidXia1989/cron
```

``
    StartByGraceful 外部传递信号， 平滑退出
    Start   阻塞式运行
``
```
以秒级别
Cron.AddFunc("*/10 * * * * ?", func() {
		func() {
			// code
		}()
	})
```

Import:

```go
import "github.com/DavidXia1989/cron"
```