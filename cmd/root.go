package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"go.uber.org/automaxprocs/maxprocs"
	"link/internal/conf"
	"link/pkg/utils"
	"link/service"
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
	services := service.InitService()
	if len(services) < 1 {
		_, _ = fmt.Fprintln(color.Output, "No service need start, exit")
		return
	}
}

func Execute() {
	_ = rootCmd.Execute()
}
