# HTTP Server created in Go

For every web based application we create there is always the need for a HTTP server to receive and respond calls from clients. It should be reusable, production-ready and highly configurable.

**Flags** are a great resource to easily configure port, enviroments, server behavior, etc, when starting the server.

If you are working in API server you should probably prefix the **version** number in all url path like so: api.com/**v1.0.0**/home. An alternative way is to include it on the headers like `Accept: app/appName-v1`

In this project we could use the standard http.ServeMux but I prefer to implement the [httprouter](https://github.com/julienschmidt/httprouter). It is well tested, reliable and fast **router**. It keeps our `main()` clean and works better for running tests. As you can see below, we update our server's Handler to call `app.routes()` instead of `http.NewServerMux()`

```
srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
```

All handlers will be stored in a routes file under cmd/app directory. 

```
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	return router
}
```

To safely shutdown our app we need to allow our handlers to finish their job through a **graceful shutdown** functionality. By using `os.Signal` we will catch signals like `SIGINT`, `SIGTERM` or `SIGKILL`. 

```
go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		app.logger.Println("shutting down server\n", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()
```

An **audit** system to check, test, update and format our code base just by running a `make audit` command. 

```
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
```

To protect our server from receiving too many request from clients is a good idea to setup **rate limiting** control. This functionality is implemented as a `middleware` so we can appli it on all routes. 



