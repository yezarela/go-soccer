# go-soccer

Simple soccer REST API

## How to run this project

This project is using go module, so you don't have to clone it inside GOPATH directory.

### Run the app

```
# Clone the repository
git clone https://github.com/yezarela/go-soccer.git

# Navigate to project directory
cd go-soccer

# Install dependencies
make deps

# Run the application
make run

# Hit the endpoint
curl localhost:1323/teams
```

If you want to use your own mongodb uri, you can update the .env file with your own mongodb uri

### Run the test
```
make test
```
