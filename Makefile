tidy:
	go mod tidy 
	go mod vendor 

build:
	go build  -o gopad ./cmd 