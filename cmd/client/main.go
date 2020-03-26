package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/osrgroup/product-model-toolkit/pkg/scanner"
	"github.com/osrgroup/product-model-toolkit/pkg/scanning"
	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

var gitCommit string

type flags struct {
	version    *bool
	lstScanner *bool
	scanner    string
	inDir      string
	set        bool
}

func main() {
	flg := initFlags()
	printInfos(&flg)

	scn := scanner.FromStr(flg.scanner)
	cfg := &scanner.Config{Tool: scn, InDir: flg.inDir, ResultDir: "/tmp/pm/"}

	scanning.Run(cfg)

}

func initFlags() flags {
	version := flag.Bool("v", false, "show version")
	lstScanner := flag.Bool("l", false, "list all available scanner")
	scanner := flag.String("s", "", "scanner to to use from list of available scanner")
	wd, _ := os.Getwd()
	inDir := flag.String("i", wd, "input dir to scan. Default is current working directory")

	flag.Parse()

	set := flag.NFlag() > 0

	return flags{
		version,
		lstScanner,
		*scanner,
		*inDir,
		set,
	}
}

func printInfos(flg *flags) {
	if *flg.version {
		printVersion()
	}

	if *flg.lstScanner {
		listScanner()
	}

}

func printVersion() {
	fmt.Println("PMT Client")
	fmt.Println("----------")
	fmt.Println("Version: " + version.Name())
	fmt.Println("Git commit: " + gitCommit)
	fmt.Println("----------")
}

func listScanner() {
	fmt.Println("Available license scanner:")
	for _, scn := range scanner.Available {
		fmt.Printf("%+v\n", scn)
	}
}
