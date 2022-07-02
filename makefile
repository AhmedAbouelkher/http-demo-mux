build:
	go build -o server .

run: build
	./server

watch:
	ulimit -n 1000
	reflex -s -r '\.go$$' make run