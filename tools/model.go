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
type Models struct {
	ModelNames []string // 模型名称
	ModelData  []string
}

func (models *Models) GetModels(path string) ([]string, []string) {
	modelPath, err := ioutil.ReadDir(path)
	if err != nil {
		panic("model path err:" + err.Error())
	}
	models.ModelNames = make([]string, len(modelPath))
	models.ModelData = make([]string, len(modelPath))
	for i, info := range modelPath {
		if !info.IsDir() {

			if strings.Split(info.Name(), ".")[1] != "java" {
				continue
			}

			models.ModelNames[i] = strings.Split(info.Name(), ".")[0]
			dataByte, err := ioutil.ReadFile(path + "/" + info.Name())
			if err != nil {
				fmt.Errorf("Read model file %s err :", info.Name())
				panic(err.Error())
			}
			models.ModelData[i] = string(dataByte)
		}
	}
	fmt.Println(models.ModelNames)
	return models.ModelNames, models.ModelData
}
