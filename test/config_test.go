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
		f.StringVar(&name, "name", "Hank", "当前人名")
		f.Int64Var(&age, "age", 18, "当前人年龄")
	}
	//调用Init之前，先把HandlerFunc设置好
	config.ManagerSet.FlagHandlerFunc = flgCurrent
	config.Init()
	fmt.Println("my name is " + name)
	fmt.Println("from config k-v,my name is " + config.GetVal("name"))
}
