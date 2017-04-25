# Exercise 9.3 (P279)

Extend the Func type and the `(*Memo).Get` method so that callers may provide an optional done channel through which they can cancel the operation (§8.9).
The results of a cancelled Func call should not be cached.
