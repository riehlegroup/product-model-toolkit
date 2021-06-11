// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import "github.com/spf13/cobra"



// Execute adds all child commands to the
// root command and sets flags appropriately.
// It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}


func main() {
	// set-up the commands
	Execute()

	// TODO (we need to have http handlers for all of the services)
	//scanning.Run(
	//	cfg,
	//	rest.NewClient(serverBaseURL),
	//)

}
