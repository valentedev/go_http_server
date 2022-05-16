# HTTP Server created in Go

For every web based application we create there is always the need for a HTTP server to receive and respond calls from clients. It should be reusable, production-ready and highly configurable. 

**Flags**
easily configure port, environments, and many others server functionalities. Here you have one example of how to use flags to define a list of Trusted Origins in a CORS situation. 

```
flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
	cfg.cors.trustedOrigins = strings.Fields(val)
	return nil
})
```

Run the following command with trusted URLs.
`go run ./cmd -cors-trusted-origins="http://localhost:3000 http://localhost:3001"`

**Versioning**

If you are working in API server, you should probably prefix the version number in all URL path like so: `api.com/v1/home`. An alternative way is to include it on the headers: `Accept: app/appName-v1`

**Graceful shutdown**

To safely shutdown our app we need to allow our handlers to finish their job through a graceful shutdown functionality. By using `os.Signal` we will catch signals like `SIGINT`, `SIGTERM` or `SIGKILL`. We use a buffed channel with size 1 to avoid missing any signals. 

`quit := make(chan os.Signal, 1)`


We create another channel to receive the HTTP server method `Shutdown()` with a configurable TimeOut context. 

`shutdownError := make(chan error)`

`shutdownError <- srv.Shutdown(ctx)`

During the context timeout execution the app will log a "shutting down server" message.

<!-- **Audit**

Will assist you to check, test, update and format our code base just by running a `make audit` command.  -->

**Rate limiting**

To protect our server from receiving too many requests from clients is a good idea to setup rate limiting control. This functionality is implemented as a `middleware` so we can apply it on all routes. 

First, we use `"github.com/tomasen/realip"` to grab the ip in the Request call `ip := realip.FromRequest(r)`. This ip number will stored as a string in the index of a map called "clients" `clients = make(map[string]*client)`. In this case, "client" is a struct that has a LIMITER (`"golang.org/x/time/rate"`) and LASTSEEN (time.Time) fields. 

```go
type client struct {
	limiter  *rate.Limiter
	lastseen time.Time
}
```

The client.limiter has a method `Allow()` that returns a bool that will control the request by releasing the request or returning a 429 error. 

**Panic recover**

We create a middleware that check if there is a `panic` by using the built in `recover()` that will release the request to the next handler or throw an 500 error. 

Check the code on [go_http_server](https://github.com/valentedev/go_http_server) Github repository.