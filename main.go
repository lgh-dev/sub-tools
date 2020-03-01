package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"sub-tools/tools"
	"time"
)

var (
	//projectPath = "D:/test/iot/device/?/src/main/java-single/com/cas/iot/device/?"
	//projectName = "iot"
	//serverName  = "device"
	//不同模块的目录
	//apiPath         = strings.Replace(projectPath, "?", "api", -1)
	//applicationPath = strings.Replace(projectPath, "?", "application", -1)
	//domainPath      = strings.Replace(projectPath, "?", "domain", -1)
	//
	////包名
	//basePackageName = "com.cas.iot.device"
	entityList []string
	entityData []string
	template   Template
)

//接口对象
//type Interface struct {
//	packageName           string            // 报名
//	annotationOfinterface string            // 方法注解
//	interfaceName         string            //public interface DeviceTypeService
//	interfaceMethods      []InterfaceMethod //方法集合
//}

//方法对象
//type InterfaceMethod struct {
//	comment               string // 注释
//	annotationOfinterface string // 方法注解
//	methodName            string //addDeviceType
//	returnDTOName         string //ResultDTO
//	paramsName            string //@RequestBody DeviceTypeDTO deviceTypeDTO
//}

// 代码模板对象
type Template struct {
	CodeType    string
	Model       CodePath
	Application CodePath
	Repository  CodePath
	API         CodePath
}
type CodePath struct {
	InPath  string //输入路径
	OutPath string //输出路径
}

/**
* @Description 代码生成器
* @author liangguohui@iot-cas.com
* @date 2020/2/25
* @version V1.0
 */
func main() {
	//subMethodStr()
	start := time.Now().UnixNano()
	readConfig() //读取配置
	getModels()  // 获取model
	//genApiJavaClassFile()
	end := time.Now().UnixNano()
	fmt.Printf("代码生成耗时：%d微秒\n", (end-start)/1000)
}

// 读取配置文件
func readConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("./config/")
	err := v.ReadInConfig()
	if err != nil {
		panic("Read config err,please check path or format !")
	}
	newV := viper.New()
	for _, key := range v.AllKeys() {
		value := v.GetString(key)
		value = replaceValue(value, v)
		newV.SetDefault(key, value)
	}
	for _, key := range newV.AllKeys() {
		fmt.Printf("%s=%s\n", key, newV.Get(key))
	}
	//初始化模板对象
	initTemplate(newV)

	return v
}

// 替换值
func replaceValue(value string, v *viper.Viper) string {
	if strings.Contains(value, "${") {
		value = strings.ReplaceAll(value, "${sub-tools.var.workspace}", v.GetString("sub-tools.var.workspace"))
		value = strings.ReplaceAll(value, "${sub-tools.var.project}", v.GetString("sub-tools.var.project"))
		value = strings.ReplaceAll(value, "${sub-tools.var.server}", v.GetString("sub-tools.var.server"))
		value = strings.ReplaceAll(value, "${sub-tools.template.javaPath}", v.GetString("sub-tools.template.javaPath"))
		value = strings.ReplaceAll(value, "${sub-tools.template.rootPath}", v.GetString("sub-tools.template.rootPath"))
	}
	return value
}

// 初始化模板对象
func initTemplate(v *viper.Viper) {
	template = Template{}
	template.CodeType = v.GetString("sub-tools.template.codeType")
	template.Model.InPath = v.GetString("sub-tools.template.code.model.inPath")
	template.Model.OutPath = v.GetString("sub-tools.template.code.model.outPath")
	template.Repository.OutPath = v.GetString("sub-tools.template.code.repository.outPath")
	template.Application.OutPath = v.GetString("sub-tools.template.code.application.outPath")
	template.API.OutPath = v.GetString("sub-tools.template.code.api.outPath")
}

// 获取model信息
func getModels() {
	model := tools.Models{}
	entityList, entityData = model.GetModels(template.Model.InPath)
}

// 获取生成repository代码。
func getRepository() {

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

//func genApiJavaClassFile() {
//	for _, entity := range entityList {
//		tools.MkdirIfNotExists(apiPath)
//		//模版引擎替换
//		engine := tools.Engine{KV: map[string]string{
//			"entity": entity, "Author": "lgh", "basePackageName": basePackageName, "entity.toLower()": entity,
//		}}
//		data, err := engine.GetTemplateFile("./template/api_template.txt").Deals()
//		if err != nil {
//			fmt.Printf(err.Error())
//			return
//		}
//		outfilePath := apiPath + "/" + entity + "Api.java-single"
//		fmt.Println("Created API file: " + outfilePath)
//		ioutil.WriteFile(outfilePath, []byte(data), os.FileMode(777))
//	}
//}
