package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	homedir "github.com/mitchellh/go-homedir"
)

// define required variables
var (
	// config file
	cfgFile string

	// crawlerCmd
	scannerName   string
	scannerOutput string
	scannerSource string

	// diffCmd
	diffFirstFile  string
	diffSecondFile string

	// exportCmd
	exportId   string
	exportType string
	exportPath string

	// importCmd
	importType string
	importPath string

	// mergeCmd
	mergeFirstFile  string
	mergeSecondFile string
	mergeOutput     string

	// searchCmd
	searchPackageName string
	searchRootDir     string
	searchFileOut     string

	// git commit version
	gitCommit string
)

// init function
func init() {

	// initializing the cobra cli application
	cobra.OnInitialize(initConfig)

	// adding the subcommands for the rootCmd
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(diffCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(scannerCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// adding the subcommands for the crawlerCmd
	scannerCmd.AddCommand(listScannerOptions)
	scannerCmd.Flags().StringVarP(&scannerName, "name", "n", "", "scanner name")
	scannerCmd.Flags().StringVarP(&scannerSource, "source", "s", "", "scanner source")
	scannerCmd.Flags().StringVarP(&scannerOutput, "out", "o", "", "scanner output path")
	_ = scannerCmd.MarkFlagRequired("name")
	_ = scannerCmd.MarkFlagRequired("source")
	_ = scannerCmd.MarkFlagRequired("out")

	// adding the subcommands for the diffCmd
	diffCmd.AddCommand(diffBasedOnPath)

	diffBasedOnPath.Flags().StringVarP(&diffFirstFile, "first", "f", "", "first file")
	diffBasedOnPath.Flags().StringVarP(&diffSecondFile, "second", "s", "", "second file")
	_ = diffBasedOnPath.MarkFlagRequired("first")
	_ = diffBasedOnPath.MarkFlagRequired("second")

	// adding the subcommands for the exportCmd
	exportCmd.AddCommand(listExportOptions)
	exportCmd.Flags().StringVarP(&exportId, "id", "i", "", "id of the product to be exported (required)")
	exportCmd.Flags().StringVarP(&exportType, "type", "t", "", "export type (required)")
	exportCmd.Flags().StringVarP(&exportPath, "path", "p", "", "export file path (required)")
	_ = exportCmd.MarkFlagRequired("id")
	_ = exportCmd.MarkFlagRequired("type")
	_ = exportCmd.MarkFlagRequired("path")


	// adding the subcommands for the importCmd
	importCmd.AddCommand(listImportOptions)
	importCmd.Flags().StringVarP(&importType, "type", "t", "", "import file type (required)")
	importCmd.Flags().StringVarP(&importPath, "path", "p", "", "import file path (required)")
	_ = importCmd.MarkFlagRequired("type")
	_ = importCmd.MarkFlagRequired("path")

	// adding the subcommands for the searchCmd
	searchCmd.Flags().StringVarP(&searchPackageName, "name", "n", "", "package name")
	searchCmd.Flags().StringVarP(&searchRootDir, "dir", "d", "", "package root dir")
	searchCmd.Flags().StringVarP(&searchFileOut, "out", "o", "", "spdx file out")
	_ = searchCmd.MarkFlagRequired("name")
	_ = searchCmd.MarkFlagRequired("dir")
	_ = searchCmd.MarkFlagRequired("out")

	// adding the subcommands for the mergeCmd
	mergeCmd.Flags().StringVarP(&mergeFirstFile, "first", "f", "", "first file")
	mergeCmd.Flags().StringVarP(&mergeSecondFile, "second", "s", "", "second file")
	mergeCmd.Flags().StringVarP(&mergeOutput, "out", "o", "", "output pat")
	_ = mergeCmd.MarkFlagRequired("first")
	_ = mergeCmd.MarkFlagRequired("second")
	_ = mergeCmd.MarkFlagRequired("out")

	// enf of the commands
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
