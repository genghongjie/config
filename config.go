package config

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	m := ManagerFlag{}
	//命令行参数获取
	m.initFlag()
	//根据命令行参数中的配置路径，文件名称优先级为 环境变量、命令行参数 读取配置文件
	m.loadConfFile()
	//缓存命令行参数值  符合 -key=value --key=value 格式的命令行参数
	m.loadCommand()
	//缓存flag参数
	m.loadFlag()

	m.printKeyValue()
}

//缓存所有的配置参数的结构体
var confMap = make(map[string]string)

const (
	//环境变量中制定config路径
	confFile = "CONFIG_FILE"
)

type ManagerFlag struct {
	//配置文件路径
	Config string `json:"config"`

	//可继续添加其它的flag
	//TODO
}

func (m *ManagerFlag) initFlag() {
	flag.CommandLine = flag.NewFlagSet("请检查命令行参数", flag.ContinueOnError)
	flag.CommandLine.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "使用说明 %s:\n\n", "优先级为 环境变量>命令行参数>配置文件，其中环境变量修改后程序不用重启")
		flag.PrintDefaults()
	}

	flag.StringVar(&m.Config, "config", "conf/app.properties", "配置文件路径 相对路径或者绝对路径，支持多个文件 可用逗号分割")

	flag.Parse()
}

//设置配置信息
func SetValue(key, value string) {
	confMap[key] = value
}

//根据key查询出数字
func GetValInt(key string, defaultVal ...int) (int, error) {
	env, ok := os.LookupEnv(key)
	if ok {
		evnVal, err := strconv.Atoi(env)
		if err != nil {
			return evnVal, nil
		}
	}

	val := confMap[key]
	if defaultVal != nil && len(defaultVal) > 0 {
		if len(val) == 0 {
			return defaultVal[0], nil
		}
	}
	if len(val) == 0 {
		return 0, nil
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, err

	}
	return v, nil
}
func GetVal(key string, defaultVal ...string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}

	val := confMap[key]
	if defaultVal != nil && len(defaultVal) > 0 {
		if val == "" {
			return defaultVal[0]
		}
	}
	return val
}

func loadConfigFromFile(path string) {
	if len(path) == 0 {
		return
	}
	//打开文件指定目录，返回一个文件f和错误信息
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
	}

	//异常处理 以及确保函数结尾关闭文件流
	if err != nil {
		log.Println("读取配置文件失败,请检查文件路径: " + path)
		return
	}

	//创建一个输出流向该文件的缓冲流*Reader
	r := bufio.NewReader(f)
	for {
		//读取，返回[]byte 单行切片给b
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		//去除单行属性两端的空格
		s := strings.TrimSpace(string(b))
		//fmt.Println(s)

		//判断等号=在该行的位置
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		//取得等号左边的key值，判断是否为空
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}

		//取得等号右边的value值，判断是否为空
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		//这样就成功吧配置文件里的属性key=value对，成功载入到内存中c对象里
		confMap[key] = value
	}
}

//The command line
func (m *ManagerFlag) loadCommand() {
	for _, v := range os.Args {
		constrainFlag := strings.Contains(v, "=")

		if constrainFlag && strings.Count(v, "=") == 1 {
			sps := strings.Split(v, "=")

			key := strings.TrimSpace(strings.Replace(sps[0], "-", "", -1))
			value := strings.TrimSpace(sps[1])

			SetValue(key, value)
		}
	}

}
func (m *ManagerFlag) loadConfFile() {
	confFileName, ok := os.LookupEnv(confFile)
	if ok {
		paths := strings.Split(confFileName, ",")
		for _, v := range paths {
			loadConfigFromFile(v)
		}
		return
	}

	paths := strings.Split(m.Config, ",")
	for _, v := range paths {
		loadConfigFromFile(v)
	}
}

func (m *ManagerFlag) printKeyValue() {
	log.Println("----------------")
	log.Println("当前配置信息:")
	for key, value := range confMap {
		log.Println(fmt.Sprintf(" %s:%s", key, value))
	}
	log.Println("----------------")
}

func (m *ManagerFlag) loadFlag() {
	SetValue("config", m.Config)
}
