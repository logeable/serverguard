package option

import (
	"github.com/spf13/pflag"
	"go.uber.org/zap/zapcore"
)

type LogOption struct {
	Level string `json:"level" mapstructure:"level"`
}

func NewLogOption() *LogOption {
	return &LogOption{
		Level: zapcore.InfoLevel.String(),
	}
}

func (o *LogOption) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("log", pflag.ExitOnError)

	fs.StringVar(&o.Level, "log.level", o.Level, "log level")
	return fs
}
