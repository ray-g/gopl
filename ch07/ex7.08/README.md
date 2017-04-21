# Exercise 7.8 (P191)

Many GUIs provide a table widget with a stateful multi-tier sort:
the primary sort key is the most recently clicked column head,
the secondary sort key is the second-most recently clicked column head, and so on.
Define an implementation of `sort.Interface` for use by such a table.
Compare that approach with repeated sorting using `sort.Stable`.
