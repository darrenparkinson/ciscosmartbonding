# Cisco Smart Bonding

[![Status](https://img.shields.io/badge/status-wip-yellow)](https://github.com/darrenparkinson/ciscosmartbonding) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/darrenparkinson/ciscosmartbonding) ![GitHub](https://img.shields.io/github/license/darrenparkinson/ciscosmartbonding?color=brightgreen) [![GoDoc](https://pkg.go.dev/badge/darrenparkinson/ciscosmartbonding)](https://pkg.go.dev/github.com/darrenparkinson/ciscosmartbonding) [![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/ciscosmartbonding)](https://goreportcard.com/report/github.com/darrenparkinson/ciscosmartbonding) [![Coverage](https://codecov.io/gh/darrenparkinson/ciscosmartbonding/branch/main/graphs/badge.svg?branch=main)](https://codecov.io/gh/darrenparkinson/ciscosmartbonding) [![CI](https://github.com/darrenparkinson/ciscosmartbonding/workflows/CI/badge.svg)](https://github.com/nwillc/snipgo/actions?query=workflow%3CI)

A Go library for interacting with the [Cisco Smart Bonding API](https://cisco-test.solvedirect.com/ws/).

Currently a work in progress.  This README will get updated as it progresses.

## Usage

**Import the library**

```go
import "github.com/darrenparkinson/ciscosmartbonding"
```

**Initialise a client:**

```go
c := ciscosmartbonding.NewClient(username, password, nil)
```

You can optionally provide your own Resty v2 client.

**Retrieve TSP Codes with the client**

```go
res, err := c.GetTSPCodes(context.Background())
```

**Use Pagination**

```go
res, err := c.GetTSPCodes(context.Background(), &ciscosmartbonding.ListParams{Limit: ciscosmartbonding.Int64(5000)})
```

**Or use the convenience functions**

```go
res, err := c.GetAllTSPCodes(context.Background())
```
