# Exercise 12.8 (P347)

The `sexpr.Unmarshal` function, like `json.Marshal`, requires the complete input in a byte slice before it can begin decoding.
Define a `sexpr.Decoder` type that, like `json.Decoder`, allows a sequence of values to be decoded from an `io.reader`.
Change `sexpr.Unmarshal` to use this new type.
