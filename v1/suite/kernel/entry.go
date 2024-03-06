package kernel

import (
	"github.com/hashicorp/go-plugin"
	"github.com/zj-sh/meta/v1/suite"
	"sync"
	"time"
)

/**
 * @description: no use database
 */

type Entry struct {
	pid     string            // 进程ID
	code    string            // 套件编码
	version string            // 套件版本
	start   time.Time         // 启动时间
	last    time.Time         // 最后一次运行时间
	delay   time.Duration     // 延迟关闭时间
	grpc    *plugin.Client    // GRPC客户端
	client  suite.SuiteClient // 套件客户端
	busy    int64             // 忙碌状态
	used    int64             // 使用次数
	file    string            // 套件文件
	sync.Mutex
}

func (e *Entry) Client() suite.SuiteClient {
	e.Lock()
	defer e.Unlock()
	e.busy++
	e.used++
	e.last = time.Now()
	return e.client
}
func (e *Entry) Release() {
	e.Lock()
	defer e.Unlock()
	e.last = time.Now()
	e.busy--
}
