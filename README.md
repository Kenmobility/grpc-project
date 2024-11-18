# gRPC Project#

## Requirements- Golang 1.20+
- PostgreSQL

## 1. Clone the repository, cd into the project folder and download required go dependencies
```bash
git clone https://github.com/Kenmobility/grpc-project.git
```
```bash
cd grpc-project
```
```bash
go mod tidy
```
## 2. Unit Testing
Run 'make test' to run the unit tests:
```bash
make test
```

## 3 Open Docker desktop application
- Ensure that docker desktop is started and running on your machine 

## 4. Run the application
- run 'make' to run application
```bash
make
```
## 4 Use Evans cli to test gRPC request
- run 'make evans' to run application
```bash
make evans
```