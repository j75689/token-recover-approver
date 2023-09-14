package cmd

import (
	"fmt"

	"github.com/bnb-chain/airdrop-service/internal/version"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:          "version",
		Short:        "Print version information",
		SilenceUsage: true,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(version.Version())
		},
	}
)
