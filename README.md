# Go Auth

An experimental project to authenticate users in Go.

## Project Features/Notes

- It uses the bcrypt hashing algorithm for passwords.
- It uses cookie-based JWTs for user verification.
- It uses no Javascript at all.
- There is no server side validation for forms (like sign up).

## Running 

- Setup: `go run setup/setup.go`
- Start server: `go run app/*.go`
