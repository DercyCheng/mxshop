// @title MXShop User API
// @version 1.0
// @description 电商系统用户服务API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.mxshop.com/support
// @contact.email support@mxshop.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8021
// @BasePath /u/v1
package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"web_api/user_web/global"
	"web_api/user_web/initialize"
	"web_api/user_web/utils"
	"web_api/user_web/utils/register/consul"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	myvalidator "web_api/user_web/validator"
)

func main() {
	//1. 初始化logger
	initialize.InitLogger()
	//2. 初始化配置文件
	initialize.InitConfig()
	//3. 初始化routers
	Router := initialize.Routers()
	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	//5. 初始化srv的连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	flag := viper.GetInt("DEV_CONFIG")
	if flag > 4 { // 1=zsz  2=comp  3=home
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	//服务注册
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
			debug最低，fetal最高，如果配置成info，所有比info低的都不会输出
			NewProduction默认日志级别为info
			NewDevelopment默认日志级别为debug
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功:")
	}
}
