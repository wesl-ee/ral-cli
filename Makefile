GOBUILD=go build
GOGET=go get

RAL_LIB=$(GOPATH)/src/github.com/wesleycoakley/ral/*.go

all: raleexplorer

raleexplorer: raleexplorer.go View.go Config.go $(RAL_LIB)
	go build -o raleexplorer raleexplorer.go View.go Config.go

clean:
	rm raleexplorer
