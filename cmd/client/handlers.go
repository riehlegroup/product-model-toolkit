// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/osrgroup/product-model-toolkit/cnst"
	"github.com/osrgroup/product-model-toolkit/pkg/server/commands"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

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

// callDiffId function for the diffCmd command -> returns error
func callDiffId(first, second string) error {

	// run the diff by id and check the error
	if err := commands.RunDiffById(first, second); err != nil {
		return err
	}

	return nil
}

// callDiffId function for the diffCmd command -> returns error
func callDiffPath(first, second string) error {

	// run the diff by path and check the error
	if err := commands.RunDiffByPath(first, second); err != nil {
		return err
	}

	return nil
}

// callExport function for the exportCmd command -> returns error
func callExport(exportId, exportType, exportPath string) error {
	// define the post path
	postPath := fmt.Sprintf("/products/export/%s", strings.ToLower(exportType))

	// run the run export and check the error
	if err := commands.RunExport(exportId, exportPath, postPath); err != nil {
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

	// run the run import and check the error
	if err := commands.RunImport(importPath, postPath); err != nil {
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
	// run the run merge command and check the error
	if err := commands.RunMerge(mergeFirstFile, mergeSecondFile, mergeOutput); err != nil {
		return err
	}
	return nil
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
