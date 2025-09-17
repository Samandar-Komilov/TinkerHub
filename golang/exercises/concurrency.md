# Concurrency Exercises

## Goroutines and Channels

### Basics

1. Launch a goroutine that prints numbers 1–10, while main prints A–J. See interleaving.
2. Use a channel to send a value from a worker goroutine to main.
3. Modify (2) so multiple workers send results to one channel.
4. Demonstrate deadlock by reading from a channel with no sender.
5. Show how closing a channel allows range-loop to finish.

### Buffered Channels

6. Use a buffered channel of size 2. Send 3 values. Observe blocking.
7. Implement a simple producer-consumer with a buffered channel.
8. Demonstrate channel capacity with `len(ch)` and `cap(ch)`.
9. Use a buffered channel to rate-limit (like a semaphore).
10. Implement a worker pool with 3 goroutines pulling from a job queue.

### `for-range` & Closing

11. Create a generator goroutine that sends 1–5 and closes. Consume with range.
12. Test receiving from a closed channel (`ok` flag).
13. Make two producers send to one channel, close only when both are done.
14. Show panic when sending to closed channel.
15. Implement fan-out: one channel feeding multiple goroutines.
