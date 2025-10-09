---
title: "Go's net/http library - Deep Dive"
description: "Go's net/http library - Deep Dive"
date: 2025-10-10
categories:
  - Backend
tags:
    - go
---


In this post we are gonna explore the standard `net/http` library internals - how it works and why it works the way it works.

<!-- more --> 

!!! abstract "Contents"
    1. HTTP Server Basics
        - `http.Request`
        - `http.ResponseWriter`
        - `http.Handle` and `http.HandleFunc`
        - `http.Handler` and `http.HandlerFunc`
        - `http.Serve` and `http.ListenAndServe`
    2. Context Usage
        - `request.Context()`
    3. Routing
        - `http.ServeMux` and `http.NewServeMux`
        - `http.DefaultServeMux`
    4. Middleware Composition
    5. Request Body, Streaming and Large Uploads
        - `r.Body (io.ReadCloser)`, `io.Copy`
        - `http.MaxBytesReader`, `multipart.Reader`
    6. HTTP Client Basics
        - `http.Client`, `http.NewRequest`, `Client.Do`, `http.Get/Post`
    7. Transport
        - `http.Transport`


## 1. HTTP Server Basics

### http.Request

The most important construct in building web applications is `Request`, no doubt. When we built our own [Nginx clone](https://github.com/Samandar-Komilov/cserve) in C, we too defined that entity at first. For this reason, I believe it worth examining `http.Request` struct initially which is defined in the standard `net/http` library:
```go
type Request struct {
    Method string
    URL *url.URL
    Proto string // e.g. "HTTP/1.0"
    Header Header
    Body io.ReadCloser
    ContentLength int64
    TransferEncoding []string
    Close bool
    Host string
    Form url.Values
    PostForm url.Values
    MultipartForm *multipart.Form
    Trailer Header
    TLS *tls.ConnectionState
    ctx context.Context
}
```

### http.ResponseWriter

It's nice that HTTP request is parsed into `http.Request` construct automatically, with all necessary fields available. But how we return the response? I mean, we should somehow build the corresponding `http.Response` construct, isn't it? But how? Before answering the question, let's consider the `http.Response` struct:
```go
type Response struct {
    Status string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto string // e.g. "HTTP/1.0"
    Header Header
    Body io.ReadCloser
    ContentLength int64
    TransferEncoding []string
    Close bool
    Trailer Header
    Request *Request
    TLS *tls.ConnectionState
}
```

Perfect, should I build the response by hand? No, Go handles that for us, but partially. It creates the response object with basic fields, but lets us change it - using `http.ResponseWriter` interface. A `http.ResponseWriter` interface is used by an HTTP handler to construct an HTTP response. Let's look at its structure as well:
```go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```


### http.Handle and http.HandleFunc

We analyzed the `Request` and `Response` constructs, but how we can "glue" them together that result an API handler? Well, that's now easy as we already know everything needed:
```go
func myhandler(w http.ResponseWriter, r *http.Request){
    fmt.Printf(w, "Hello World!\n")
}
```
In Go, any function that has the `http.ResponseWriter` and `*http.Request` parameters is considered as a handler. This function is simply responding to the request with `Hello World!`, but now another question arises again - how does it know which path it responds to? `/hello`?   

At this point, we introduce `http.HandleFunc` that associates the handler with the path. Here is a quick example:
```go
func main() {
	http.HandleFunc("/my", myhandler)
}
```


### http.Serve and http.ListenAndServe

Although, it seems we are done, when we run the code nothing works. Well, it's natural - where does it know which port it is listening to? Every server application should listen at some specific port to accept requests and respond accordingly, but where is that? Here:
```go
func main() {
	http.HandleFunc("/my", myhandler)

    http.ListenAndServe(":8090", nil)
}
```