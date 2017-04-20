# Exercise 7.5 (P175)

The `LimitReader` function in the io package accepts an `io.Reader r` and a number of bytes `n`,
and returns another `Reader` that reads from `r` but reports an end-of-file condition after n bytes.
Implement it.

``` Go
    func LimitReader(r io.Reader, n int64) io.Reader
```
