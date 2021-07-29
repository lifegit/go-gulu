/**
* @Author: TheLife
* @Date: 2021/7/23 上午10:52
 */
package app

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/ghodss/yaml"
	"github.com/go-openapi/spec"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/gen"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func Build(config *gen.Config) error {
	if _, err := os.Stat(config.SearchDir); os.IsNotExist(err) {
		return fmt.Errorf("dir: %s is not exist", config.SearchDir)
	}

	p, err := BuildSwag(config)
	if err != nil {
		return err
	}
	v2, err := BuildV2(config, p)
	if err != nil {
		return err
	}
	_, err = BuildV3(config, v2)
	if err != nil {
		return err
	}
	return nil
}

func BuildSwag(config *gen.Config) (s *spec.Swagger, err error) {
	p := swag.New(swag.SetMarkdownFileDirectory(config.MarkdownFilesDir),
		swag.SetExcludedDirsAndFiles(config.Excludes),
		swag.SetCodeExamplesDirectory(config.CodeExampleFilesDir))
	p.PropNamingStrategy = config.PropNamingStrategy
	p.ParseVendor = config.ParseVendor
	p.ParseDependency = config.ParseDependency
	p.ParseInternal = config.ParseInternal
	err = p.ParseAPI(config.SearchDir, config.MainAPIFile, config.ParseDepth)
	if err != nil {
		return nil, err
	}
	s = p.GetSwagger()

	return s, err
}

func ReplaceSpot(searchIn []byte) (res []byte, err error) {
	replaceFunc := func(pat string, searchIn []byte) (res []byte, err error) {
		if ok, err := regexp.Match(pat, searchIn); !ok {
			return nil, err
		}

		re, err := regexp.Compile(pat)
		if err != nil {
			return nil, err
		}
		str2 := re.ReplaceAllStringFunc(string(searchIn), func(s string) string {
			return strings.ReplaceAll(s, ".", "")
		})

		return []byte(str2), nil
	}

	res, pats := searchIn, []string{
		`"\$ref": "#/definitions/(.*)"`,
		`"definitions": {([\s\S]*) `,
	}
	for _, pat := range pats {
		res, err = replaceFunc(pat, res)
		if err != nil {
			return
		}
	}

	return
}

func BuildV2(config *gen.Config, s *spec.Swagger) (res []byte, err error) {
	log.Println("Generate swagger 2 docs....")
	swaggerV2, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return
	}
	swaggerV2,err = ReplaceSpot(swaggerV2)
	if err != nil {
		return
	}

	return swaggerV2, OutPut(swaggerV2, path.Join(config.OutputDir, "v2"), "swagger")
}

func BuildV3(config *gen.Config, swaggerV2 []byte) (res []byte, err error) {
	log.Println("Generate openapi 3 docs....")
	var doc2 openapi2.T
	if err = json.Unmarshal(swaggerV2, &doc2); err != nil {
		return
	}
	spec3, err := openapi2conv.ToV3(&doc2)
	if err != nil {
		return
	}
	openapiV3, err := json.MarshalIndent(spec3, "", "    ")
	if err != nil {
		return
	}

	return openapiV3, OutPut(openapiV3, path.Join(config.OutputDir, "v3"), "openapi")
}

func OutPut(jsonBytes []byte, outputDir string, name string) error {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	// json
	jsonFileName := filepath.Join(outputDir, fmt.Sprintf("%s.json", name))
	err := file.WriteFile(jsonBytes, jsonFileName)
	if err != nil {
		return err
	}

	// yaml
	yamlFileName := filepath.Join(outputDir, fmt.Sprintf("%s.yaml", name))
	yamlBytes, err := yaml.JSONToYAML(jsonBytes)
	if err != nil {
		return fmt.Errorf("cannot convert json to yaml error: %s", err)
	}
	err = file.WriteFile(yamlBytes, yamlFileName)
	if err != nil {
		return err
	}

	log.Printf("create %s.json at %+v", name, jsonFileName)
	log.Printf("create %s.yaml at %+v", name, yamlFileName)

	return nil
}
