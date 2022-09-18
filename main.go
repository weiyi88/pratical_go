package main

// import 这里我习惯把官方库，开源库，本地module依次分开列出
import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"gin_demo/config"
	"gin_demo/connect/db"
	"gin_demo/router"
)

var (
	conf = pflag.StringP("config", "c", "", "config filepath")
)

func main() {
	pflag.Parse()

	// 初始化配置
	if err := config.Run(*conf); err != nil {
		panic(err)
	}

	// 连接mysql数据库
	btn := db.GetInstance().InitPool()
	if !btn {
		log.Println("init database pool failure...")
		panic(errors.New("init database pool failure"))
	}

	// redis
	db.InitRedis()

	gin.SetMode(viper.GetString("mode"))
	g := gin.New()
	g = router.Load(g)
	err := g.Run(viper.GetString("addr"))
	if err != nil {
		return
	}
}
