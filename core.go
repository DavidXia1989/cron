package cron

import (
	"time"
	"github.com/robfig/cron/v3"
	"context"
)

// *cron.Cron
var cron_v3 = cron.New(cron.WithParser(cron.NewParser(cron.Second | cron.Minute |
	cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)), cron.WithChain())



// 根据signal 等待正在运行的定时任务执行结束后，返回退出
func StartByGraceful(ctx context.Context) {
	//在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制
	cron_v3.Start()
	//会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
	t1 := time.NewTimer(time.Second * 10)
	for {
		//阻塞 select 等待 channel
		select {
		case <- ctx.Done():
			for {
				select {
				case <-cron_v3.Stop().Done():
					return
				}
			}
		case <-t1.C:
			//会重置定时器，让它重新开始计时
			t1.Reset(time.Second * 10)
		}
	}
}

// 阻塞运行
func Start() {
	//在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制
	cron_v3.Start()
	//会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
	t1 := time.NewTimer(time.Second * 10)
	for {
		//阻塞 select 等待 channel
		select {
		case <-t1.C:
			//会重置定时器，让它重新开始计时
			t1.Reset(time.Second * 10)
		}
	}
}

// 添加定时任务 穿一个定时任务包
func AddFunc(spec string, cmd func()) (int, error) {
	entryId, err:=  cron_v3.AddFunc(spec, cmd)
	return int(entryId),err
}

// 添加定时任务 传递一个run的抽象接口
func AddJob(spec string, job Job) (int, error) {
	entryId, err:=  cron_v3.AddJob(spec, job)
	return int(entryId),err
}

// 获取内部*Cron的实例
func GetCron() *cron.Cron {
	return cron_v3
}