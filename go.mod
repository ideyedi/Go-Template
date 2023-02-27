module github.com/ideyedi/go_tutorial

go 1.19

require (
	github.com/gin-gonic/gin v1.3.0
	go.template/models v0.0.0-00010101000000-000000000000
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/ugorji/go/codec v1.2.10 // indirect
	go.template/faker v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// Local modules
replace go.template/models => ./models

replace go.template/faker => ./faker
