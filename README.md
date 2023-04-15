# Simple-Web-Backend-ChatGPT

Hello! Welcome to my Simple Web Backend experiment with ChatGPT. Feel free to clone the repo and test it out! If you want to know how I built this with ChatGPT, you can read the full blog [here](https://medium.com/@romalms10/how-to-build-a-simple-web-backend-with-chatgpt-784e031a485f)

## Get It Running
`go run main.go`

## Test that it works
Open a new window in your terminal and try out the following commands:

### See which people are already in memory
```
$ curl http://localhost:8080/people
[]
```
### Add a Person
```
$ curl -X POST -H "Content-Type: application/json" -d '{"name": "Alice", "age": 27}' http://localhost:8080/people
$
```
### See if your person is now in memory
```
$ curl http://localhost:8080/people
[{"name":"Alice","age":27}]
```

### Delete people from memory
```
$ curl -X DELETE -H "Content-Type: application/json" -d '"Alice"' http://localhost:8080/people
$
```

### Check which people are left
```
$ curl http://localhost:8080/people
[]
```