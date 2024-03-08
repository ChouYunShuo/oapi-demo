# OAPI-Codegen Boilerplate

## Introduction

This project is a Go-based application that utilizes `oapi-codegen` for generating boilerplate code from OpenAPI specifications. This approach ensures type-safe integrations of HTTP services defined by the OpenAPI (formerly Swagger) specifications. Parts of this codebase are referenced from the project [chi-demo](https://github.com/tendant/chi-demo/) by github.com/tendant related to server setup and sqlc integration.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Dependencies](#dependencies)
- [Configuration](#configuration)
- [Documentation](#documentation)
- [Examples](#examples)
- [Troubleshooting](#troubleshooting)
- [Contributors](#contributors)
- [License](#license)

## Installation

### Prerequisites

Ensure you have Go installed on your system. This project requires Go 1.22.0.

### Installing `oapi-codegen`

To install the `oapi-codegen` tool, run the following command:

```bash
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
```

This command installs the latest version of `oapi-codegen`, which is essential for generating code based on your OpenAPI specifications.

To successfully run oapi-codegen from the shell, you need to include the Go binary directory in your system's PATH environment variable. You can do this by executing the following command:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

This command appends the directory where Go binaries are stored ($(go env GOPATH)/bin) to your existing PATH, ensuring that the shell can find and execute oapi-codegen."

## Usage

After installing `oapi-codegen`, you can generate HTTP API functions and types with the following commands:

To generate types:

```bash
oapi-codegen -generate types -o idm_types.gen.go -package api idm.yaml
```

To generate the Chi server:

```bash
oapi-codegen -generate chi-server -o idm_server.gen.go -package api idm.yaml
```

These commands will generate Go code in the output files you specify, namely `idm_types.gen.go` and `idm_server.gen.go`, based on the OpenAPI specification file `idm.yaml`. Please adjust the variable names to match your specific file names.

## Features

- Type-safe HTTP API integration based on OpenAPI specifications.
- Automated code generation for boilerplate HTTP handlers and types.
- Easy integration with Chi router for HTTP server functionality.
- Authentication middleware for route protection

## Dependencies

- Go 1.22.0
- `oapi-codegen` for generating code from OpenAPI specs.

Refer to `go.mod` and `go.sum` for a detailed list of dependencies.

## Configuration

For the API to function correctly, you need to set up a local PostgreSQL server. This server will act as the storage system for all the API requests. In addition, you can utilize the provided sqlc configuration file. sqlc is a command-line tool that generates type-safe code from SQL. By using this configuration file, sqlc will automatically generate the necessary database functions.

## Documentation

For more detailed documentation on `oapi-codegen`, visit [https://github.com/deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen).

## Examples

For examples of how to implement handlers using the interface generated by oapi-codegen, please refer to the idmstore.go file located in the private_api or public_api directory. This file showcases the use of the generated code to create handlers that implements the interface generated by `oapi-codegen`.

## Contributors

[Yun-Shuo Chou](https://www.linkedin.com/in/yschou/)

## License

This project is licensed under the [LICENSE](LICENSE).