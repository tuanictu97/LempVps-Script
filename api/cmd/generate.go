package cmd

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"io.github.tuanictu97/api/pkg/logger"
	"io.github.tuanictu97/api/tools/generate"
)

var GenerateCmd = &cli.Command{
	Name:  "gen",
	Usage: "Generate module files by template",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "logcfg",
			Usage: "Logger configuration file",
			Value: "configs/logger.toml",
		},
		&cli.StringFlag{
			Name:  "projectdir",
			Usage: "Project directory",
			Value: ".",
		},
		&cli.StringFlag{
			Name:  "tpldir",
			Usage: "Template directory",
			Value: "tools/generate/tpls",
		},
		&cli.StringSliceFlag{
			Name:  "result",
			Usage: "Specify the generated result, can be multiple",
			Value: cli.NewStringSlice("api", "biz", "dao", "typed"),
		},
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"},
			Usage:    "Configuration file (yaml)",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		defer func() {
			err := zap.L().Sync()
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("Failed to sync zap logger: %v \n", err))
			}
		}()

		ctx := logger.NewTag(context.Background(), "generate")
		// Initialize the logger first
		cleanLogger, err := logger.InitWithConfig(c.String("logcfg"), HandleLoggerHook)
		if err != nil {
			return err
		}
		defer cleanLogger()

		err = generate.Generate(ctx, c.String("projectdir"), c.String("tpldir"), c.StringSlice("result"), c.String("config"))
		if err != nil {
			logger.Context(ctx).Error("Failed to generate", zap.Error(err))
		}
		return err
	},
}
