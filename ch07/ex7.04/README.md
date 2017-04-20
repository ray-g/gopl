# Exercise 7.4 (P175)

The `strings.NewReader` function retruns a value that satisfies
the `io.Reader` interface (and others) by reading from its argument, a string.
Implement a simple version of `NewReader` yourself, and use it to make
the HTML parser (§5.2) take input from a string.
