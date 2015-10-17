CURR_PATH=`pwd`
GOPATH_DEPS=${CURR_PATH}/lib
GCC_GO=GOPATH=${GOPATH_DEPS} go
GCC_GO_TEST=${GCC_GO} test
GCC_GO_GET=${GCC_GO} get
TEST_PATHS=.

test: install_deps
	${GCC_GO_TEST} ${TEST_PATHS}

full_test: install_deps
	${GCC_GO_TEST} -v -benchmem -cover ${TEST_PATHS}

install_deps:
	@echo "Installing dependencies..."
	@mkdir -p ${GOPATH_DEPS}
	${GCC_GO_GET} github.com/shiroyuki/tameshigiri
	${GCC_GO_GET} github.com/shiroyuki/yotsuba-go
	${GCC_GO_GET} github.com/shiroyuki/re
