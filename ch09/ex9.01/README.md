# Exercise 9.1 (P262)

Add a function `Withdraw(amount int) bool` to the `gopl.io/ch9/bank1` program.
The result should indicate whether the transaction succeeded or failed due to insufficient funds.
The message send to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back to `Withdraw`.
