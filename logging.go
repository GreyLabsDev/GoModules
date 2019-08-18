import (
	"fmt"
	"time"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

var GlobalLogger Logger

type EventType int

const (
	LogDebug EventType = 0
	LogError EventType = 1
	LogInfo  EventType = 2
)

type Logger struct {
	StartTime string
}

func (logger *Logger) init() {
	logger.StartTime = time.Now().String()
	fmt.Println("LOGGER STARTED - " + logger.StartTime)
}

func (logger *Logger) Log(eventType EventType, message string) {
	switch eventType {
	case LogDebug:
		fmt.Printf(DebugColor, "LOG.DEBUG: "+message)
	case LogError:
		fmt.Printf(ErrorColor, "LOG.ERROR: "+message)
	case LogInfo:
		fmt.Printf(InfoColor, "LOG.INFO: "+message)
	}
	fmt.Println()
}

func InitGlobalLogger() {
	GlobalLogger = Logger{}
	GlobalLogger.init()
}
