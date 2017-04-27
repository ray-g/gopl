# Exercise 13.4 (P366)

Depending on C libraries has its drawbacks.
Provide an alternative pure-Go implementation of `bzip.NewWriter` that uses the `os/exec` package to run `/bin/bzip2` as a subprocess.
