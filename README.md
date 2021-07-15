# config

- 使用方法

1. 如何引用 
main.go 中直接引用
```goland
import (
	"fmt"
	_ "github.com/genghongjie/config"
)
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