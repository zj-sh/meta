package suite

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/zohu/zlog"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

/**
 * @description: no use database
 */

// 套件运行管理器
type manager struct {
	path   string            // 套件路径
	max    int               // 最大运行数量
	suites map[string]*Entry // 套件列表
	sync.Mutex
}

// run
// @Description: PS: 必须搭配Lock使用
// @receiver m
// @param rid
// @param e
func (m *manager) run(rid string, e *Entry) {
	m.suites[rid] = e
}

// load
// @Description:  PS: 必须搭配Lock使用
// @receiver m
// @param rid
// @return *Entry
// @return bool
func (m *manager) load(rid string) (*Entry, bool) {
	if e, ok := m.suites[rid]; ok {
		if _, err := os.Stat(e.file); os.IsNotExist(err) {
			delete(m.suites, rid)
			return nil, false
		}
		return e, true
	}
	return nil, false
}
func (m *manager) shutdown(rid string) {
	m.Lock()
	defer m.Unlock()
	if e, ok := m.suites[rid]; ok {
		if e.grpc != nil {
			zlog.Infof("套件%s[%s] 正在关闭，运行期间共使用%d次", e.code, e.version, e.used)
			if protocol, err := e.grpc.Client(); err == nil {
				_ = protocol.Close()
			}
			e.grpc.Kill()
			zlog.Infof("套件%s[%s] 已关闭", e.code, e.version)
		}
		delete(m.suites, rid)
	}
}
func (m *manager) shutdownPid(pid string) {
	for rid, e := range m.suites {
		if e.pid == pid {
			m.shutdown(rid)
			break
		}
	}
}
func (m *manager) shutdownAll() {
	for rid, _ := range m.suites {
		m.shutdown(rid)
	}
}
func (m *manager) isMaxRunning() bool {
	return m.max != 0 && len(m.suites) >= m.max
}
func (m *manager) onStatus() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	var last time.Time
	t1 := time.NewTicker(time.Second)
	for {
		select {
		case <-t1.C:
			if last.IsZero() {
				last = time.Now()
			}
			var running []*Entry
			for _, e := range m.suites {
				if (e.busy <= 0 && e.delay != time.Duration(0) && time.Now().Sub(e.last) > e.delay) || e.grpc.Exited() {
					zlog.Infof("套件%s[%s] 闲置超过延时(%s)，即将自动关闭!", e.code, e.version, e.delay.String())
					m.shutdown(UniqueSuiteFileName(e.code, e.version))
				} else {
					running = append(running, e)
				}
			}
			count := time.Now().Sub(last)
			if (len(running) > 0 && count > time.Minute) || (len(running) == 0 && count > time.Hour) {
				zlog.Infof("当前运行套件: \n%s", fmtSuiteInfo(running))
			}
			last = time.Now()
		case <-quit:
			m.shutdownAll()
		}
	}
}

// UniqueSuiteFileName
// @Description: 获取套件唯一文件名称[套件编码+版本号]
// @param suiteCode
// @param version
// @return string
func UniqueSuiteFileName(suiteCode, version string) string {
	return fmt.Sprintf("%s_%s", suiteCode, strings.ReplaceAll(version, ".", "_"))
}

func fmtSuiteInfo(infos []*Entry) string {
	table, _ := gotable.Create("PID", "CODE", "VERSION", "START", "LAST", "BUSY", "USED", "DELAY")
	for _, e := range infos {
		delay := e.delay
		if e.busy <= 0 {
			delay = delay - time.Now().Sub(e.last)
		}
		_ = table.AddRow([]string{
			e.pid,
			e.code,
			e.version,
			e.start.Format("2006-01-02 15:04:05"),
			e.last.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%d", e.busy),
			fmt.Sprintf("%d", e.used),
			delay.String(),
		})
	}
	return table.String()
}
