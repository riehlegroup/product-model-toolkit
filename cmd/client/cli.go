// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import "github.com/spf13/cobra"


// Execute adds all child commands
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}


func main() {
	Execute()
}
