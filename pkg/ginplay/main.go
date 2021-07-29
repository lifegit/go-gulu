/**
* @Author: TheLife
* @Date: 2021/7/19 下午3:03
 */
package main

import (
	"github.com/lifegit/go-gulu/v2/pkg/ginplay/app"
	"log"
)

//go:generate go build -o main
func main() {
	if err := app.RestCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}