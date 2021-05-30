package commands

import (
	"fmt"
	version2 "github.com/osrgroup/product-model-toolkit/pkg/services/version"
)

func RunVersion(gitCommit string) error {
	fmt.Printf(
		"PMT Client\n----------\nVersion: %s\nGit commit: %s\n",
		version2.Name(),
		gitCommit,
	)

	return nil
}
