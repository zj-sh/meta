package base

import (
	"github.com/bytedance/sonic"
	"os"
	"runtime"
)

func (s *JsonString) Marshal(val interface{}) error {
	if val == "" {
		return nil
	}
	if d, err := sonic.MarshalString(val); err != nil {
		return err
	} else {
		s.Value = d
	}
	return nil
}
func (s *JsonString) UnMarshal(ptr interface{}) error {
	if s.GetValue() == "" {
		return nil
	}
	return sonic.UnmarshalString(s.GetValue(), ptr)
}
func NewJsonString(val string) *JsonString {
	return &JsonString{Value: val}
}

// GetPageRows
// @Description: 获取分页参数, page默认1, rows默认50, rows最大1000
// @receiver p
// @return int64
// @return int64
func (p *Pages) GetPageRows() (int, int) {
	if p.GetPage() == 0 {
		p.Page = 1
	}
	if p.GetRows() == 0 {
		p.Rows = 50
	}
	if p.GetRows() > 1000 {
		p.Rows = 1000
	}
	return int(p.GetPage()), int(p.GetRows())
}

func (x *Authorization) Map() map[string]interface{} {
	var m = make(map[string]interface{})
	d, _ := sonic.Marshal(x)
	_ = sonic.Unmarshal(d, &m)
	return m
}

func IsDebug() bool {
	return os.Getenv("MORED_DEBUG") == "true"
}
func IsLinux() bool {
	return __goos() == "linux"
}
func IsWindows() bool {
	return __goos() == "windows"
}
func IsMac() bool {
	return __goos() == "darwin"
}
func __goos() string {
	return runtime.GOOS
}
