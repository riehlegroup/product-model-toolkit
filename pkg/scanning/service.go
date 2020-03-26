package scanning

import (
	"log"

	"github.com/osrgroup/product-model-toolkit/pkg/scanner"
)

// Run executes a scan with a scanner tool for a given directory.
func Run(cfg *scanner.Config) {
	log.Printf("[Scanner] Selected : %v", cfg.Tool.String())
	log.Printf("[Scanner] Input directory: %v", cfg.InDir)
	log.Printf("[Scanner] Result directory: %v", cfg.ResultDir)
}
