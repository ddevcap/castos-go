[![Go Reference](https://pkg.go.dev/badge/github.com/ddevcap/castos-go.svg)](https://pkg.go.dev/github.com/ddevcap/castos-go)

# Castos Golang Client

## About
This client allows you to integrate with the Castos API using the golang programming language. More information about Castos here: https://castos.com/

## Install
```bash
go get github.com/ddevcap/castos-go
```

## Example
```go
token := "apitoken" // Castos API token.
podcastId := 1234   // Castos podcast id.

client := castos.NewClient(token)

episodes, err := client.Episodes.GetAll(podcastId)
if err != nil {
	log.fatal(err)
}

for _, episode := range episodes {
	// Do awesome episode stuff
}
```

## WIP