MODULE = github.com/stevesloka/envoy-xds-server

GO_BUILD_VARS = \
	github.com/projectcontour/contour/internal/build.Version=${BUILD_VERSION} \
	github.com/projectcontour/contour/internal/build.Sha=${BUILD_SHA} \
	github.com/projectcontour/contour/internal/build.Branch=${BUILD_BRANCH}

GO_LDFLAGS := -s -w $(patsubst %,-X %, $(GO_BUILD_VARS))

build:
	mkdir -p build
	mkdir -p configs
	go build -o build/bin/envoy-xds-server -mod=readonly -v -ldflags="$(GO_LDFLAGS)" $(MODULE)/cmd/server

clean:
	rm -rf build
	rm test/configs/*