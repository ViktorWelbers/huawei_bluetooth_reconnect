build:
	go build  -ldflags "-H windowsgui"  -o headphones.exe

run:
	go run main.go