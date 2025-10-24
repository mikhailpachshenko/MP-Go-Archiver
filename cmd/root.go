package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "simple archiver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleErr(err)
		os.Exit(1)
	}
}

func handleErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
}
