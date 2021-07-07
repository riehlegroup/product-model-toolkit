// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package cnst

const (
	ServerBaseURL = "http://localhost:8081/api/v1"
)

// commands descriotion
const (
	Cli      = "cli"
	CliShort = "Product Model Toolkit"
	CliLong  = `Product Model Toolkit for Managing Open Source Dependencies in Products`

	Scanner      = "scanner"
	ScannerShort = "Scan the licenses"
	ScannerLong  = `Scan the licenses`

	Diff      = "diff"
	DiffShort = "Difference between two component graphs"
	DiffLong  = `Difference between two component graphs`

	Id      = "id"
	IdShort = "Difference based on id"
	IdLong  = `Difference based on the id of saved products`

	Path      = "path"
	PathShort = "Difference based on path"
	PathLong  = `Difference based on the path of spdx files`

	Export      = "export"
	ExportShort = "Export the component graph"
	ExportLong  = `Export the component graph from spdx, composer or file-hasher`

	Import      = "import"
	ImportShort = "Import the component graph"
	ImportLong  = `Import the component graph from spdx, composer or file-hasher`

	Merge      = "merge"
	MergeShort = "Merge two components"
	MergeLong  = `Merge two components`

	Search      = "search"
	SearchShort = "Search for components"
	SearchLong  = `Search for components by their name and meta-data`

	Version      = "version"
	VersionShort = "Show the version of Product Model Toolkit"
	VersionLong  = "This command will show the current using version of the application"

	List             = "list"
	ListShortScanner = "List all available scanners"
	ListLongScanner  = `List all available scanners for selecting as a license crawler`
	ListShortImport  = "List all available import types"
	ListLongImport   = `List all available file types for importing as a product into the PMT`
	ListShortExport  = "List all available export types"
	ListLongExport   = `List all available file types for exporting from BoM artifacts`
)

// general
const (
	APIGroup = "/api/v1"
)

// docker images
const (
	LicenseeImage   = "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-licensee:9.13.0"
	ScancodeImage   = "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-scancode:3.1.1"
	PhpscannerImage = "docker.pkg.github.com/osrgroup/product-model-toolkit/php-scanner:1.0.0"
	ComposerImage   = "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-composer:dummy"
	FilehasherImage = "docker.pkg.github.com/osrgroup/product-model-toolkit/file-hasher:latest"
)
