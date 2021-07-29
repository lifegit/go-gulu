# swagr

## 📌What's

#### swag brother

[#396](https://github.com/swaggo/swag/issues/386#issuecomment-833913287) I think this method (use docker) is too cumbersome.
And I need to change something else. So I use [kin-openapi](https://github.com/getkin/kin-openapi) to convert to openapi 3 doc.

## ✌️change

1. add output openapi3.
2. delete output doc.go.
3. abandon "." in openapi doc definitions. example:  "web.APIError" -> "webAPIError". [#903](https://github.com/swaggo/swag/issues/903)


## 📦 install

```bash
# go get
go get https://github.com/lifegit/go-gulu/pkg/swagr

echo "go build && ./swagr init -h"

# source code
git clone https://github.com/lifegit/go-gulu
cd pkg/swagr
go mod download

go install
echo "add swagr to system path"
```

### ⚠️command

```bash
[root]# go build && ./swagr init -h
Automatically generate RESTful API documentation with Swagger 2.0 and convert to openapi 3.0 (spec3) for Go.

Usage:
  swagr init [flags]

Examples:
swagr init -g http/api.go

Flags:
  -c, --codeExampleFiles string   Parse folder containing code example files to use for the x-codeSamples extension, disabled by default
  -d, --dir string                Directory you want to parse (default "./")
      --exclude string            Exclude directories and files when searching, comma separated
  -g, --generalInfo string        Go file path in which 'swagger general API Info' is written (default "main.go")
  -h, --help                      help for init
  -m, --markdownFiles string      Parse folder containing markdown files to use as description, disabled by default
  -o, --output string             Output directory for all the generated files(swagger.json, swagger.yaml) (default "./docs")
      --parseDependency           Parse go files in outside dependency folder, disabled by default
      --parseDepth int            Dependency parse depth (default 100)
      --parseInternal             Parse go files in internal packages, disabled by default
      --parseVendor               Parse go files in 'vendor' folder, disabled by default
  -p, --propertyStrategy string   Property Naming Strategy like snakecase,camelcase,pascalcase (default "camelcase")
      --version                   version for init
```


## 🙏thank
- [Jordan Moore/OneCricketeer](https://github.com/OneCricketeer)
- [swaggo/swag](https://github.com/swaggo/swag)
- [getkin/kin-openapi](https://github.com/getkin/kin-openapi)