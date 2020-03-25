package scanner

// Scanner struct define a scanner tool
type Scanner struct {
	Name      string
	Version   string
	DockerImg string
}

// Available list all available scanner tools that can be used.
var Available = [...]Scanner{
	Licensee,
	Scancode,
}

// Licensee represents the latest usable Licensee version
var Licensee = Scanner{
	Name:      "Licensee",
	Version:   "9.13.0",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-licensee:9.13.0",
}

// Scancode represents the latest usable Scancode version
var Scancode = Scanner{
	Name:      "Scancode",
	Version:   "3.1.1",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-scancode:3.1.1",
}
