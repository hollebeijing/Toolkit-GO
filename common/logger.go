package common

import (
	"flag"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

/***
path:日志路径
MaxSize:文件大小限制,单位MB
MaxBackups:最大保留日志文件数量
MaxAge:日志文件保留天数
*/
/***
使用方法

go.mod 的项目导入
	1、在go.mod的require中添加: naas/sd-toolkit-go v0.0.0-incompatible
	2、在go.mod的replace中添加: naas/sd-toolkit-go v0.0.0-incompatible => ../SD-Toolkit-GO

初始化：
	在项目main文件中，
		toolkit "naas/sd-toolkit-go/src/common"
		init(){
			_ = toolkit.InitLogger("./logs/", 50, 2, 10)
		}

使用：
	toolkit "naas/sd-toolkit-go/src/common"
	toolkit.Logger.Info("xxxx", zap.String("version", version.VERSION))
*/

func init() {
	flag.BoolVar(&logging.StdOut, "stdout", false, "logs print to Stdout")
	flag.BoolVar(&logging.InfoOption.Enable, "log_file_enable", true, "logs to standard info as well as files")
	flag.StringVar(&logging.InfoOption.Path, "log_path", "./logs", "logs path")
	flag.StringVar(&logging.InfoOption.FileName, "log_file_name", "info.log", "log file name")
	flag.IntVar(&logging.InfoOption.MaxSize, "log_max_size", 50, "log file max size(MB)")
	flag.IntVar(&logging.InfoOption.MaxBackups, "log_max_backups", 10, "log file max retain count ")
	flag.IntVar(&logging.InfoOption.MaxAge, "log_max_age", 30, "log file max retain days")
	flag.BoolVar(&logging.ErrorOption.Enable, "error_file_enable", true, "logs to standard info as well as files")
	flag.StringVar(&logging.ErrorOption.Path, "error_path", "./logs", "logs path")
	flag.StringVar(&logging.ErrorOption.FileName, "error_file_name", "error.log", "log file name")
	flag.IntVar(&logging.ErrorOption.MaxSize, "error_max_size", 50, "log file max size(MB)")
	flag.IntVar(&logging.ErrorOption.MaxBackups, "error_max_backups", 10, "log file max retain count ")
	flag.IntVar(&logging.ErrorOption.MaxAge, "error_max_age", 30, "log file max retain days")

}

var logging LoggerOption

func InitLogger() {
	var coreArr []zapcore.Core

	//NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	//获取编码器
	var encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//日志解蔽
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件writeSyncer
	infoOption := logging.InfoOption
	if infoOption.Enable {
		infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s", infoOption.Path, infoOption.FileName), //日志文件存放目录，如果文件夹不存在会自动创建
			MaxSize:    infoOption.MaxSize,                                         //文件大小限制,单位MB
			MaxBackups: infoOption.MaxBackups,                                      //最大保留日志文件数量
			MaxAge:     infoOption.MaxAge,                                          //日志文件保留天数
			Compress:   false,                                                      //是否压缩处理
		})
		infoFileCore := zapcore.NewCore(encoder, zapcore.AddSync(infoFileWriteSyncer), lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
		coreArr = append(coreArr, infoFileCore)
	}

	//error文件writeSyncer
	errorOption := logging.ErrorOption
	if errorOption.Enable {
		errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s", errorOption.Path, errorOption.FileName), //日志文件存放目录
			MaxSize:    errorOption.MaxSize,                                          //文件大小限制,单位MB
			MaxBackups: errorOption.MaxBackups,                                       //最大保留日志文件数量
			MaxAge:     errorOption.MaxAge,                                           //日志文件保留天数
			Compress:   false,                                                        //是否压缩处理
		})
		errorFileCore := zapcore.NewCore(encoder, zapcore.AddSync(errorFileWriteSyncer), highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
		coreArr = append(coreArr, errorFileCore)
	}

	if logging.StdOut {
		stdCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lowPriority)
		coreArr = append(coreArr, stdCore)
	}

	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
}

type LoggerOption struct {
	StdOut     bool
	InfoOption struct {
		Enable     bool
		Path       string
		FileName   string
		MaxSize    int
		MaxBackups int
		MaxAge     int
	}

	ErrorOption struct {
		Enable     bool
		Path       string
		FileName   string
		MaxSize    int
		MaxBackups int
		MaxAge     int
	}
}

func InitLoggerWithOption(option *LoggerOption) {

	var coreArr []zapcore.Core

	//NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	//获取编码器
	var encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//日志解蔽
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件writeSyncer
	infoOption := option.InfoOption
	if infoOption.Enable {
		infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s", infoOption.Path, infoOption.FileName), //日志文件存放目录，如果文件夹不存在会自动创建
			MaxSize:    infoOption.MaxSize,                                         //文件大小限制,单位MB
			MaxBackups: infoOption.MaxBackups,                                      //最大保留日志文件数量
			MaxAge:     infoOption.MaxAge,                                          //日志文件保留天数
			Compress:   false,                                                      //是否压缩处理
		})
		infoFileCore := zapcore.NewCore(encoder, zapcore.AddSync(infoFileWriteSyncer), lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
		coreArr = append(coreArr, infoFileCore)
	}

	//error文件writeSyncer
	errorOption := option.ErrorOption
	if errorOption.Enable {
		errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s", errorOption.Path, errorOption.FileName), //日志文件存放目录
			MaxSize:    errorOption.MaxSize,                                          //文件大小限制,单位MB
			MaxBackups: errorOption.MaxBackups,                                       //最大保留日志文件数量
			MaxAge:     errorOption.MaxAge,                                           //日志文件保留天数
			Compress:   false,                                                        //是否压缩处理
		})
		errorFileCore := zapcore.NewCore(encoder, zapcore.AddSync(errorFileWriteSyncer), highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
		coreArr = append(coreArr, errorFileCore)
	}

	if option.StdOut {
		stdCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(os.Stdout, os.Stderr), lowPriority)
		coreArr = append(coreArr, stdCore)
	}

	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
}

func GetLogger() *zap.Logger {
	return Logger
}
