# Exercise 8.15 (P256)

Failure of any client program to read data in timely manner ultimately causes all clients to get stuck.
Modify the broadcaster to skip a message rather than wait if a client writer is not ready to acceptit.
Alternatively, add buffering to each client's outgoing message channel so that most messages are not dropped; the boradcaster should use a non-blocking send to this channel.
