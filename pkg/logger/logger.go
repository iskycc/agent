package logger

import (
	"fmt"
	"sync"
	"time"
)

var (
	defaultLogger = &ServiceLogger{enabled: true, logger: newConsoleLogger()}

	loggerOnce sync.Once
)

// consoleLogger 实现控制台日志输出
type consoleLogger struct{}

func newConsoleLogger() *consoleLogger {
	return &consoleLogger{}
}

func (c *consoleLogger) Infof(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func (c *consoleLogger) Errorf(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(msg)
	return fmt.Errorf(msg)
}

type ServiceLogger struct {
	enabled bool
	logger  *consoleLogger
}

func InitDefaultLogger(enabled bool, _ interface{}) {
	loggerOnce.Do(func() {
		defaultLogger.enabled = enabled
	})
}

func SetEnable(enable bool) {
	defaultLogger.SetEnable(enable)
}

func Println(v ...interface{}) {
	defaultLogger.Println(v...)
}

func Printf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}

func Error(v ...interface{}) error {
	return defaultLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) error {
	return defaultLogger.Errorf(format, v...)
}

func NewServiceLogger(enable bool, _ interface{}) *ServiceLogger {
	return &ServiceLogger{
		enabled: enable,
		logger:  newConsoleLogger(),
	}
}

func (s *ServiceLogger) SetEnable(enable bool) {
	s.enabled = enable
}

func (s *ServiceLogger) Println(v ...interface{}) {
	if s.enabled {
		s.logger.Infof("NEZHA@%s>> %v", time.Now().Format(time.DateTime), fmt.Sprint(v...))
	}
}

func (s *ServiceLogger) Printf(format string, v ...interface{}) {
	if s.enabled {
		s.logger.Infof("NEZHA@%s>> "+format, append([]interface{}{time.Now().Format(time.DateTime)}, v...)...)
	}
}

func (s *ServiceLogger) Error(v ...interface{}) error {
	if s.enabled {
		return s.logger.Errorf("NEZHA@%s>> %v", time.Now().Format(time.DateTime), fmt.Sprint(v...))
	}
	return nil
}

func (s *ServiceLogger) Errorf(format string, v ...interface{}) error {
	if s.enabled {
		return s.logger.Errorf("NEZHA@%s>> "+format, append([]interface{}{time.Now().Format(time.DateTime)}, v...)...)
	}
	return nil
}
