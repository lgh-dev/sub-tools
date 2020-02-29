package main

import (
	"./tools"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	projectPath = "D:/test/iot/device/?/src/main/java/com/cas/iot/device/?"
	projectName = "iot"
	serverName  = "device"
	//不同模块的目录
	apiPath         = strings.Replace(projectPath, "?", "api", -1)
	applicationPath = strings.Replace(projectPath, "?", "application", -1)
	domainPath      = strings.Replace(projectPath, "?", "domain", -1)

	//包名
	basePackageName = "com.cas.iot.device"
	entityList   []string
)

//接口对象
type Interface struct {
	packageName           string            // 报名
	annotationOfinterface string            // 方法注解
	interfaceName         string            //public interface DeviceTypeService
	interfaceMethods      []InterfaceMethod //方法集合
}

//方法对象
type InterfaceMethod struct {
	comment               string // 注释
	annotationOfinterface string // 方法注解
	methodName            string //addDeviceType
	returnDTOName         string //ResultDTO
	paramsName            string //@RequestBody DeviceTypeDTO deviceTypeDTO
}

/**
* @Description 代码生成器
* @author liangguohui@iot-cas.com
* @date 2020/2/25
* @version V1.0
 */
func main() {
	//subMethodStr()
	start:=time.Now().UnixNano()
	initFunc()
	genApiJavaClassFile()
	end:=time.Now().UnixNano()
	fmt.Printf("代码生成耗时：%d毫秒\n",(end-start)/1000000)
}

func initFunc()  {
	model:=tools.Model{}
	entityList=model.GetModels("D:/test/iot/device/model")
}
/**
快速从接口类中筛选出接口方法。
*/
//func subMethodStr() {
//	con, err := ioutil.ReadFile("D:/test/new_file.md")
//	if err == nil {
//		strs := strings.Split(string(con), "\n")
//		var newstrs string
//		for _, str := range strs {
//			if strings.Contains(str, ";") || strings.Contains(str, "*") {
//				newstrs += str + "\n"
//			}
//		}
//		ioutil.WriteFile("D:/test/new_file2.md", []byte(newstrs), os.FileMode(777))
//	}
//}

func genApiJavaClassFile() {
	for _, entity := range entityList {
		tools.MkdirIfNotExists(apiPath)
		//模版引擎替换
		engine := tools.Engine{KV: map[string]string{
			"entity": entity, "Author": "lgh", "basePackageName": basePackageName, "entity.toLower()": entity,
		}}
		data, err := engine.GetTemplateFile("./codetemplate/api_template.txt").Deals()
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		outfilePath := apiPath + "/" + entity + "Api.java"
		fmt.Println("Created API file: " + outfilePath)
		ioutil.WriteFile(outfilePath, []byte(data), os.FileMode(777))
	}
}
