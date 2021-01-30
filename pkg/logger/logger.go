package logger

import (
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/sirupsen/logrus"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

type Logger struct {
	log    *logrus.Logger
	fields logrus.Fields

	addr  string
	level string
}

func New(fields map[string]interface{}, addr, level string, debug bool) *Logger {
	log := logrus.New()

	if !debug {
		log.SetFormatter(&GelfFormatter{fields: fields})
	}

	return &Logger{
		log:    log,
		fields: fields,
		addr:   addr,
		level:  level,
	}
}

func (l *Logger) Connect() error {
	lvl, err := logrus.ParseLevel(l.level)
	if err != nil {
		return err
	}

	l.log.SetLevel(lvl)

	hook := graylog.NewGraylogHook(l.addr, l.fields)

	l.log.Hooks.Add(hook)

	return nil
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log.WithFields(l.fields).Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.log.WithFields(l.fields).Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log.WithFields(l.fields).Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log.WithFields(l.fields).Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log.WithFields(l.fields).Fatalf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.log.WithFields(l.fields).Panicf(format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.log.WithFields(l.fields).Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.log.WithFields(l.fields).Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log.WithFields(l.fields).Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.log.WithFields(l.fields).Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.log.WithFields(l.fields).Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.log.WithFields(l.fields).Panic(args...)
}
