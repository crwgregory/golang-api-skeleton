# Golang Skeleton API
This is a project to help you get up an running fast with a secure, light weight, extensible api structure.

## Skill level
***Intermediate***
This project gives you a basic structure to work with. There is no ORM.
### Prerequisites:
- Docker
- Secret Key Generation and handling (for JWT's)
- MySQL

## Features
- Built in load balancer
- MC (MVC with out the view) esc structure
- Request logging to NoSql DynamoDB
- JWT Generation for user authentication


## Application Structure
### Routing & Request workflow:
Routing in the API takes a slightly different approach from an MVC architecture.
Step by step:
1. ***Initial Request:*** Request comes in and a `handler` is found that maps to the base request. See `main.go`'s routes.
2. ***Handlers:*** The handler registers it's own routes and passes them to a routing function that parses the path and 
determines which callback function to call.
3. ***Handler Functions:*** The called handler function instantiates the proper `record` and calls the associated functions and returns the resulting response.
4. ***Records:*** A Record is a representation of a database table in the form a struct. It holds all the necessary ***Queries***.
6. ***Connection***: A connection is a connection to a data source.

## Other Things
### Request Logging
The api is set up to log requests to a dynamodb on AWS. You must supply the proper credentails for the api to connect to the service.
See amazons credential documentation. You can comment out where the log is pushed to the log channel in the RouteController 
function in the `handler.go` file if you want to **disable** logging.

### Creating a super secret key
Run the php script in the templates folder in the docker folder, that will generate a byte key that will then be copied to 
the docker container and be used for generating jwt's. The key can be used with other crypting algorithms like blowfish.

### Verbose
Verbose can be turned off in the Dockerfile under API_VERBOSE. Verbose will print err's and stack traces when set in the api response.
 
### External libraries
[SHA3](http://golang.org/x/crypto/sha3)
