package commands

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"os/exec"
)

func RunCrawler(name, source, output string, client *rest.Client) error {
	logServerVersion(client)
	// now we just have one crawler plugin then there is no need for checking the name
	// Then the name would be php-scanner
	if name == "php-scanner" {
		name = "docker.pkg.github.com/osrgroup/product-model-toolkit/php-scanner:1.0.0"
	}

	// source should have a slash at the end
	//if okSource := strings.HasSuffix("source", "/"); !okSource {
	//	source = fmt.Sprintf(source+"%v", "/")
	//}
	//
	//if okOutput := strings.HasSuffix("source", "/"); !okOutput {
	//	output = fmt.Sprintf(output+"%v", "/")
	//}

	fmt.Println(source)
	if source == "." {
		source = "$PWD"
	}
	if output == "." {
		output = "$PWD"
	}

	dockerCmd := fmt.Sprintf("sudo docker run"+
		" -e USE_DEFAULT_REPO=0 "+
		"-v %v/source:/source "+
		"-v %v/output:/output %v",
		source, output, name)

	// log information
	fmt.Println("Running crawler")
	// execute docker command
	fmt.Println("Executing the docker command ...")
	// executing the command
	_, err := exec.Command("/bin/sh", "-c", dockerCmd).CombinedOutput()
	// check error
	if err != nil {
		return err
	}

	fmt.Println("Crawling licenses successfully completed")

	// return
	return nil
}
