GCC_GO_TEST=go test
TEST_PATHS=./tameshigiri ./cache ./re ./web .

test:
	${GCC_GO_TEST} ${TEST_PATHS}

full_test:
	${GCC_GO_TEST} -benchmem -cover ${TEST_PATHS}
