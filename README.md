# Cisco Smart Bonding

[![Status](https://img.shields.io/badge/status-wip-yellow)](https://github.com/darrenparkinson/ciscosmartbonding) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/darrenparkinson/ciscosmartbonding) ![GitHub](https://img.shields.io/github/license/darrenparkinson/ciscosmartbonding?color=brightgreen) [![GoDoc](https://pkg.go.dev/badge/darrenparkinson/ciscosmartbonding)](https://pkg.go.dev/github.com/darrenparkinson/ciscosmartbonding) [![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/ciscosmartbonding)](https://goreportcard.com/report/github.com/darrenparkinson/ciscosmartbonding) [![Coverage](https://codecov.io/gh/darrenparkinson/ciscosmartbonding/branch/main/graphs/badge.svg?branch=main)](https://codecov.io/gh/darrenparkinson/ciscosmartbonding) [![CI](https://github.com/darrenparkinson/ciscosmartbonding/workflows/CI/badge.svg)](https://github.com/darrenparkinson/ciscosmartbonding/actions?query=workflow%3CI)

A Go library for interacting with the [Cisco Smart Bonding API](https://cisco-test.solvedirect.com/ws/).

Currently a work in progress.  This README will get updated as it progresses.

## Usage

**Import the library**

```go
import "github.com/darrenparkinson/ciscosmartbonding"
```

**Initialise a client:**

```go
c := ciscosmartbonding.NewClient(clientID, secret, nil)
```

You can optionally provide your own Resty v2 client.

**Retrieve TSP Codes with the client**

```go
res, err := c.GetAllTSPCodes(context.Background())
```

## Examples

There are various examples to perform the various steps required. They all rely on having the 
appropriate environment variables which can be within a `.env` file at the top level of the repository.

### Retriever

Firstly, there is the `retriever` example.  This will perform the pull update requests every 30 seconds 
and display summary information. 

This should be opened in a separate window whilst you run the other samples:

```sh
$ go run ./examples/0-retriever-service -env development
```

Refer to the [README](examples/0-retriever-service/README.md) for more information.

### Create Ticket

This can be used to create a new shadow ticket.  You will need a lot more variables for this example.  
Refer to the [README](examples/2-create-ticket/README.md) for more information.

```sh
$ go run ./examples/2-create-ticket
```

There is a variation on this to create an an escalated P1 ticket:

```sh
$ go run ./examples/9-create-ticket-escalated-p1
```

You should update the details to use a new ticket number each time.

### Update Work Notes

This demonstrates the use of a helper function and the equivalent PushUpdate function for updating work notes in a ticket.

### TSP Codes

This provides an example of retrieving the TSP codes. It will output the total number
of TSP codes available and save them to a CSV. Optionaly you can specify to save as json.

```sh
$ go run ./examples/1-tsp
Retrieved 19078 codes
```

### Pull Update

An example of pulling down updates.  This is similar to the retriever example, but performs a single pull.