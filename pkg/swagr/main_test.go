package main_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/pkg/swagr/app"
	"github.com/swaggo/swag/gen"
	"testing"
)

func TestName(t *testing.T) {
	err := app.Build(&gen.Config{
		ParseDepth:         100,
		MainAPIFile:        "main.go",
		PropNamingStrategy: "camelcase",
		SearchDir:          "./example/basic",
		OutputDir:          "./example/basic/docs/swagger",
	})
	fmt.Println(err)
}
