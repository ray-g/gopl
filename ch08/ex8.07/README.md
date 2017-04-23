# Exercise 8.7 (P243)

Write a concurrent program that creates a local mirror of a web site,
fetching each reachable page and writing it to a directory on the local disk.
Only pages within theoriginal domain (for instance, `golang.org`) should be fetched.
URLs within mirrored pages should be altered as needed so that they refer to the mirrored page, not the original.
