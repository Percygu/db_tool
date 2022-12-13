package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	rlog "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

// logFormatter 日志数据格式化
type logFormatter struct {
	log.TextFormatter
}

// Format 自定义日志输出格式
func (c *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	prettyCaller := func(frame *runtime.Frame) string {
		_, fileName := filepath.Split(frame.File)
		return fmt.Sprintf("%s:%d", fileName, frame.Line)
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(fmt.Sprintf("[%s] %s", entry.Time.Format(c.TimestampFormat), // 输出日志时间
		strings.ToUpper(entry.Level.String())))
	if entry.HasCaller() {
		b.WriteString(fmt.Sprintf("[%s]", prettyCaller(entry.Caller))) // 输出日志所在文件，行数位置
	}
	b.WriteString(fmt.Sprintf(" %s\n", entry.Message)) // 输出日志内容
	return b.Bytes(), nil
}

// InitConfig 初始化日志
func InitConfig() {
	globalConf := GetGlobalConf()
	// 设置日志级别
	fmt.Printf("globalConf = %v\n", globalConf)
	level, err := log.ParseLevel(globalConf.LogConf.Level)
	if err != nil {
		panic("log level parse err:" + err.Error())
	}
	// 设置日志格式为json格式
	log.SetFormatter(&logFormatter{
		log.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		}})
	log.SetReportCaller(true) // 打印文件位置，行号
	log.SetLevel(level)

	switch globalConf.LogConf.LogPattern {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	case "file":
		logger, err := rlog.New(
			globalConf.LogConf.LogPath+".%Y%m%d",
			//rlog.WithLinkName(globalConf.LogConf.LogPath),
			rlog.WithRotationCount(globalConf.LogConf.SaveDays),
			//rlog.WithMaxAge(time.Minute*3),
			rlog.WithRotationTime(time.Hour*24),
		)
		if err != nil {
			panic("log conf err: " + err.Error())
		}
		log.SetOutput(logger)
	default:
		panic("log conf err, check log_pattern in logsvr.yaml")
	}
	log.Infof("app Conf %#v", globalConf.LogConf)
}
