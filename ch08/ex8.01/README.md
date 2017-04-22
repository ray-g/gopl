# Exercise 8.1 (P222)

Modify `clock2` to accept a port number, and write a program, `clockwall`, that acts as a client of several clock servers at once,
reading the times from each one and displaying the results in a table, akin to the wall of clocks seen in some business offices.
If you have access to `geographically` distributed computers, run instances remotely;
otherwise run local instances on different ports with fake time zones.

``` text
    $ TZ=US/Eastern    ./clock2 -port 8010 &
    $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
    $ TZ=Europe/London ./clock2 -port 8030 &
    $ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
```
