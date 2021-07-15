# config
## 特点
- 可同时兼容 环境变量参数、命令行参数、多配置文件参数
- 具有优先级特性
- 使用方便，兼容flag



- 使用方法

默认配置路径 conf/app.properties

1. 如何引用 
main.go 中直接引用
package main

import (
	"fmt"
	"github.com/genghongjie/config"//必要
)
//非必要
var name string
//非必要
func init() {
	config.FlagSet.StringVar(&name, "name", "Jack", "人名")

}

func main() {
    //必要
	config.Init()
	
	fmt.Printf("hello %s", config.GetVal("name"))
	fmt.Println()
}

```

2. 查询配置中的key值
```cassandraql
# 查询key config的值 string类型
    config.GetVal("config")
# 默认值方式 string类型
    config.GetVal("config","conf/app.properties")


# 查询 整型  int类型
  config.GetValInt("age")
  config.GetValInt("age",2)
```