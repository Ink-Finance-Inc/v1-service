// main.go
package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"inkfinance.xyz/chain"
	"inkfinance.xyz/conf"

	// doc "inkfinance.xyz/docs"
	http "inkfinance.xyz/server"
)

// @title           Inkfinance db API
// @version         1.0

// @contact.name   InkFinance.xyz.TECH
// @contact.url    http://www.swagger.io/support
// @contact.email  support@inkfinance.xyz

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func init() {
	conf.DB = conf.ConnectDB(conf.MySQLDSN).Debug()

}

func initConfig(settingFile string) {

	viper.SetConfigName(settingFile) // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")         // optionally look for config in the working directory
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	fmt.Println("init db connection:", viper.GetString("mysql.connection"))
	conf.MySQLDSN = viper.GetString("mysql.connection")
	conf.DB = conf.ConnectDB(conf.MySQLDSN).Debug()

	fmt.Println("init finished")

}

func main() {

	var config string
	if len(os.Args) > 1 {
		config = os.Args[1]
	}

	fmt.Println(("config: " + config))
	if config == "" || len(config) == 0 {
		config = "config-local"
	}

	fmt.Println(config)
	initConfig(config)
	errc := make(chan error)
	var url = "0.0.0.0:" + viper.GetString("port")
	go http.Run(url, errc)
	// 等grpc服务完成后，在这里启动 grpc
	fmt.Println("started at: ", url)

	// engine := gin.New()
	// engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// startChainScan()
	log.WithField("error", <-errc).Info("Exit")

}

func startChainScan() {
	chain.StartChainScan()
}
