# Exercise 8.4 (P239)

Modify the `reverb2` server to use a `sync.WaitGroup` per connection to count the number of active `echo` goroutines.
When it falls to zero, close the write half of the TCP connection as described in Exercise 8.3.
Verify that your modified `netcat3` client from that exercise waits fro the final echos of multiple concurrent shouts,
even after the standard input has been closed.
