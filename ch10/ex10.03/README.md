# Exercise 10.3 (P293)

Using `fetch http://gopl.io/ch1/helloworld?go-get=1`, find out which service hosts the code samples for this book.
(HTTP requests from `go get` include the `go-get` parameter so that servers can disginguish them from ordinary browser requests.)

## Result

```shell
go run fetch.go http://gopl.io/ch1/helloworld\?go-get\=1 | grep go-import
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
```
