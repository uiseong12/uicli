/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"uicli/amazon/ec2"

	"github.com/spf13/cobra"
	uicliEc2 "github.com/uiseong12/uicli/amazon/ec2"
)

var ec2LsFlag = uicliEc2.LsFlag

// ec2lsCmd represents the ec2ls command
var ec2lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ec2.Ls(&ec2LsFlag)
	},
}

func init() {
	ec2Cmd.AddCommand(ec2lsCmd)
	ec2lsCmd.Flags().BoolVarP(&ec2LsFlag.All, "all", "a", false, "List all instances")
}
