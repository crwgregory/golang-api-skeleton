# Setup:
- Use the db.json template in the config dir to set the db settings and put it in docker/settings
- Copy the aws credentials file to the same directory

## Commands
`docker build --rm -t api-skeleton .`

`docker run -it --rm --name api-skeleton-container -p 8080:8080 -v {your-home}\go\src\github.com\crwgregory\golang-api-skeleton:/go/src/github.com/crwgregory/golang-api-skeleton -w /go/src/github.com/crwgregory/golang-api-skeleton api-skeleton`

`go get`

`go build && ./golang-api-skeleton`