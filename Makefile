BINARY_NAME=clipse
INSTALL_DIR?=/usr/local/bin/

wayland:
	CGO_ENABLED=0 go build -tags wayland -o $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

install: build
	install -m 755 $(BINARY_NAME) $(INSTALL_DIR)

clean:
	go clean
	rm -f $(BINARY_NAME)
