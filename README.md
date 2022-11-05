# Castos golang Client

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