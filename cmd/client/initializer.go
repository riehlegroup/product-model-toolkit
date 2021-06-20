package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	homedir "github.com/mitchellh/go-homedir"
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
	rootCmd.AddCommand(crawlerCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// adding the subcommands for the crawlerCmd
	crawlerCmd.AddCommand(listCrawlerOptions)
	crawlerCmd.Flags().StringVarP(&crawlerName, "name", "n", "", "crawler name")
	crawlerCmd.Flags().StringVarP(&crawlerSource, "source", "s", "", "crawler source")
	crawlerCmd.Flags().StringVarP(&crawlerOutput, "out", "o", "", "crawler output path")
	_ = crawlerCmd.MarkFlagRequired("name")
	_ = crawlerCmd.MarkFlagRequired("source")
	_ = crawlerCmd.MarkFlagRequired("out")

	// adding the subcommands for the diffCmd
	diffCmd.AddCommand(diffBasedOnId)
	diffCmd.AddCommand(diffBasedOnPath)

	diffBasedOnId.Flags().StringVarP(&diffFirstId, "first", "f", "", "first id")
	diffBasedOnId.Flags().StringVarP(&diffSecondId, "second", "s", "", "second id")
	_ = diffBasedOnId.MarkFlagRequired("first")
	_ = diffBasedOnId.MarkFlagRequired("second")

	diffBasedOnPath.Flags().StringVarP(&diffFirstFile, "first", "f", "", "first file")
	diffBasedOnPath.Flags().StringVarP(&diffSecondFile, "second", "s", "", "second file")
	_ = diffBasedOnPath.MarkFlagRequired("first")
	_ = diffBasedOnPath.MarkFlagRequired("second")

	// adding the subcommands for the exportCmd
	exportCmd.AddCommand(listExportOptions)
	exportCmd.Flags().StringVarP(&exportType, "type", "t", "", "export file type (required)")
	exportCmd.Flags().StringVarP(&exportPath, "path", "p", "", "export file path (required)")
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
