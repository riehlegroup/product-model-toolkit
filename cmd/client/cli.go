// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	homedir "github.com/mitchellh/go-homedir"
	"log"
	"time"

	"github.com/osrgroup/product-model-toolkit/cnst"
	"github.com/osrgroup/product-model-toolkit/pkg/server/commands"
	"github.com/pterm/pterm"
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
	// rootCmd.AddCommand(mergeCmd)
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

// Execute adds all child commands
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}


func main() {
	Execute()
}



// introScreen creates a fancy intro message
func introScreen() {
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("PMT", pterm.NewStyle(pterm.FgLightGreen))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)

	pterm.DefaultCenter.Print(pterm.DefaultHeader.
		WithFullWidth().
		WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).
		WithMargin(10).
		Sprint(cnst.CliShort))
	pterm.Info.Println(pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()

}

// callscanner function for the scannerCmd command -> returns error
func callscanner(name, source, output string) error {
	// run the scanner and check the error
	if err := commands.RunScanner(name, source, output); err != nil {
		return err
	}

	return nil
}

// callDiffPath function for the diffCmd command -> returns error
func callDiffPath(firstPath, secondPath string) error {

	// run the diff by path and check the error
	if err := commands.RunDiffByPath(firstPath, secondPath); err != nil {
		return err
	}

	return nil
}

// callExport function for the exportCmd command -> returns error
func callExport(exportId, exportType, exportPath string) error {
	// run the run export and check the error
	if err := commands.RunExport(exportId, exportType, exportPath); err != nil {
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
		// "custom-report",
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

	// run the run import and check the error
	if err := commands.RunImport(importType, importPath); err != nil {
		return err
	}
	return nil
}

// listAvailablescannerTypes function for the scannerCmd command
func listAvailablescannerTypes() error {
	// list available scanner types and check the error
	if err := commands.ListAvailableScannerTypes(); err != nil {
		return err
	}

	return nil
}

// listAvailableImportTypes function for the importCmd command
func listAvailableImportTypes() {
	// define the available import options list
	availableImportOptions := []string{
		"spdx",
		"composer",
		"file-hasher",
		"scanner",
	}

	// print instruction, loop over the
	// list and print the available options
	fmt.Println("Available import options:")
	for key, name := range availableImportOptions {
		fmt.Printf("%v) %v\n", key+1, name)
	}
}

// callSearch function for the searchCmd command
func callSearch(searchPackageName, searchRootDir, searchFileOut string) error {
	// run the run search command and check the error
	if err := commands.RunSearch(searchPackageName, searchRootDir, searchFileOut); err != nil {
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   cnst.Cli,
	Short: cnst.CliShort,
	Long:  cnst.CliLong,
	Run: func(cmd *cobra.Command, args []string) {
		introScreen()
		cmd.Help()
	},
}

// scannerCmd
var scannerCmd = &cobra.Command{
	Use:   cnst.Scanner,
	Short: cnst.ScannerShort,
	Long:  cnst.ScannerLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callscanner(scannerName, scannerSource, scannerOutput); err != nil {
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
	Use:   cnst.Version,
	Short: cnst.VersionShort,
	Long:  cnst.VersionLong,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callVersion(); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// listscannerOptions command, this works as a subcommand for the scannerCmd command
var listScannerOptions = &cobra.Command{
	Use:   cnst.List,
	Short: cnst.ListShortScanner,
	Long:  cnst.ListLongScanner,
	Run: func(cmd *cobra.Command, args []string) {
		listAvailablescannerTypes()
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
