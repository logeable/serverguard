package servermgr

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/pflag"

	"github.com/logeable/serverguard/internal/servermgr/option"

	"github.com/spf13/cobra"
)

type App struct {
	cmd    *cobra.Command
	option *option.Option
}

func NewApp() *App {
	opt := option.NewOption()
	cmd := &cobra.Command{
		Use: "servermgr",
		Run: cobraRun(opt),
	}
	addFlags(cmd, opt.Flags())

	app := &App{
		cmd:    cmd,
		option: opt,
	}
	return app
}

func addFlags(cmd *cobra.Command, flags []*pflag.FlagSet) {
	for _, fs := range flags {
		cmd.Flags()
		cmd.Flags().AddFlagSet(fs)
	}
}

func (app *App) Run() {
	err := app.cmd.Execute()
	if err != nil {
		log.Fatalf("execute command failed: %v", err)
	}
}

func cobraRun(opt *option.Option) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		run(opt)
	}
}

func run(opt *option.Option) {
	spew.Dump(opt)
}
