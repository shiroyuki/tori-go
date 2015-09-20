GCC_GO_TEST=go test

test:
	${GCC_GO_TEST} ./tameshigiri
	${GCC_GO_TEST} ./cache
	${GCC_GO_TEST} ./re
	${GCC_GO_TEST} ./web
	${GCC_GO_TEST}

full_test:
	${GCC_GO_TEST} -benchmem -cover ./web
