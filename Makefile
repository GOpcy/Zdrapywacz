compile: 
	go build -o bin/output.exe

run: compile
	./bin/output.exe