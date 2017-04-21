# Exercise 7.7 (P181)

Explain why the help message contains `°C` when the default value of `20.0` does not.

**Answer**: The flag's value is a `Stringer`, and is used in command-line help messages.
`Celsius` has a `func (c Celsius) String() string` method, so `Celsius` is a `Stringer`.
When `flag` shows the help messages, it will call `Celsius' String` method to format the value.
