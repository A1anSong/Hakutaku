package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"server/config"
	"server/handler"
	"server/hook"
	_ "server/model"
	"server/util"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	addr = flag.String("s", ":8000", "server addr")
)

func main() {
	flag.Parse()
	if addr == nil {
		flag.PrintDefaults()
		return
	}

	hook.OnEvent(hook.InitHook, nil)

	viper.SetConfigName("server")
	viper.AddConfigPath("./")
	viper.AddConfigPath(util.GetExecDir())
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = config.LoadConfig(viper.GetViper())
	if err != nil {
		log.Fatal(err.Error())
	}
	initLog()
	err = hook.OnEvent(hook.LoadConfigHook, config.Configs)
	if err != nil {
		log.Fatal(err.Error())
	}

	g := gin.New()
	g.Use(gin.Recovery(), logrusLogger())
	r := g.Group("/api/v1/")

	handler.InitRouter(r)
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	sc := make(chan os.Signal)
	// signal.Notify(sc)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		sig := <-sc
		l.Close()
		log.Infof("exit: signal=<%d>.", sig)
		if sig != nil {
			log.Infof("exit: bye :-(.")
		}
	}()
	err = g.RunListener(l)
	err = hook.OnEvent(hook.CloseHook, err)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func logrusLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		log.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func initLog() {
	lv, err := log.ParseLevel(config.Configs.Log.Level)
	if err != nil {
		log.Errorf("parse log level failed, use info level,err: %s", err.Error())
		return
	}
	log.SetLevel(lv)
	log.SetReportCaller(true)
	fileName := config.Configs.Log.File
	logF, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal("open log file %s error:", fileName, err.Error())
	}
	log.SetOutput(logF)
	hook.Register(hook.CloseHook, "log_close", func(v interface{}) error {
		return logF.Close()
	})
}
