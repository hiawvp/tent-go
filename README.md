# tent-go
Rewriting info282 project because why not

## About

Writing the backend in Go but I'm not really committed to this right now, just playing around with testing and trying out Vim plugins.

## For future me

`cd tento`

- Live reload with [Air](https://github.com/cosmtrek/air) 

`air`

http://localhost:5150/api/v1/

- Run a single test, use `gotest` for colors
 
`./gotest ./test/api -run TestPostProduct -v`

- Run all the tests

`./gotest ./test/... -v  `

- test with coverage

`go test -cover -coverpkg "./pkg/..." "./test/..." `

## Resources

https://gorm.io/docs/index.html

https://github.com/go-playground/validator

https://go.dev/doc/tutorial/web-service-gin

https://github.com/golang-standards/project-layout
