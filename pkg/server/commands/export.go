package commands

import (
	// "bytes"
	"encoding/json"
	// "errors"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/osrgroup/product-model-toolkit/cnst"
	"github.com/osrgroup/product-model-toolkit/model"
	"github.com/spdx/tools-golang/builder"
	"github.com/spdx/tools-golang/tvsaver"
)

func RunExport(exportId, exportPath string) error {
	// creating a new http client
	client := newClient(cnst.ServerBaseURL)

	// log server version with respect to client
	logServerVersion(client)

	// get the id from the database
	result, err := client.GetProductId(exportId)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}

	var prod model.Product
	if err := json.Unmarshal(body, &prod); err != nil {
		return err
	}


	// get the command-line arguments
	packageName := "test"
	packageRootDir := exportPath

	// to use the SPDX builder package, the first step is to define a
	// builder.Config2_2 struct. this config data can be reused, in case you
	// are building SPDX documents for several directories in sequence.
	config := &builder.Config2_2{

		// NamespacePrefix is a prefix that will be used to populate the
		// mandatory DocumentNamespace field in the Creation Info section.
		// Because it needs to be unique, the value that will be filled in
		// for the document will have the package name and verification code
		// appended to this prefix.
		NamespacePrefix: "https://example.com/whatever/testdata-",

		// CreatorType will be used for the first part of the Creator field
		// in the Creation Info section. Per the SPDX spec, it can be
		// "Person", "Organization" or "Tool".
		CreatorType: "Person",

		// Creator will be used for the second part of the Creator field in
		// the Creation Info section.
		Creator: "Jane Doe",

		// note that builder will also add the following, in addition to the
		// Creator defined above:
		// Creator: Tool: github.com/spdx/tools-golang/builder

		// Finally, you can define one or more paths that should be ignored
		// when walking through the directory. This is intended to omit files
		// that are located within the package's directory, but which should
		// be omitted from the SPDX document.
		PathsIgnored: []string{

			// ignore all files in the .git/ directory at the package root
			"/.git/",

			// ignore all files in all __pycache__/ directories, anywhere
			// within the package directory tree
			"**/__pycache__/",

			// ignore the file with this specific path relative to the
			// package root
			"/.ignorefile",

			// or ignore all files with this filename, anywhere within the
			// package directory tree
			"**/.DS_Store",
		},
	}

	// now, when we actually ask builder to walk through a directory and
	// build an SPDX document, we need to give it three things:
	//   - what to name the package; and
	//   - where the directory is located on disk; and
	//   - the config object we just defined.
	doc, err := builder.Build2_2(packageName, packageRootDir, config)
	if err != nil {
		fmt.Printf("error while building document: %v\n", err)
		return err
	}

	// if we got here, the document has been created.
	// all license info is marked as NOASSERTION, but file hashes and
	// the package verification code have been filled in appropriately.
	fmt.Printf("successfully created document for package %s\n", packageName)

	// we can now save it to disk, using tvsaver.

	// create a new file for writing
	w, err := os.Create("result.spdx")
	if err != nil {
		fmt.Printf("error while opening %v for writing: %v\n", "result.spdx", err)
		return err
	}
	defer w.Close()

	err = tvsaver.Save2_2(doc, w)
	if err != nil {
		fmt.Printf("error while saving %v: %v", "result.spdx", err)
		return err
	}

	fmt.Printf("successfully saved %v\n", "result.spdx")


	// save the spdx file
	
	// return the path of the spdx file

	fmt.Println(prod)
	// return
	return nil
}

