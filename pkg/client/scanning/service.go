package scanning

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/osrgroup/product-model-toolkit/pkg/client/scanner"
)

// Run executes a scan with a scanner tool for a given directory.
func Run(cfg *scanner.Config) {
	log.Printf("[Scanner] Selected : %v", cfg.Tool.String())
	log.Printf("[Scanner] Input directory: %v", cfg.InDir)
	log.Printf("[Scanner] Result directory: %v", cfg.ResultDir)

	execDockerCall(cfg)
	checkResults(cfg)
}

func execDockerCall(cfg *scanner.Config) {
	dockerCmd := execStr(cfg)
	log.Println("[Docker] ", dockerCmd)

	_, err := exec.Command("/bin/sh", "-c", dockerCmd).CombinedOutput()
	if err != nil {
		log.Printf("[Docker] Error during execution: %v", err.Error())
	}
}

func execStr(cfg *scanner.Config) string {
	return fmt.Sprintf("docker run --rm -v %s:/input -v %s:/result %s %s", cfg.InDir, cfg.ResultDir, cfg.Tool.DockerImg, cfg.Tool.Cmd)
}

func checkResults(cfg *scanner.Config) {
	infos, err := ioutil.ReadDir(cfg.ResultDir)
	if err != nil {
		log.Printf("[Docker] Error during checking files: %v", err)
	}
	for i, f := range infos {
		log.Printf("[Docker] Result %v: \n%s", i, f.Name())

		if f.Name() == "result.json" {
			path := cfg.ResultDir + f.Name()
			log.Printf("Found : %v", path)
			data, _ := ioutil.ReadFile(path)
			log.Printf("\n%s", data)
		}
	}
}
