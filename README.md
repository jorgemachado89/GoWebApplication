# GoWebApplication
Go web application server.

Built considering lessons took from PluralSight course: [
Creating Web Applications with Go](https://app.pluralsight.com/library/courses/creating-web-applications-go-update)

## Notes

### HTTP client

* In order to build the client type ```npm run grunt``` and to run it ```npm run server```

### Go Server
* GOPATH set to projects root directory: $HOME/Projects/go/go-web-application/server
* In order for ```go install main``` cmd to work, file structure %PROJECT_ROOT/src/main/ must be in place.
* When dealing with Handlers in Go to match URLs Go is gonna pick the one that has the most specific match of all patterns.
* Handle vs HandleFunc Handlers provide greater flexibility.
* There are several Built In Handlers that provide many of the repetitive logic required during web development.
    * NotFoundHandler - Returns a 404 error to the requester
    * RedirectHandler - Redirects requests to the passed URL
    * StripPrefix - Specifies a handler to handle URL that does not expect the passed prefix to exist
    * TimeoutHandler - Decorates the passed handler and timeouts according to the max duration passed and returns to the requester the msg also passed.
    * FileServer - Takes in a FileSystem object to implement possible custom FS although most of the times it will be passed the OS FS to serve local files.
* The defer keyword defers the execution of the function call until the surrounding function returns.
* Purpose of templates is to bind data to templates to generate documents
* When considering template imports html vs text escapes text being pulled from the data. Great for dealing with security vulnerabilities.
    * Use text import when dealing with non HTML content.
* Reusing template names overwrites the original. The last defined template with duplicate name stands.
* Defining an empty template allows for optionaly redifining that same template in a later stage.
    * The "block" function allows for the same effect.
* Pipelines containing functions and/or method calls expect these to return one or two values, being the second of the error type.
* When dealing with template curly braces adding a minus "-" sign:
    * right after the opening braces will trim all aforegoing white spaces 
    * doing it before the closing braces will trim all preceding ones
