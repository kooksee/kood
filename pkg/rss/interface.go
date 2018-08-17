package rss

import "time"

type IRss interface {
	// 添加源
	AddSource(url string) error

	// 更新源内容
	UpdateSource(url string) error

	// 定时更新源内容
	WatchSource(url string, tm time.Duration) error

	// 删除源地址
	DeleteSource(url string) error
}
