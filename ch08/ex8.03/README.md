# Exercise 8.3 (P227)

In `netcat3`, the interface value conn has the concrete type `*net.TCPConn`, which represents a TCP connection.
A TCP connection consists of two halves that may be closed independently using its `CloseRead` and `CloseWrite` methods.
Modify the main goroutine of `netcat3` to close only the write half of the connection so that the program will continue to print the final echoes from the `reverb1` server even after the standard input has been closed.
(Doing this for the `reverb2` server is harder; see Exercise 8.4)
