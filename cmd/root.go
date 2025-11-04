package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

type Note struct {
	Description string
	CreatedAt string
	IsComplete bool
}

var rootCmd = &cobra.Command {
	Use: "tasks",
	Short: "love me some tasks",
	Long: "Note taking app",
}

func Execute() {
	if err:= rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
