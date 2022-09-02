source_files = src/main.go src/routes.go
test_files   = 

build_filename = run.build
test_filename  = testing.test

build: $(source_files)
	go build -o $(build_filename) $(source_files)

nobuild: $(source_files)
	go run $(source_files)

test: $(source_files) $(test_files)
	go test -v -c -o $(test_filename) $(source_files) $(test_files) && \
	./$(test_filename) && \
	rm $(test_filename)

clean:
	rm -f $(build_filename) $(test_filename)
