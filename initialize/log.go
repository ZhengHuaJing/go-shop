package initialize

import (
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/pkg/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

func Log() {
	encoder := getEncoder()
	logFile := getLogFile()
	core := zapcore.NewCore(encoder, logFile, zapcore.DebugLevel)
	global.Log = zap.New(core, zap.AddCaller()).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogFile() *os.File {
	logCfg := global.Config.Log
	fileCfg := global.Config.File
	timeCfg := global.Config.Time
	fileName := logCfg.LogSaveName + time.Now().Format(timeCfg.DateFormat) + "." + logCfg.LogFileExt
	filePath := fileCfg.RuntimeRootPath + logCfg.LogSavePath
	logFile, err := file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatal("logging.Setup err: " + err.Error())
	}

	return logFile
}
