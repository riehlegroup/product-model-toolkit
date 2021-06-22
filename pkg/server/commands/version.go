package commands

import (
	"fmt"
)

func RunVersion(gitCommit string) error {
	fmt.Printf(
		"PMT Client\n----------\nVersion: %s\nGit commit: %s\n",
		"1.0.0",
		gitCommit,
	)

	return nil
}
