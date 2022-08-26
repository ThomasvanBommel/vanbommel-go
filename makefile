objects = src/main.go src/RequestHandler.go

build: $(objects)
	go build -o run $(objects)

nobuild: $(objects)
	go run $(objects)

clean: run
	rm run
