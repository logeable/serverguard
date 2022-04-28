package option

import "github.com/spf13/pflag"

type GRPCOption struct {
	BindIP   string `json:"bind_ip" mapstructure:"bind_ip"`
	BindPort int    `json:"bind_port" mapstructure:"bind_port"`
}

func NewGRPCOption() *GRPCOption {
	return &GRPCOption{
		BindIP:   "0.0.0.0",
		BindPort: 7420,
	}
}

func (o *GRPCOption) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("grpc", pflag.ExitOnError)
	fs.StringVar(&o.BindIP, "grpc.bind-ip", o.BindIP, "grpc bind ip")
	fs.IntVar(&o.BindPort, "grpc.bind-port", o.BindPort, "grpc bind port")
	return fs
}
