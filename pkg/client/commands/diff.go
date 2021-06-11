package commands

import (
	"fmt"
	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"github.com/spdx/tools-golang/licensediff"
	"github.com/spdx/tools-golang/spdxlib"
	"github.com/spdx/tools-golang/tvloader"
	"os"
)

func RunDiffById(first, second string) error {
	//logServerVersion(client)
	//
	//log information
	//fmt.Println("Running the diff by Id")
	//
	//open the first spdx file
	//r, err := os.Open(first)
	//if err != nil {
	//	fmt.Printf("Error while opening %v for reading: %v", first, err)
	//	return err
	//}
	//defer r.Close()
	//
	//try to load the first SPDX file's contents as a tag-value file, version 2.2
	//docFirst, err := tvloader.Load2_2(r)
	//if err != nil {
	//	fmt.Printf("Error while parsing %v: %v", first, err)
	//	return err
	//}
	//
	//check whether the SPDX file has at least one package that it describes
	//pkgIDsFirst, err := spdxlib.GetDescribedPackageIDs2_2(docFirst)
	//if err != nil {
	//	fmt.Printf("Unable to get describe packages from first SPDX document: %v\n", err)
	//	return err
	//}
	//
	// if we got here, the file is now loaded into memory.
	//fmt.Printf("Successfully loaded first SPDX file %s\n", first)

	// open the second SPDX file
	//r, err = os.Open(second)
	//if err != nil {
	//	fmt.Printf("Error while opening %v for reading: %v", second, err)
	//	return err
	//}
	//defer r.Close()
	//
	// try to load the second SPDX file's contents as a tag-value file, version 2.2
	//docSecond, err := tvloader.Load2_2(r)
	//if err != nil {
	//	fmt.Printf("Error while parsing %v: %v", second, err)
	//	return err
	//}
	// check whether the SPDX file has at least one package that it describes
	//pkgIDsSecond, err := spdxlib.GetDescribedPackageIDs2_2(docSecond)
	//if err != nil {
	//	fmt.Printf("Unable to get describe packages from second SPDX document: %v\n", err)
	//	return err
	//}

	// if we got here, the file is now loaded into memory.
	//fmt.Printf("Successfully loaded second SPDX file %s\n\n", second)

	// compare the described packages from each Document, by SPDX ID
	// go through the first set first, report if they aren't in the second set
	//for _, pkgID := range pkgIDsFirst {
	//	fmt.Printf("================================\n")
	//	p1, okFirst := docFirst.Packages[pkgID]
	//	if !okFirst {
	//		fmt.Printf("Package %s has described relationship in first document but ID not found\n", string(pkgID))
	//		continue
	//	}
	//	fmt.Printf("Package %s (%s)\n", string(pkgID), p1.PackageName)
	//	p2, okSecond := docSecond.Packages[pkgID]
	//	if !okSecond {
	//		fmt.Printf("  Found in first document, not found in second\n")
	//		continue
	//	}
	//
	//	now, run a diff between the two
		//pairs, err := licensediff.MakePairs2_2(p1, p2)
		//if err != nil {
		//	fmt.Printf("  Error generating licensediff pairs: %v\n", err)
		//	continue
		//}

		// take the pairs and turn them into a more structured results set
		//resultSet, err := licensediff.MakeResults(pairs)
		//if err != nil {
		//	fmt.Printf("  Error generating licensediff results set: %v\n", err)
		//	continue
		//}

		// print some information about the results
		//fmt.Printf("  Files in first only: %d\n", len(resultSet.InFirstOnly))
		//fmt.Printf("  Files in second only: %d\n", len(resultSet.InSecondOnly))
		//fmt.Printf("  Files in both with different licenses: %d\n", len(resultSet.InBothChanged))
		//fmt.Printf("  Files in both with same licenses: %d\n", len(resultSet.InBothSame))
	//}
	//
	//now report if there are any package IDs in the second set that aren't
	// in the first
	//for _, pkgID := range pkgIDsSecond {
	//	p2, okSecond := docSecond.Packages[pkgID]
	//	if !okSecond {
	//		fmt.Printf("================================\n")
	//		fmt.Printf("Package %s has described relationship in second document but ID not found\n", string(pkgID))
	//		continue
	//	}
	//	_, okFirst := docFirst.Packages[pkgID]
	//	if !okFirst {
	//		fmt.Printf("================================\n")
	//		fmt.Printf("Package %s (%s)\n", string(pkgID), p2.PackageName)
	//		fmt.Printf("  Found in second document, not found in first\n")
	//	}
	//}
	//

	// return
	return nil
}

func RunDiffByPath(first, second string) error  {
	logServerVersion(client)

	// log information
	fmt.Println("Running the diff by Id")

	// open the first spdx file
	r, err := os.Open(first)
	if err != nil {
		fmt.Printf("Error while opening %v for reading: %v", first, err)
		return err
	}
	defer r.Close()

	// try to load the first SPDX file's contents as a tag-value file, version 2.2
	docFirst, err := tvloader.Load2_2(r)
	if err != nil {
		fmt.Printf("Error while parsing %v: %v", first, err)
		return err
	}

	// check whether the SPDX file has at least one package that it describes
	pkgIDsFirst, err := spdxlib.GetDescribedPackageIDs2_2(docFirst)
	if err != nil {
		fmt.Printf("Unable to get describe packages from first SPDX document: %v\n", err)
		return err
	}

	// if we got here, the file is now loaded into memory.
	fmt.Printf("Successfully loaded first SPDX file %s\n", first)

	// open the second SPDX file
	r, err = os.Open(second)
	if err != nil {
		fmt.Printf("Error while opening %v for reading: %v", second, err)
		return err
	}
	defer r.Close()

	// try to load the second SPDX file's contents as a tag-value file, version 2.2
	docSecond, err := tvloader.Load2_2(r)
	if err != nil {
		fmt.Printf("Error while parsing %v: %v", second, err)
		return err
	}
	// check whether the SPDX file has at least one package that it describes
	pkgIDsSecond, err := spdxlib.GetDescribedPackageIDs2_2(docSecond)
	if err != nil {
		fmt.Printf("Unable to get describe packages from second SPDX document: %v\n", err)
		return err
	}

	// if we got here, the file is now loaded into memory.
	fmt.Printf("Successfully loaded second SPDX file %s\n\n", second)

	// compare the described packages from each Document, by SPDX ID
	// go through the first set first, report if they aren't in the second set
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

