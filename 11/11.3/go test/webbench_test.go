package gotest

import (
	"testing"
)

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	// 做一些初始化的工作，例如读取文件数据，数据库连接之类的
	// 这些时间不影响测试函数本身的性能
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
