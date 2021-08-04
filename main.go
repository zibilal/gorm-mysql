package main

import (
	"github.com/spf13/cobra"
	"gorm-mysql/app"
)

func main() {
	var rootCmd = &cobra.Command{Use: "gorm-mysql"}
	rootCmd.AddCommand(app.CmdRunApi)
	_ = rootCmd.Execute()
}
