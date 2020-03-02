package scanner

// Scanner struct define a scanner tool
type Scanner struct {
	Name      string
	Version   string
	DockerImg string
}

// Available lsit all available scanner tools that can be used.
var Available = [...]Scanner{
	Scanner{
		Name:      "Licensee",
		Version:   "9.13.0",
		DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-licensee:9.13.0",
	},
}
