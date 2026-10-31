# Go notes (Sprint 1)

How a request flows
An HTTP server receives a request, matches the URL to a handler function, and that function writes the response. Since Go 1.22 the standard router matches the method and path (GET /v1/cakes/{id}) and gives you r.PathValue("id").
    • Go blog: Routing enhancements in Go 1.22
    • Go by Example: HTTP server
Packages
A package is a folder of .go files that share a name. Names starting with a capital letter (NewService, Cake) can be used from other packages; lowercase ones (cakeHandlers, writeJSON) are private to the package. Code inside internal/ can only be imported by your own module.
    • How to write Go code (packages and modules)
    • Organizing a Go module (why cmd/ and internal/)
Errors are values
Go has no exceptions. Functions return an error as their last value, and you check it right away with if err != nil. Fixed errors like ErrNotFound are compared with errors.Is, which still works when an error has been wrapped with more context.
    • Tour: Errors
    • Go by Example: Errors
    • Go blog: Working with errors (errors.Is)
Interfaces are implicit
An interface is a list of methods. Any type that has those methods satisfies it automatically, with no implements keyword. store.Memory has ListCakes and GetCake, so it is a cakes.Repository, and a DynamoDB store with the same two methods will be one too.
    • Tour: Interfaces are implemented implicitly
    • Go by Example: Interfaces
defer
defer f() schedules f to run when the surrounding function returns, whichever return it takes. It keeps cleanup next to the setup: lock, then defer unlock. Deferred calls run in reverse order.
    • Tour: Defer
    • Go blog: Defer, panic and recover (also explains recoverPanics in the middleware)
Pointer receivers
func (s *Service) Quote(...) means the method works on a pointer to the Service, not a copy. Use pointer receivers when the method changes the value, or when copying it would be wasteful or wrong. Memory contains a mutex, which must never be copied, so its methods use *Memory.
    • Tour: Pointer receivers
    • Go by Example: Methods
Struct tags
The text in backticks after a field, such as json:"amount", is metadata that packages read at runtime. encoding/json uses it to name the JSON field, so Amount becomes amount. Other libraries use their own tags, like dynamodbav in Sprint 2.
    • Go by Example: JSON
    • Go blog: JSON and Go
Closures and middleware
A closure is a function that uses variables from the place where it was created. cors(allowed) returns a function that still remembers allowed. Middleware is a function that takes a handler and returns a new handler that does something before or after calling the original.
    • Tour: Function closures
    • Go by Example: Closures
    • Alex Edwards: Making and using HTTP middleware
Bonus: goroutines and signals in main.go
go func(){...}() runs code in the background as a goroutine. <-stop waits on a channel until a value arrives, which is how the program waits for Ctrl+C before shutting down.
    • Tour: Goroutines
    • Go by Example: Signals
For a quick reference on any concept, keep Go by Example bookmarked: one short runnable page per topic.
