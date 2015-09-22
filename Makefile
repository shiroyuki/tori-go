GCC_GO=GOPATH=`pwd`:$GOPATH go
GCC_GO_TEST=${GCC_GO} test
TEST_PATHS=./tameshigiri ./cache ./re ./web .

test:
	${GCC_GO_TEST} ${TEST_PATHS}

full_test:
	${GCC_GO_TEST} -v -benchmem -cover ${TEST_PATHS}
