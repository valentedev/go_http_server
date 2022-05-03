# HTTP Server created in Go

For every web based application we create there is always the need for a HTTP server to receive and respond calls from clients. In this post we are going to create a basic server written in Go that should be production ready, highly configurable, ready to hook databases and much more. 

First of all let's create a folder for our project.

> `$ mkdir httpserver-go`

Since this is a Go project we should init our mod file. My preference is to use 'repository/username/filename' as the mod path as you can see below. 

> `$ cd httpserver-go`

> `$ go mod init github/valentedev/httpserver-go`

Now we can create the directory structure of our application that should look like this:

```
- bin  
- cmd
  - app
        main.go
- go.mod
- Makefile
```

> `mkdir -p bin cmd/app`

> `touch Makefile cmd/app/main.go`

Now let's check our app is working by adding some code to `main.go` file

```
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```
> `$ go run ./cmd/app`

We should have "Hello world!" printed in our terminal. Great! Now we can move on and start writting our http server.

