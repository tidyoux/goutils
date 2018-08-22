package goutils

import (
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

// InitRotationLogger inits rotation log config.
func InitRotationLogger(filePath string, maxAge, rotationTime time.Duration, formatter log.Formatter) error {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	writer, err := rotatelogs.New(
		filePath+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(filePath),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		return err
	}

	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			log.DebugLevel: writer,
			log.InfoLevel:  writer,
			log.WarnLevel:  writer,
			log.ErrorLevel: writer,
			log.FatalLevel: writer,
			log.PanicLevel: writer,
		},
		formatter,
	))
	return nil
}
