files = $(wildcard src/*.go)

outfile = "out/Release/x64/vizz"

test: $(files)
	go run $(files)

build: $(files)
	go build -o $(outfile) $(files) 
