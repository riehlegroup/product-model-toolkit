// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	//"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"github.com/osrgroup/product-model-toolkit/pkg/client/scanner"
	//"github.com/osrgroup/product-model-toolkit/pkg/client/scanning"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)


var gitCommit string
var (
	importType string
	importPath string
)
const serverBaseURL = "http://localhost:8081/api/v1"

type flags struct {
	scanner string
	inDir   string
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Product Model Toolkit",
	Long: `Product Model Toolkit for Managing Open Source Dependencies in Products`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Show the version of Product Model Toolkit",
	Long: "This command will show the current using version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import the component graph",
	Long: `Import the component graph from spdx, composer or file-hasher`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")
	},
}

var listImportOptions = &cobra.Command{
	Use:   "list",
	Short: "List all available import options",
	Long: `List all available file types for importing as a product into the PMt`,
	Run: func(cmd *cobra.Command, args []string) {
		listAvailableImportOptions()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(versionCmd)
	importCmd.AddCommand(listImportOptions)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	importCmd.Flags().StringVarP(&importType, "type", "t", "", "import file type (required)")
	importCmd.MarkPersistentFlagRequired("type")

	importCmd.Flags().StringVarP(&importPath, "path", "p", "", "import file path (required)")
	importCmd.MarkPersistentFlagRequired("path")
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
	Execute()


	//flg, abort := checkFlags()
	//if abort {
	//	os.Exit(0)
	//}
	//
	//scn := scanner.FromStr(flg.scanner)
	//cfg := &scanner.Config{Tool: scn, InDir: flg.inDir, ResultDir: "/tmp/pm/"}
	//
	//scanning.Run(
	//	cfg,
	//	rest.NewClient(serverBaseURL),
	//)

}

func checkFlags() (flags, bool) {
	version := flag.Bool("v", false, "show version")

	lstScanner := flag.Bool("l", false, "list all available scanner")

	scanner := flag.String("s", "", "scanner to to use from list of available scanner")
	wd, _ := os.Getwd()
	inDir := flag.String("i", wd, "input dir to scan. Default is current working directory")

	flag.Parse()

	if *version {
		printVersion()
	}

	if *lstScanner {
		listScanner()
	}

	abortAfterFlags := *version || *lstScanner

	return flags{
			*scanner,
			*inDir,
		},
		abortAfterFlags
}

func printVersion() {
	fmt.Printf(
		"PMT Client\n----------\nVersion: %s\nGit commit: %s\n",
		version.Name(),
		gitCommit,
	)
}


func listAvailableImportOptions() {
	availableImportOptions := []string{
		"spdx",
		"composer",
		"file-hasher",
	}
		fmt.Println("Available import options:")
	for key, name := range availableImportOptions {
		fmt.Printf("%v) %v\n",key+1, name)
	}
}

func listScanner() {
	fmt.Println("Available license scanner:")
	for _, scn := range scanner.Available {
		fmt.Printf("----------\nName:    %s\nVersion: %s\nImage:   %s\n", scn.Name, scn.Version, scn.DockerImg)
	}
	fmt.Printf("----------\n")
}
