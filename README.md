# golang_leanchat

## Usage
```
# Install client
go install github.com/starnuik/golang_leanchat@latest

# Run client
golang_leanchat --help


# Install server
curl https://api.github.com/repos/starnuik/golang_leanchat/contents/compose.yml > compose.yml
docker-compose pull

# Run Server
docker-compose up
```