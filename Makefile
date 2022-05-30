#
# Makefile to perform "live code reloading" after changes to .go files.
#
# To start live reloading run the following command:
# $ make serve
#

# binary name to kill/restart
PROG = 2022Q2GO-Bootcamp

# clean up
clean:
	go clean

# run formatting tool and build
build: clean
	go fmt
	go build

# attempt to kill running server
kill:
	-@killall -9 $(PROG) 2>/dev/null || true

# attempt to build and start server
serve:
	@make kill
	@make build;
	./${PROG}