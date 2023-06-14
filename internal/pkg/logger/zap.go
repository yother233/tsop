package logger

import (
	"strings"

	"go.uber.org/zap"

	"tsop/internal/pkg/conf"
)

type Zap struct {
	log *zap.SugaredLogger
}

// NewZap create zap
func NewZap(config *conf.Config) (Logger, func(), error) {
	opts := []zap.Option{zap.AddCallerSkip(1)}
	l, err := zap.NewProduction(opts...)
	if err != nil {
		return nil, nil, err
	}

	if strings.ToLower(config.Mode) == conf.Debug {
		l, err = zap.NewDevelopment(opts...)
		if err != nil {
			return nil, nil, err
		}
	}

	return &Zap{
		log: l.Sugar(),
	}, func() { l.Sync() }, nil
}

func (z *Zap) Info(args ...interface{}) {
	z.log.Info(args)
}

// Debug implements Logger.
func (z *Zap) Debug(args ...interface{}) {
	z.log.Debug(args...)
}

// Debugf implements Logger.
func (z *Zap) Debugf(template string, args ...interface{}) {
	z.log.Debugf(template, args...)
}

// Error implements Logger.
func (z *Zap) Error(args ...interface{}) {
	z.log.Error(args...)
}

// Errorf implements Logger.
func (z *Zap) Errorf(template string, args ...interface{}) {
	z.log.Errorf(template, args...)
}

// Infof implements Logger.
func (z *Zap) Infof(template string, args ...interface{}) {
	z.log.Infof(template, args...)
}

// Panic implements Logger.
func (z *Zap) Panic(args ...interface{}) {
	z.log.Panic(args...)
}

// Panicf implements Logger.
func (z *Zap) Panicf(template string, args ...interface{}) {
	z.log.Panicf(template, args...)
}

// Warn implements Logger.
func (z *Zap) Warn(args ...interface{}) {
	z.log.Warn(args...)
}

// Warnf implements Logger.
func (z *Zap) Warnf(template string, args ...interface{}) {
	z.log.Warnf(template, args...)
}
