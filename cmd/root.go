package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/automaxprocs/maxprocs"
	"link/internal/conf"
	"link/pkg/utils"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:   "link",
		Short: `a community`,
		Long:  `a community`,
		Run:   runServe,
	}
)

func runServe(*cobra.Command, []string) {
	utils.PrintBanner("LINKING...")
	_, _ = maxprocs.Set(maxprocs.Logger(log.Printf))

	conf.Initial()
}

func Execute() {
	_ = rootCmd.Execute()
}
