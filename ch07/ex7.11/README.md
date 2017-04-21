# Exercise 7.11 (P195)

Add additional handlers so that clients can create, read, update, and delete database entries.
For example, a request of the form `/update?item=socks&price=6` will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid.
(Warning: this change introduces concurrent variable update.)
