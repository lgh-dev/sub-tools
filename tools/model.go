package tools

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/**
* @Description 获取model信息
* @author liangguohui@iot-cas.com
* @date 2020/2/26
* @version V1.0
 */
type Model struct {
	ModelNames []string // 模型名称
}

func (model *Model) GetModels(path string) []string {
	models, err := ioutil.ReadDir(path)
	if err != nil {
		panic("model path err")
	}
	model.ModelNames=make([]string, len(models))
	for i, info := range models {
		if !info.IsDir(){
			model.ModelNames[i] = strings.Split(info.Name(),".")[0]
		}
	}
	fmt.Println(model.ModelNames)
	return model.ModelNames
}
