package option

import "github.com/spf13/pflag"

type Option struct {
	GRPCOption *GRPCOption
	LogOption  *LogOption
}

func NewOption() *Option {
	return &Option{
		GRPCOption: NewGRPCOption(),
		LogOption:  NewLogOption(),
	}
}

func (o *Option) Flags() []*pflag.FlagSet {
	var ret []*pflag.FlagSet
	ret = append(ret,
		o.GRPCOption.Flags(),
		o.LogOption.Flags(),
	)
	return ret
}
