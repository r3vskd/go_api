# Go API
Testing API's written in go step by step, from basic functions to API construction

# Resources
https://github.com/gorilla/mux

# Why i use go for this API?
Go have:
- Fast compile time
- Built in concurrency
- Simplicity

# Go modules and packages structure

![alt text](https://github.com/r3vskd/go_api/blob/main/images/Screenshot_2.png)

# Basic data types in go

![alt text](https://github.com/r3vskd/go_api/blob/main/images/Screenshot_3.png)

# Go Project Structure

![alt text](https://github.com/r3vskd/go_api/blob/main/images/Screenshot_4.png)

## Key concepts about Golang:

- Arrays and Slices: Used for input data and unique string processing.
- Maps: Used to count character occurrences.
- Loops: Used for iterating over input data and results.
- Channels and Goroutines: Goroutines and channels in Go are the primary tools for managing concurrency.
- Functions and Pointers: Used in the processString function.
- Strings and Runes: Processed to count letter and digit occurrences.
- Generics: Used to create a reusable filterUnique function.
- Structs and Interfaces: Used to define and implement the Processor interface.
- Concurrency: Concurrency is about dealing with lots of things at once.

## Common packages i Golang:

- fmt: For formatted I/O operations.
- sync: Provides synchronization primitives, such as WaitGroup.
- time: For time-related functions, used here to simulate work with time.Sleep.

## What is defer for?

Defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

## Slices

![alt text](https://github.com/r3vskd/go_api/blob/main/images/Untitled-Diagram38.jpg)

## Maps
![alt text](https://github.com/r3vskd/go_api/blob/main/images/0qFxv.jpg)

# Running the API
``` go run main.go ```

## Access the API endpoints:
To count runes:
``` http://localhost:8080/count?text=Hello&text=World ```
 
## To find unique strings:
``` http://localhost:8080/unique?text=Hello&text=World&text=Hello ```


