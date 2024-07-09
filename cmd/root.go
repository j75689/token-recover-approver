package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/bnb-chain/token-recover-app/internal/app"
	"github.com/bnb-chain/token-recover-app/pkg/util"
)

// Root command
var (
	rootCmd = &cobra.Command{
		Run: func(_ *cobra.Command, _ []string) {
			application, err := app.Initialize(cfgFile)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			startWithModules := func() error {
				appModules := []app.Modules{}
				for _, module := range strings.Split(modules, ",") {
					appModules = append(appModules, app.Modules(module))
				}
				return application.Start(appModules)
			}
			util.Launch(startWithModules, application.Stop, time.Duration(timeout)*time.Second)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}

// Flags
var (
	timeout uint
	cfgFile string
	modules string
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(toolCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().UintVar(&timeout, "timeout", 300, "graceful shutdown timeout (second)")
	rootCmd.PersistentFlags().StringVar(&modules, "modules", app.APIModule.String(), "modules to start, separated by comma(api,tracker,bot)")
}
