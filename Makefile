tidy:
	go mod tidy 
	go mod vendor 

run:
	go run ./cmd editor/editor.go