// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/osrgroup/product-model-toolkit/cnst"
	"github.com/osrgroup/product-model-toolkit/pkg/client/commands"
	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	// config file
	cfgFile string

	// crawlerCmd
	crawlerName   string
	crawlerOutput string
	crawlerSource string

	// diffCmd
	diffFirstId    string
	diffSecondId   string
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

// base url for connecting to server
const serverBaseURL = cnst.ServerBaseURL

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   cnst.Cli,
	Short: cnst.CliShort,
	Long:  cnst.CliLong,
}

// crawlerCmd
var crawlerCmd = &cobra.Command{
	Use:   cnst.Crawler,
	Short: cnst.CrawlerShort,
	Long:  cnst.CrawlerLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callCrawler(crawlerName, crawlerSource, crawlerOutput); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// diffCmd
var diffCmd = &cobra.Command{
	Use:   cnst.Diff,
	Short: cnst.DiffShort,
	Long:  cnst.DiffLong,
}

// diffBasedOnId command, this works as a subcommand for the diffCmd command
var diffBasedOnId = &cobra.Command{
	Use:   cnst.Id,
	Short: cnst.IdShort,
	Long:  cnst.IdLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callDiffId(diffFirstId, diffSecondId); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// diffBasedOnPath command, this works as a subcommand for the diffCmd command
var diffBasedOnPath = &cobra.Command{
	Use:   cnst.Path,
	Short: cnst.PathShort,
	Long:  cnst.PathLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callDiffPath(diffFirstFile, diffSecondFile); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// exportCmd
var exportCmd = &cobra.Command{
	Use:   cnst.Export,
	Short: cnst.ExportShort,
	Long:  cnst.ExportLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callExport(exportId, exportType, exportPath); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// importCmd
var importCmd = &cobra.Command{
	Use:   cnst.Import,
	Short: cnst.ImportShort,
	Long:  cnst.ImportLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callImport(importType, importPath); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// mergeCmd
var mergeCmd = &cobra.Command{
	Use:   cnst.Merge,
	Short: cnst.MergeShort,
	Long:  cnst.MergeLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callMerge(mergeFirstFile, mergeSecondFile, mergeOutput); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// searchCmd
var searchCmd = &cobra.Command{
	Use:   cnst.Search,
	Short: cnst.SearchShort,
	Long:  cnst.SearchLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callSearch(searchPackageName, searchRootDir, searchFileOut); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// versionCmd
var versionCmd = &cobra.Command{
	Use: cnst.Version,
	Short: cnst.VersionShort,
	Long: cnst.VersionLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callVersion(); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// listCrawlerOptions command, this works as a subcommand for the crawlerCmd command
var listCrawlerOptions = &cobra.Command{
	Use:   cnst.List,
	Short: cnst.ListShortCrawler,
	Long:  cnst.ListLongCrawler,
	Run: func(cmd *cobra.Command, args []string) {
		listAvailableCrawlerTypes()
	},
}

// listImportOptions command, this works as a subcommand for the importCmd command
var listImportOptions = &cobra.Command{
	Use:   cnst.List,
	Short: cnst.ListShortImport,
	Long:  cnst.ListLongImport,
	Run: func(cmd *cobra.Command, args []string) {
		listAvailableImportTypes()
	},
}

// listExportOptions command, this works as a subcommand for the exportCmd command
var listExportOptions = &cobra.Command{
	Use:   cnst.List,
	Short: cnst.ListShortExport,
	Long:  cnst.ListLongExport,
	Run: func(cmd *cobra.Command, args []string) {
		listAvailableExportTypes()
	},
}

// Execute adds all child commands to the
// root command and sets flags appropriately.
// It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

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

func main() {
	// set-up the commands
	Execute()

	// TODO (we need to have http handlers for all of the services)
	//scanning.Run(
	//	cfg,
	//	rest.NewClient(serverBaseURL),
	//)

}

// callCrawler function for the crawlerCmd command -> returns error
func callCrawler(name, source, output string) error {
	// creating a new http client
	client := rest.NewClient(serverBaseURL)

	// run the crawler and check the error
	if err := commands.RunCrawler(name, source, output, client); err != nil {
		return err
	}

	return nil
}

// callDiffId function for the diffCmd command -> returns error
func callDiffId(first, second string) error {
	// creating a new http client
	client := rest.NewClient(serverBaseURL)

	// run the diff by id and check the error
	if err := commands.RunDiffById(first, second, client); err != nil {
		return err
	}

	return nil
}

// callDiffId function for the diffCmd command -> returns error
func callDiffPath(first, second string) error {
	// creating a new http client
	client := rest.NewClient(serverBaseURL)

	// run the diff by path and check the error
	if err := commands.RunDiffByPath(first, second, client); err != nil {
		return err
	}

	return nil
}

// callExport function for the exportCmd command -> returns error
func callExport(exportId, exportType, exportPath string) error {
	// define the post path
	postPath := fmt.Sprintf("/products/export/%s", strings.ToLower(exportType))

	// creating a new http client
	client := rest.NewClient(serverBaseURL)

	// run the run export and check the error
	if err := commands.RunExport(exportId, exportPath, postPath, client); err != nil {
		return err
	}
	return nil
}

// listAvailableExportTypes function for the exportCmd command
func listAvailableExportTypes() {
	// define the available export options list
	availableExportOptions := []string{
		"spdx",
		"human-read",
		"custom-report",
	}

	// print instrcution, loop over the
	// list and print the available options
	fmt.Println("Available import options:")
	for key, name := range availableExportOptions {
		fmt.Printf("%v) %v\n", key+1, name)
	}
}

// callImport function for the importCmd command -> returns error
func callImport(importType, importPath string) error {
	// define the post path
	postPath := fmt.Sprintf("/products/import/%s", strings.ToLower(importType))

	// creating a new http client
	client := rest.NewClient(serverBaseURL)

	// run the run import and check the error
	if err := commands.RunImport(importPath, postPath, client); err != nil {
		return err
	}
	return nil
}

// listAvailableCrawlerTypes function for the crawlerCmd command
func listAvailableCrawlerTypes() {
	// define the available crawler options
	availableCrawlerOptions := []string{
		"php-scanner",
	}

	// print instruction, loop over the
	// list and print the available options
	fmt.Println("Available crawler options:")
	for key, name := range availableCrawlerOptions {
		fmt.Printf("%v) %v\n", key+1, name)
	}
}

// listAvailableImportTypes function for the importCmd command
func listAvailableImportTypes() {
	// define the available import options list
	availableImportOptions := []string{
		"spdx",
		"composer",
		"file-hasher",
	}

	// print instruction, loop over the
	// list and print the available options
	fmt.Println("Available import options:")
	for key, name := range availableImportOptions {
		fmt.Printf("%v) %v\n", key+1, name)
	}
}

// callMerge function for the mergeCmd command -> returns error
func callMerge(mergeFirstFile, mergeSecondFile, mergeOutput string) error {
	// creating a new http client
	client := rest.NewClient(serverBaseURL)

	// run the run merge command and check the error
	if err := commands.RunMerge(mergeFirstFile, mergeSecondFile, mergeOutput, client); err != nil {
		return err
	}
	return nil
}

// callSearch function for the searchCmd command
func callSearch(searchPackageName, searchRootDir, searchFileOut string) error {
	// creating a new http client
	client := rest.NewClient(serverBaseURL)

	// run the run search command and check the error
	if err := commands.RunSearch(searchPackageName, searchRootDir, searchFileOut, client); err != nil {
		return err
	}
	return nil
}

// printVersion function for the versionCmd command
func callVersion() error {
	// run the run version command and check the error
	if err := commands.RunVersion(gitCommit); err != nil {
		return err
	}
	return nil
}
