// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"fmt"
	"github.com/spdx/tools-golang/licensediff"
	"github.com/spdx/tools-golang/spdxlib"
	"github.com/spdx/tools-golang/tvloader"
	"os"
)

func RunDiffByPath(first, second string) error  {
	// log information
	fmt.Println("Running the diff by path")

	// open the first spdx file
	r, err := os.Open(first)
	if err != nil {
		fmt.Printf("error while opening %v for reading: %v", first, err)
		return err
	}
	defer r.Close()

	// try to load the first SPDX file's contents as a tag-value file, version 2.2
	docFirst, err := tvloader.Load2_2(r)

	// check the error
	if err != nil {
		fmt.Printf("error while parsing %v: %v", first, err)
		return err
	}

	// check SPDX file for packages
	pkgIDsFirst, err := spdxlib.GetDescribedPackageIDs2_2(docFirst)

	// check the erro
	if err != nil {
		fmt.Printf("unable to get describe packages from first SPDX document: %v\n", err)
		return err
	}

	// print the successful message -> first file is loaded
	fmt.Printf("successfully loaded first SPDX file %s\n", first)

	// open the second file
	r, err = os.Open(second)

	// check the error
	if err != nil {
		fmt.Printf("Error while opening %v for reading: %v", second, err)
		return err
	}

	// close the file in defer
	defer r.Close()

	// load the second file as tag-value
	docSecond, err := tvloader.Load2_2(r)

	// check the error
	if err != nil {
		fmt.Printf("Error while parsing %v: %v", second, err)
		return err
	}

	// check SPDX file for packages
	pkgIDsSecond, err := spdxlib.GetDescribedPackageIDs2_2(docSecond)

	// check the error
	if err != nil {
		fmt.Printf("Unable to get describe packages from second SPDX document: %v\n", err)
		return err
	}

	// print the successful message -> second file is loaded
	fmt.Printf("Successfully loaded second SPDX file %s\n\n", second)

	// compare the packages by SPDX ID
	for _, pkgID := range pkgIDsFirst {
		fmt.Printf("================================\n")
		p1, okFirst := docFirst.Packages[pkgID]
		if !okFirst {
			fmt.Printf("Package %s has described relationship in first document but ID not found\n", string(pkgID))
			continue
		}
		fmt.Printf("Package %s (%s)\n", string(pkgID), p1.PackageName)
		p2, okSecond := docSecond.Packages[pkgID]
		if !okSecond {
			fmt.Printf("  Found in first document, not found in second\n")
			continue
		}

		// now, run a diff between the two
		pairs, err := licensediff.MakePairs2_2(p1, p2)
		if err != nil {
			fmt.Printf("  Error generating licensediff pairs: %v\n", err)
			continue
		}

		// take the pairs and turn them into a more structured results set
		resultSet, err := licensediff.MakeResults(pairs)
		if err != nil {
			fmt.Printf("  Error generating licensediff results set: %v\n", err)
			continue
		}

		// print some information about the results
		fmt.Printf("  Files in first only: %d\n", len(resultSet.InFirstOnly))
		fmt.Printf("  Files in second only: %d\n", len(resultSet.InSecondOnly))
		fmt.Printf("  Files in both with different licenses: %d\n", len(resultSet.InBothChanged))
		fmt.Printf("  Files in both with same licenses: %d\n", len(resultSet.InBothSame))
	}

	// now report if there are any package IDs in the second set that aren't
	// in the first
	for _, pkgID := range pkgIDsSecond {
		p2, okSecond := docSecond.Packages[pkgID]
		if !okSecond {
			fmt.Printf("================================\n")
			fmt.Printf("Package %s has described relationship in second document but ID not found\n", string(pkgID))
			continue
		}
		_, okFirst := docFirst.Packages[pkgID]
		if !okFirst {
			fmt.Printf("================================\n")
			fmt.Printf("Package %s (%s)\n", string(pkgID), p2.PackageName)
			fmt.Printf("  Found in second document, not found in first\n")
		}
	}

	// return
	return nil
}

