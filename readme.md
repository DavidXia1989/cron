## Installation logging

cron定时任务包
基于github.com/robfig/cron/v3的封装
Env：

golang >= 1.13.0

Install:

```
// 设置环境变量使得go支持私有库
GOPRIVATE=code.zm.shzhanmeng.com
// 安装
go get -u code.zm.shzhanmeng.com/go-common/cron
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
import "code.zm.shzhanmeng.com/go-common/cron"
```