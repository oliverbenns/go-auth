# Go Auth

An experimental project to authenticate users in Go.

## Project Features/Notes

- It uses the bcrypt hashing algorithm for passwords.
- It uses cookie-based JWTs for user verification.
- It uses no Javascript at all.
- There is no server side validation for forms (like sign up).
- Error handling isn't particularly well done.
- Test cases don't handle every case.
- Jwts & Cookies are not set to expire.

## Running 

- Setup: `go run setup/setup.go`
- Start server: ` go build app/*.go && ./app/command-line-arguments`
- Test: `go test ./app`