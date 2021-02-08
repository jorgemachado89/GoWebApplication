# GoWebApplication
Go web application server

## Notes

### HTTP client

* In order to build the client type ```npm run grunt``` and to run it ```npm run server```

### Go Server
* GOPATH set to projects root directory: $HOME/Projects/go/go-web-application/server
* In order for ```go install firstapp``` cmd to work, file structure %PROJECT_ROOT/src/firstapp/ must be in place.
* When deaaling with Handlers in Go to match URLs Go is gonna pick the one that has the most specific match of all patterns. 