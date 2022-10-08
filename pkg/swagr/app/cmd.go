package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/gen"
)

var c gen.Config

// RestCmd represents the rest command
var RestCmd = &cobra.Command{
	Use:     "init",
	Version: fmt.Sprintf("swagr: 1.0 (powered by swag %s)", swag.Version),
	Short:   "Automatically generate RESTful API documentation with Swagger 2.0 and convert to openapi 3.0 (spec3) for Go. ",
	Example: "swagr init -g http/api.go",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		switch c.PropNamingStrategy {
		case swag.CamelCase, swag.SnakeCase, swag.PascalCase:
		default:
			return fmt.Errorf("not supported %s propertyStrategy", c.PropNamingStrategy)
		}

		return Build(&c)
	},
}

func init() {
	RestCmd.Flags().StringVarP(&c.MainAPIFile, "generalInfo", "g", "main.go", "Go file path in which 'swagger general API Info' is written")
	RestCmd.Flags().StringVarP(&c.SearchDir, "dir", "d", "./", "Directory you want to parse")
	RestCmd.Flags().StringVar(&c.Excludes, "exclude", "", "Exclude directories and files when searching, comma separated")
	RestCmd.Flags().StringVarP(&c.PropNamingStrategy, "propertyStrategy", "p", "camelcase", "Property Naming Strategy like snakecase,camelcase,pascalcase")
	RestCmd.Flags().StringVarP(&c.OutputDir, "output", "o", "./docs/swagger", "Output directory for all the generated files(swagger.json, swagger.yaml)")
	RestCmd.Flags().BoolVar(&c.ParseVendor, "parseVendor", false, "Parse go files in 'vendor' folder, disabled by default")
	RestCmd.Flags().BoolVar(&c.ParseDependency, "parseDependency", false, "Parse go files in outside dependency folder, disabled by default")
	RestCmd.Flags().StringVarP(&c.MarkdownFilesDir, "markdownFiles", "m", "", "Parse folder containing markdown files to use as description, disabled by default")
	RestCmd.Flags().StringVarP(&c.MarkdownFilesDir, "codeExampleFiles", "c", "", "Parse folder containing code example files to use for the x-codeSamples extension, disabled by default")
	RestCmd.Flags().BoolVar(&c.ParseInternal, "parseInternal", false, "Parse go files in internal packages, disabled by default")
	RestCmd.Flags().IntVar(&c.ParseDepth, "parseDepth", 100, "Dependency parse depth")
}
