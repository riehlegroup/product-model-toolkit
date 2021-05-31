// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/client/commands"
	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	//"github.com/osrgroup/product-model-toolkit/pkg/client/scanning"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	//"github.com/osrgroup/product-model-toolkit/pkg/client/importing"
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
const serverBaseURL = "http://localhost:8081/api/v1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Product Model Toolkit",
	Long:  `Product Model Toolkit for Managing Open Source Dependencies in Products`,
}

// crawlerCmd
var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Crawl the licenses",
	Long:  `Crawl the licenses`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callCrawler(crawlerName, crawlerSource, crawlerOutput); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// diffCmd
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Difference between two component graphs",
	Long:  `Difference between two component graphs`,
}

var diffBasedOnId = &cobra.Command{
	Use:   "id",
	Short: "Difference based on id",
	Long:  `Difference based on the id of saved products`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callDiffId(diffFirstId, diffSecondId); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

var diffBasedOnPath = &cobra.Command{
	Use:   "path",
	Short: "Difference based on path",
	Long:  `Difference based on the path of spdx files`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callDiffPath(diffFirstFile, diffSecondFile); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// exportCmd
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the component graph",
	Long:  `Export the component graph from spdx, composer or file-hasher`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callExport(exportId, exportType, exportPath); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// importCmd
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import the component graph",
	Long:  `Import the component graph from spdx, composer or file-hasher`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callImport(importType, importPath); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// mergeCmd
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge two components",
	Long:  `Merge two components`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callMerge(mergeFirstFile, mergeSecondFile, mergeOutput); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// searchCmd
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for components",
	Long:  `Search for components by their name and meta-data`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := callSearch(searchPackageName, searchRootDir, searchFileOut); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

// versionCmd
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of Product Model Toolkit",
	Long:  "This command will show the current using version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		if err := callVersion(); err != nil {
			log.Fatalln(err)
			return
		}
	},
}

var listImportOptions = &cobra.Command{
	Use:   "list",
	Short: "List all available import types",
	Long:  `List all available file types for importing as a product into the PMt`,
	Run: func(cmd *cobra.Command, args []string) {
		listAvailableImportTypes()
	},
}

