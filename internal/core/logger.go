package core

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/natefinch/lumberjack"
)

func (c *core) loggerInit() {
	ljack := &lumberjack.Logger{
		Filename:   "/home/robot/logs/aprialgatto.log",
		MaxSize:    2, // megabytes
		MaxBackups: 20,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	mWriter := io.MultiWriter(os.Stderr, ljack)
	log.SetOutput(mWriter)

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)
}
