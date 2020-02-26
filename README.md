<p align="center"><img src="logo.png"></p>

Product Model Toolkit for Managing Open Source Dependencies in Products

![Build](https://github.com/osrgroup/product-model-toolkit/workflows/Go/badge.svg)

## Introduction

The Product Model Toolkit helps you to manage third-party open source dependencies in your product. The toolkit itself is not a license scanner. Instead, it facilitates other license scanners to incorporate license and other information together with architectural information into a unified model.

## Architecture Overview

![Architecture concept](docs/architecture-concept-v0.3.png)

1. The small CLI client shall facilitate already existing scanners. For that, it will start Docker container which itself contains the actual scanner and its dependencies. The result then will be sent to the server application or stored as a file for further use. This should help to compare the developed PHP specific deep scanner with other tools.

2. The server application contains all functionalities needed to generate a software bill-of-materials (SBOM) of a product, represented by the elements in the figure. It is also responsible for storing a component graph into a database.

3. A database optimized for graphs shall store the data. The DB shall provide a GraphQL interface, or allows to add a GraphQL interface to it.

4. The PHP scanner performs a deep analysis of a web project and sends its result as a standardized representation (like the CLI client) to the server.

## Project Structure

```
├── cmd/.................Main applications of this project which will be compiled as executables
│   ├── client/
│   │   └── main.go......Client application entry point
│   └── server/
│       └── main.go......Server application entry point
├── docs/................Documentation
├── pkg/.................Library code for client and server
├── model/...............The model for representing a software product
└── README.md
```

## How to Use

All important commands needed to build, test, and run the applications are represented as Makefile rule.
All available rules can be displeyed with `make help`.

### Build

Build with `make build` the client and server application. The generated artifacts are `pmtserver` and `pmtclient`.

### Test

Test with `make test`. This also produces a code coverage report as a file called `coverage.out`.

### Run

Already built executables can be executed directly. For exampl `./pmtclient` or `./pmtserver`.

During development, `go run cmd/client/main.go` or `go run cmd/server/main.go` can be used to run a specific application.

## Requirements

Here you can find the functional requirements for the toolkit. We strive to implement these features in an agile fashion.

**BOM Generation**

- [ ] The system shall generate BOM artifacts as SDPX document.
- [ ] The system shall generate BOM artifacts as human readable representation.
- [ ] The system will be able to provide BOM information for custom reports.

**SPDX support**

- [ ] The system shall import the component graph from a SPDX document.
- [ ] The system shall export the component graph as SPDX document.
- [ ] The system shall import licence information from a SPDX document.

**Diff**

- [ ] The system shall validate if two component graphs are the same.
- [ ] The system shall validate if two components are the same.
- [ ] The system shall present the difference in components between two component graphs of the same product.
- [ ] The system shall present the difference in meta-data between two component graphs of the same product.

**Search**

- [ ] The system shall be able to search for components by its name.
- [ ] The system shall be able to search for components by its meta-data.

**Data Merging**

- [ ] The system should merge license information from different sources into a SDPX license identifier representation.
- [ ] The system should merge sub component graphs into the component graph.

**Client Crawler**

- [ ] The crawler should be executable in a CI environment.
- [ ] The crawler shall be able to facilitate other scanners running in Docker containers to collect license information.
- [ ] The crawler shall send scanned information to the server application via HTTP calls (REST).
- [ ] The crawler shall store scanned information as structured representation (SPDX, SBOM, etc.) as files.


## License

TBA
