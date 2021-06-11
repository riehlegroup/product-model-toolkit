package commands

import (
	"errors"
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"os/exec"
)

func RunCrawler(name, source, output string, client *rest.Client) error {
	logServerVersion(client)

	// the formal name for docker image
	var dockerImageName string

	// now we just have one crawler plugin then there is no need for checking the name
	// Then the name would be php-scanner

	switch name {
	case "php-scanner":
		dockerImageName = "docker.pkg.github.com/osrgroup/product-model-toolkit/php-scanner:1.0.0" // TODO
	default:
		return errors.New("invalid crawler name")
	}


	fmt.Println(source)
	if source == "." {
		source = "$PWD"
	}
	if output == "." {
		output = "$PWD"
	}

	// create the dockerCmd from input values
	dockerCmd := fmt.Sprintf("sudo docker run"+
		" -e USE_DEFAULT_REPO=0 "+
		"-v %v:/source "+
		"-v %v:/output %v",
		source, output, dockerImageName)

	// log information
	fmt.Println("Running crawler")

	// execute docker command
	fmt.Println("Executing the docker command ...")

	// print the docker command
	fmt.Println(dockerCmd)

	// executing the command
	_, err := exec.Command("/bin/sh", "-c", dockerCmd).CombinedOutput()
	// check error
	if err != nil {
		return err
	}

	fmt.Println("Crawling licenses successfully completed")
	fmt.Printf("The output path: %v\n", output)


	// return
	return nil
}
