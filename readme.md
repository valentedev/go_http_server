# HTTP Server created in Go

For every web based application we create there is always the need for a HTTP server to receive and respond calls from clients. It should be reusable, production-ready and highly configurable.

`flags` are a great resource to easily configure port, enviroments, server behavior, etc, when starting the server.

If you are working in API server you should probably prefix the `version` number in your url path like so: api.com/**v1.0.0**/home. An alternative way is to include it on the headers like `Accept: app/appName-v1`

In this project we could use the standard http.ServeMux but I prefer to implement the [httprouter](https://github.com/julienschmidt/httprouter). It is well tested, reliable and fast. It keeps our `main()` clean and works better for running tests. As you can see below, we update our server's Handler to call app.routes() instead of `http.NewServerMux()`

```
srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
```


