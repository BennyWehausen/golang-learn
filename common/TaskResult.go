package common

import "time"

// 任务执行结果结构体
type TaskResult struct {
	ID       int
	Duration time.Duration
}
