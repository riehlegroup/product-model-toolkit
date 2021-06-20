package main


var (
	// config file
	cfgFile string

	// crawlerCmd
	crawlerName   string
	crawlerOutput string
	crawlerSource string

	// diffCmd
	diffFirstId    string
	diffSecondId   string
	diffFirstFile  string
	diffSecondFile string

	// exportCmd
	exportId   string
	exportType string
	exportPath string

	// importCmd
	importType string
	importPath string

	// mergeCmd
	mergeFirstFile  string
	mergeSecondFile string
	mergeOutput     string

	// searchCmd
	searchPackageName string
	searchRootDir     string
	searchFileOut     string

	// git commit version
	gitCommit string
)
