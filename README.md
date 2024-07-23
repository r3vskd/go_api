# Go API
Testing API's written in go step by step, from basic functions to API construction.

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

# Running the AP
``` go run main.go ```

## Access the API endpoints:
To count runes:
``` http://localhost:8080/count?text=Hello&text=World ```
 
## To find unique strings:
``` http://localhost:8080/unique?text=Hello&text=World&text=Hello ```


