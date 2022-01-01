package root

import (
	"github.com/spf13/cobra"
	"ws/cmd/fake"
	"ws/cmd/migrate"
	"ws/cmd/serve"
	"ws/config"
)

func NewRootCommand(name string) *cobra.Command {

	var configFile string

	var rootCmd = &cobra.Command{
		Use:                        name,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
		CompletionOptions:          cobra.CompletionOptions{},
		TraverseChildren: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config.Setup(configFile)
		},
	}

	flag := rootCmd.PersistentFlags()
	flag.StringVar(&configFile, "config", "config.yaml", "config file")


	rootCmd.AddCommand(serve.NewServeCommand(),
		migrate.NewMigrateCommand(),
		fake.NewFakeCommand(),
	)

	return rootCmd
}
