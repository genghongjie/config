package test

import (
	"flag"
	"fmt"
	"genghongjie/config"
	"testing"
)

func TestInit(t *testing.T) {
	var name string
	var age int64

	flgCurrent := func(f *flag.FlagSet) {
		f.StringVar(&name, "名称", "Hank", "当前人名")
		f.Int64Var(&age, "年龄", 18, "当前人年龄")
	}
	//调用Init之前，先把HandlerFunc设置好
	config.ManagerSet.FlagHandlerFunc = flgCurrent

	config.Init()
	fmt.Println("my name is " + name)
}

func TestInit_init(t *testing.T) {
	config.Init()
}
