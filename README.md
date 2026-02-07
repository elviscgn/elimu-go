# elimu-api 

but in golang (mostly porting the logic I will do in the java springboot and compare the two, so far i'm enjoying golang more who would've guessed)

## To get started
1. Edit the `.env.example` file to `.env` then fill in your google client id and secrets, you can find out how to get these with a google search "google oauth credentials"
2. To run the server `go run cmd/main.go`
3. To run tests `go test ./...`
4. To generate documentation `swag init -g cmd/api/main.go` make sure you have swaggo installed


## Features
### Modules
- [x] Identity
- [ ] LMS bot in rust (in progress)
- [ ] User management (`student || staff`)
- [ ] Curriculum & Projects
- [ ] Documents
- [ ] Communications
- [ ] Audit
