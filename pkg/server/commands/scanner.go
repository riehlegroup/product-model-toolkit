package commands

import (
	"errors"
	"fmt"
	"strings"
	"os/exec"

	"github.com/osrgroup/product-model-toolkit/cnst"
)

// Scanner provides license scanner operations.
type Scanner interface {
	Exec(cfg Config)
}

// Tool struct define a scanner tool
type Tool struct {
	Name      string
	Version   string
	DockerImg string
	Cmd       string
	Results   []string
}

// Config represents a configuration for a tool to execute.
type Config struct {
	Tool
	InDir     string
	ResultDir string
}

// Available list all available scanner tools that can be used.
var Available = [...]Tool{
	PhpScanner,
	Licensee,
	Scancode,
	Composer,
	FileHasher,
}

// Default is the default scanner tools that shall be used if no particular tool is selected.
var Default Tool = Licensee

// Licensee represents the latest usable Licensee version
var Licensee = Tool{
	Name:      "licensee",
	Version:   "9.13.0",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-licensee:9.13.0",
	Cmd:       `/bin/bash -c "licensee detect /input/ --json > /result/result.json"`,
	Results:   []string{"result.json"},
}

// Scancode represents the latest usable Scancode version
var Scancode = Tool{
	Name:      "scancode",
	Version:   "3.1.1",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-scancode:3.1.1",
	Cmd:       `/bin/bash -c "./scancode --spdx-tv /result/result.spdx /input"`,
	Results:   []string{"result.spdx"},
}


// Composer represents a container which produces a Composer result artifact.
var Composer = Tool{
	Name:      "composer",
	Version:   "dummy",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-composer:dummy",
	Cmd:       `/bin/sh -c "cp example.json result/example.json"`,
	Results:   []string{"example.json"},
}

// FileHasher represents a container which produces a JSON file with hashes of all file of the input folder.
var FileHasher = Tool{
	Name:      "filehasher",
	Version:   "latest",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/file-hasher:latest",
	Cmd:       `/bin/sh -c "./fh -i /input -o /result/result.json"`,
	Results:   []string{"result.json"},
}

var PhpScanner = Tool {
	Name: "phpscanner",
	Version: "1.0.0",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/php-scanner:1.0.0",
	Cmd: `/bin/sh -c  "phpScanner.php --sourcedir=<path/to/scanned/folder> --outputdir=<path/to/output/folder>"`, // TODO
	Results: []string{"phpScanner.json"},

}

// FromStr returns a tool with the given name. If unable to find a tool with the given name it returns the default tool.
func FromStr(name string) Tool {
	search := strings.ToLower(name)
	for _, t := range Available {
		if strings.ToLower(t.Name) == search {
			return t
		}
	}

	return Default
}

// String return the name and the version of the tool.
func (t *Tool) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.Version)
}


func RunScanner(name, source, output string) error {
	// creating a new http client
	client := newClient(cnst.ServerBaseURL)

	// log server version with respect to client
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

	// define the url of respected http call
	url := ""

	// create a json data for sending to the server
	var jsonStr = []byte(fmt.Sprintf(`{"path":%v}`, output))

	// send the results to the server
	client.postData(url, jsonStr)

	fmt.Println("Crawling licenses successfully completed")
	fmt.Printf("The output path: %v\n", output)


	// return
	return nil
}


func ListAvailableScannerTypes() error {
	// print instruction, loop over the
	// list and print the available options
	fmt.Println("Available scanner options:")
	for key, item := range Available {
		fmt.Printf("%v) %v\n", key+1, item.Name)
	}

	return nil
}