var listExportOptions = &cobra.Command{
	Use:   "list",
	Short: "List all available export types",
	Long:  `List all available file types for exporting from BoM artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		listAvailableExportTypes()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// init
	cobra.OnInitialize(initConfig)

	// rootCmd
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(diffCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(crawlerCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// crawlerCmd
	crawlerCmd.Flags().StringVarP(&crawlerName, "name", "n", "", "crawler name")
	crawlerCmd.Flags().StringVarP(&crawlerSource, "source", "s", "", "crawler source")
	crawlerCmd.Flags().StringVarP(&crawlerOutput, "out", "o", "", "crawler output path")
	_ = crawlerCmd.MarkFlagRequired("name")
	_ = crawlerCmd.MarkFlagRequired("source")
	_ = crawlerCmd.MarkFlagRequired("out")

	// diffCmd
	diffCmd.AddCommand(diffBasedOnId)
	diffCmd.AddCommand(diffBasedOnPath)
	diffBasedOnId.Flags().StringVarP(&diffFirstId, "first", "f", "", "first id")
	diffBasedOnId.Flags().StringVarP(&diffSecondId, "second", "s", "", "second id")
	_ = diffBasedOnId.MarkFlagRequired("first")
	_ = diffBasedOnId.MarkFlagRequired("second")

	// exportCmd
	exportCmd.AddCommand(listExportOptions)
	exportCmd.Flags().StringVarP(&exportType, "type", "t", "", "export file type (required)")
	exportCmd.Flags().StringVarP(&exportPath, "path", "p", "", "export file path (required)")
	_ = exportCmd.MarkFlagRequired("type")
	_ = exportCmd.MarkFlagRequired("path")

	// importCmd
	importCmd.AddCommand(listImportOptions)
	importCmd.Flags().StringVarP(&importType, "type", "t", "", "import file type (required)")
	importCmd.Flags().StringVarP(&importPath, "path", "p", "", "import file path (required)")
	_ = importCmd.MarkFlagRequired("type")
	_ = importCmd.MarkFlagRequired("path")


	diffBasedOnPath.Flags().StringVarP(&diffFirstFile, "first", "f", "", "first file")
	diffBasedOnPath.Flags().StringVarP(&diffSecondFile, "second", "s", "", "second file")
	_ = diffBasedOnPath.MarkFlagRequired("first")
	_ = diffBasedOnPath.MarkFlagRequired("second")

	// searchCmd
	searchCmd.Flags().StringVarP(&searchPackageName, "name", "n", "", "package name")
	searchCmd.Flags().StringVarP(&searchRootDir, "dir", "d", "", "package root dir")
	searchCmd.Flags().StringVarP(&searchFileOut, "out", "o", "", "spdx file out")
	_ = searchCmd.MarkFlagRequired("name")
	_ = searchCmd.MarkFlagRequired("dir")
	_ = searchCmd.MarkFlagRequired("out")

	// mergeCmd
	mergeCmd.Flags().StringVarP(&mergeFirstFile, "first", "f", "", "first file")
	mergeCmd.Flags().StringVarP(&mergeSecondFile, "second", "s", "", "second file")
	mergeCmd.Flags().StringVarP(&mergeOutput, "out", "o", "", "output pat")
	_ = mergeCmd.MarkFlagRequired("first")
	_ = mergeCmd.MarkFlagRequired("second")
	_ = mergeCmd.MarkFlagRequired("out")

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

	//scn := scanner.FromStr(flg.scanner)
	//cfg := &scanner.Config{Tool: scn, InDir: flg.inDir, ResultDir: "/tmp/pm/"}
	//
	//scanning.Run(
	//	cfg,
	//	rest.NewClient(serverBaseURL),
	//)

}

// callCrawler of crawlerCmd
func callCrawler(name, source, output string) interface{} {
	client := rest.NewClient(serverBaseURL)
	if err := commands.RunCrawler(name, source, output, client); err != nil {
		return err
	}

	return nil
}

// callDiffId of diffCmd
func callDiffId(first, second string) error {
	client := rest.NewClient(serverBaseURL)
	if err := commands.RunDiffById(first, second, client); err != nil {
		return err
	}

	return nil
}

// callDiffId of diffCmd
func callDiffPath(first, second string) error {
	client := rest.NewClient(serverBaseURL)
	if err := commands.RunDiffByPath(first, second, client); err != nil {
		return err
	}

	return nil
}

// callExport of exportCmd
func callExport(exportId, exportType, exportPath string) error {
	postPath := fmt.Sprintf("/products/export/%s", strings.ToLower(exportType))
	client := rest.NewClient(serverBaseURL)
	if err := commands.RunExport(exportId, exportPath, postPath, client); err != nil {
		return err
	}
	return nil
}

// listAvailableExportTypes of exportCmd
func listAvailableExportTypes() {
	availableExportOptions := []string{
		"spdx",
		"human-read",
		"custom-report",
	}
	fmt.Println("Available import options:")
	for key, name := range availableExportOptions {
		fmt.Printf("%v) %v\n", key+1, name)
	}
}

// callImport of importCmd
func callImport(importType, importPath string) error {
	postPath := fmt.Sprintf("/products/import/%s", strings.ToLower(importType))
	client := rest.NewClient(serverBaseURL)
	if err := commands.RunImport(importPath, postPath, client); err != nil {
		return err
	}
	return nil
}

// listAvailableImportTypes of importCmd
func listAvailableImportTypes() {
	availableImportOptions := []string{
		"spdx",
		"composer",
		"file-hasher",
	}
	fmt.Println("Available import options:")
	for key, name := range availableImportOptions {
		fmt.Printf("%v) %v\n", key+1, name)
	}
}

// callMerge of mergeCmd
func callMerge(mergeFirstFile, mergeSecondFile, mergeOutput string) error {
	client := rest.NewClient(serverBaseURL)
	if err := commands.RunMerge(mergeFirstFile, mergeSecondFile, mergeOutput, client); err != nil {
		return err
	}
	return nil
}

// callSearch of searchCmd
func callSearch(searchPackageName, searchRootDir, searchFileOut string) error {
	client := rest.NewClient(serverBaseURL)
	if err := commands.RunSearch(searchPackageName, searchRootDir, searchFileOut, client); err != nil {
		return err
	}
	return nil
}

// printVersion of versionCmd
func callVersion() error {
	if err := commands.RunVersion(gitCommit); err != nil {
		return err
	}
	return nil
}