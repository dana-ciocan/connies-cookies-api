# Creating an API using Go and Gin

I'm learning Go so I'm trying out this tutorial: https://go.dev/doc/tutorial/web-service-gin

## How to run this API

- Clone this repo
- Run `go run .`

You should now see the API at http://localhost:8181 (not it's HTTP not HTTPS)

## Available routes

- [/cookies](http://localhost:8181/cookies) - see a list of available cookies
- [/cookies/:id](http://localhost:8181/cookies/1) - see a particular cookie, selected by its ID
