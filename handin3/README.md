# Cryptographic Computing

## How to run
To run the main test in `main.go`, write the following command in the root directory of the project:
```go
go run main.go
```

The program will output amount of cases where the protocol equals the expected value.

## Structure
Following the suggested structure in the assignment description, the protocol runs in a loop of `send`- and `receive` functions, between parties *Alice* and *Bob*. 

The protocol is expected to loop 14 times, and each loop consists of
```
Bob.Receive(Alice.Send())
Alice.Receive(Bob.Send())
```
This protocol finishes when Alice notifies the system that an output is ready.

In `main.go`, the correctness of the output in the protocol is verified by comparing it to the given boolean logic used in handin 1 & 2.