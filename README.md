# GoWebApplication
Go web application server

## Notes

### HTTP client

* In order to build the client type ```npm run grunt``` and to run it ```npm run server```

### Go Server
* GOPATH set to projects root directory: $HOME/Projects/go/go-web-application/server
* In order for ```go install firstapp``` cmd to work, file structure %PROJECT_ROOT/src/firstapp/ must be in place.
* When dealing with Handlers in Go to match URLs Go is gonna pick the one that has the most specific match of all patterns.
* Handle vs HandleFunc Handlers provide greater flexibility.
* There are several Built In Handlers that provide many of the repetitive logic required during web development.
    * NotFoundHandler - Returns a 404 error to the requester
    * RedirectHandler - Redirects requests to the passed URL
    * StripPrefix - Specifies a handler to handle URL that does not expect the passed prefix to exist
    * TimeoutHandler - Decorates the passed handler and timeouts according to the max duration passed and returns to the requester the msg also passed.
    * FileServer - Takes in a FileSystem object to implement possible custom FS although most of the times it will be passed the OS FS to serve local files.
* The defer keyword defers the execution of the function call until the surrounding function returns.