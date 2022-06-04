vendor:
	go mod tidy && go mod vendor
	echo "all code dependencies have being downloaded into the 'vendor' folder"

test: 
	go test ./... -count=1
	go test ./... -v -json > report.json
	go test ./... -v -coverprofile=coverage.out

server:
	go run main.go server

user:
	go run main.go user

simulate:
	eval './simulate.sh $(CURDIR)'
