# Go Concurrency Tasks - Integration to Real World Through Simulations

## 🟢 Part 1: Goroutines & Channels Basics

1. **Concurrent Counter**: Write a program where multiple goroutines increment a shared counter, but synchronize updates via a channel (no `sync.Mutex`).
2. **Ping-Pong**: Two goroutines pass a string `"ping"` and `"pong"` back and forth through a channel N times.
3. **Buffered Logger**: Create a buffered channel logger that queues log messages and a goroutine that writes them to stdout.
4. **File Writer Queue**: Producer goroutine sends lines to a channel, consumer goroutine writes them to a file. Stop gracefully with channel close.
5. **Math Service**: Implement a goroutine that takes requests `(operation, numbers)` from a channel and sends back results through another channel.

## 🟡 Part 2: for-range & Select

6. **Ticker with Stop**: Create a ticker goroutine that sends time every second to a channel. Stop it after 5 ticks with a `done` channel.
7. **Race Timer**: Run two goroutines that finish at random times. Use `select` to report which one finishes first.
8. **Multiplex Input**: Create two producer goroutines (one sends odd numbers, one sends even). Use `select` to consume from both concurrently.
9. **Timeout Fetch**: Simulate fetching data by sleeping. Cancel with `select` if it takes >2 seconds.
10. **Worker Cancellation**: Start a worker that processes jobs from a channel. Use `select` and a `quit` channel to stop it.

## 🟠 Part 3: Producer-Consumer Simulations

11. **Task Queue**: Build a producer-consumer system: producers send tasks, multiple consumers process them concurrently.
12. **Job Dispatcher**: Write a dispatcher that distributes jobs across N workers using buffered channels.
13. **Batch Processor**: Collect jobs into a slice of size 5 using a buffered channel, then process the batch at once.
14. **Bounded Queue**: Implement a bounded queue using a buffered channel, demonstrate blocking when full.
15. **Pipeline**: Chain 3 stages: generate numbers → square them → print results, using goroutines and channels.

## 🔴 Part 4: Errors, Panic, Recover

16. **Safe Division Service**: Worker goroutine receives `(a, b)` and returns `a/b` or an error if `b=0`.
17. **Recover Logger**: Goroutine panics randomly; use `defer + recover` to catch panic and log an error instead of crashing.
18. **Error Pipeline**: Each stage of a pipeline can fail (e.g., parsing, processing). Collect errors into a dedicated error channel.
19. **Retry Logic**: Worker fails randomly. Implement retry logic using channels and errors.
20. **Supervisor**: Start N workers; if any panics, supervisor recovers and restarts it.

## 🔵 Part 5: Structs & Interfaces with Concurrency

21. **Job Interface**: Define an interface `Job { Run() error }`. Implement `EmailJob`, `ReportJob`. Run them concurrently via a worker pool.
22. **Task Scheduler**: Create a struct `Scheduler` with methods `AddTask(func() error)` and `Run()`. Run tasks concurrently and collect errors.
23. **Concurrent Cache**: Build a `Cache` struct that uses a goroutine + channel internally to handle `Get/Set` requests safely.
24. **Publisher-Subscriber**: Implement a pub/sub system using channels. Subscribers register to topics, publishers send messages.
25. **Concurrent Downloader**: Define an interface `Downloader` with `Download(url string) ([]byte, error)`. Run multiple downloaders concurrently and collect results.

## 🟣 Part 6: Reality-Near Simulations

26. **Banking System**: Simulate bank accounts with deposit/withdraw requests sent over channels. Handle overdraft with errors.
27. **Chat Server Simulation**: Multiple clients send messages to a central `chatroom` goroutine, which broadcasts to all clients.
28. **Web Crawler Simulation**: Given a set of URLs (strings), workers fetch them concurrently (simulate with `time.Sleep`), collect successes/errors.
29. **Order Processing System**: Orders flow through a pipeline: validate → process payment → ship. Each stage runs in its own goroutine.
30. **Traffic Light Simulation**: Simulate traffic lights with goroutines. Use `select` to switch between red/green/yellow signals. Panic if two lights go green at once, recover and fix.

💡 These tasks integrate:

* **goroutines** (workers, pipelines, services),
* **channels** (sync, producer-consumer, signaling),
* **for-range & select** (timeouts, multiplexing, cancellation),
* **errors/panic/recover** (graceful error handling),
* **structs/interfaces** (abstraction, real design).

