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
* Templates will consider any of the following an empty value as part of logical tests:
    * false
    * zero
    * nil
    * empty collection
* When dealing with Logical tests all commands are evaluated even if the one of the parts is false.
* There are some third party providing routes handling although we should always consider if the overhead of implementing our own routing mechanism exceeds the gain of using a tool built by someone else. Two valuable choices when considering going elsewhere:
    * [Gorilla Mux](https://github.com/gorilla/mux)
    * [HttpRouter](https://github.com/julienschmidt/httprouter)
* HTTP Middleware common use cases:
    * Logging
    * Security Validations
    * Request timeouts
    * Response compression
* Contexts are useful to apply a context based on the request context
* Since Golang 1.6 modules are mandatory by default.
* When using HTTPS in Go, the network protocol used will be HTTP/2
* For using the PostGres Database the developer should start by initializing the driver prior to creating a connections pool.
* HTTP/2
    * Allows for request multiplexing
        * In each TCP connection allows for multiple streams in the same connection without interference
        * In each stream information is sent in frames instead of sending both body and header as a single unattached block
    * Header Compression
        * Splitting information in chuncks allows for better optimization
            * Sending headers (C->S) in indeppendent frames allows for the server to start processing the request as soon as the minimum ammount of information necessary is available.
            * Since both Client and Server will now know they are receiving solely headers they can apply compress and decompress algorithms.
        * Reduces the ammount of most of the times fixed data volume being sent back and forth
    * Secure by default
        * THis isn't part of the protocol since it can work via HTTP and HTTPS although most organizations are not complying with non secure HTTP connections
    * Server Push
        * Allows for preemptively request resources such as CSS file dependencies without having the client explicitly request the resource.
        * When such a request would have been made the resource would already be cached in the client side, rendering the request avoidable.
        * Validation for Server Push assertion needs to be set up since not all clients are guaranteed to support it.
        * Beware of using Server Push because the server is going to push the resource not knowing if the client is going to need it or not.