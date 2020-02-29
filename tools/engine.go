package tools

import (
	"errors"
	"io/ioutil"
	"strings"
	"time"
)

const label = "{{?}}"

//模版引擎
type Engine struct {
	KV   map[string]string
	data string
}

//替换参数 key:value结构
func (engine *Engine) Add(k string, v string) *Engine {
	engine.KV[k] = v
	return engine
}

//获取参数值
func (engine *Engine) get(k string) string {
	return engine.KV[k]
}

func toLabel(key string) string {
	return strings.Replace(label, "?", key, -1)
}

//获取模版内容
func (engine *Engine) GetTemplateFile(templatePath string) *Engine {

	dataByte, err := ioutil.ReadFile(templatePath)
	if err == nil {
		engine.data = string(dataByte)
	}
	return engine
}

// 替换date
func (engine *Engine) dealDate() *Engine {
	engine.data = strings.Replace(engine.data, toLabel("date"), time.Now().Format("2006-01-02 15:04:05"), -1)
	return engine
}

//执行替换map里面的所有参数。
func (engine *Engine) Deals() (string, error) {
	for k, v := range engine.KV {
		engine.Deal(k, v)
	}
	// 默认替换时间
	engine.dealDate()

	if len(engine.data) == 0 {
		return engine.data, errors.New("文件为空，请检查路径 !")
	}
	return engine.data, nil
}

func (engine *Engine) Deal(k string, v string) {
	if strings.Contains(k, "toLower()") {
		v = strings.ToLower(v)
	}
	engine.data = strings.Replace(engine.data, toLabel(k), v, -1)
}
