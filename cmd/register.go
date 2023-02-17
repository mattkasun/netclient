/*
Copyright © 2022 Netmaker Team <info@netmaker.io>
*/
package cmd

import (
	"github.com/gravitl/netclient/functions"
	"github.com/gravitl/netmaker/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register to a Netmaker instance",

	Long: `register to a Netmaker instance using: 
token: netclient register -t <token> // join using an enrollment token`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := viper.New()
		flags.BindPFlags(cmd.Flags())
		// CLI should always take password from stdin
		flags.Set("readPassFromStdIn", true)
		token := flags.GetString("token")
		if token == "" {
			cmd.Usage()
			return
		}
		if err := functions.Register(token); err != nil {
			logger.Log(0, "registration failed", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("token", "t", "", "enrollment token for registering to a Netmaker instance")
}
