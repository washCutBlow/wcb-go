package conf

import (
	"flag"
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
)

var (
	ENV *string
	CONF map[string]interface{}

)

// 解析命令行启动参数，设置全局环境变量
func init() {
	ENV = flag.String("e","test","dev|test|prod")
	flag.Parse()

	if *ENV != "test" && *ENV != "dev" && *ENV != "prod"{
		*ENV = "develop"
	}
}

// 解析配置文件
func init()  {
	filename := fmt.Sprintf("conf/%s.ini",*ENV)
	cfg, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		fmt.Println("读取配置文件错误")
		os.Exit(0)
	}
	//c,err := cfg.GetValue("develop","ceshi")
	fmt.Println(cfg.GetSectionList())
}

func init()  {
	fmt.Println(3)
}

func GetConf()  {

}

func GetDbConf() {
	
}