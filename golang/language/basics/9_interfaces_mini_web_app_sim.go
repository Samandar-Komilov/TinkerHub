package basics

import "fmt"

func Main_interfaces_mini_app() {
	// Choose DI storage implementation
	storage := RealStorage{} // or MockStorage{}

	// Choose DI logger implementation
	logger := ConsoleLogger{} // or FileLogger{}

	// Core handler
	saveHandler := SaveHandler{storage: storage}

	// Wrap handler with middleware
	loggedHandler := LoggingMiddleware(logger)(saveHandler)

	// Simulate request
	req := Request{Path: "/submit", Body: "Hello, Go!"}
	loggedHandler.Handle(req)
}

// Interfaces
type Handler interface {
	Handle(req Request)
}

type Storage interface {
	Save(data string)
}

type Loggerr interface {
	Log(msg string)
}

// Types
type Request struct {
	Path string
	Body string
}
type HandlerFunc func(Request)

type RealStorage struct{}
type MockStorage struct{}

type SaveHandler struct {
	storage Storage
}

type ConsoleLogger struct{}
type FileLogger struct{}

type Middleware func(Handler) Handler

// Methods
func (f HandlerFunc) Handle(req Request) {
	f(req)
}

func (r RealStorage) Save(data string) {
	fmt.Println("Saving to Real DB...", data)
}

func (m MockStorage) Save(data string) {
	fmt.Println("Saving to Mock DB...", data)
}

func (h SaveHandler) Handle(req Request) {
	h.storage.Save(req.Body)
}

func (c ConsoleLogger) Log(msg string) {
	fmt.Println("LOG:", msg)
}

func (f FileLogger) Log(msg string) {
	fmt.Println("File Log:", msg)
}

func LoggingMiddleware(logger Loggerr) Middleware {
	return func(next Handler) Handler {
		return HandlerFunc(func(req Request) {
			logger.Log("Received request: " + req.Path)
			next.Handle(req)
		})
	}
}
