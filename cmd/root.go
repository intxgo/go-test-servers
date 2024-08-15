package cmd

import (
	"log"
	"os"
	"os/signal"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-test-server",
		Short: "A command line tool for running test servers",
	}
)

func Execute() error {


	return rootCmd.Execute()
}