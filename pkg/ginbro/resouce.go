package ginbro

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"strings"
)

type Property struct {
	ColumnName string
	//DataType      string
	//ColumnType    string
	ColumnComment    string
	ColumnCommentTag []ColumnCommentTag
	ModelProp        string
	ModelType        string
	ModelTag         string
	//ColumnKey     string
	//ColumnType   string
	SwaggerType   string
	SwaggerFormat string
	IsAuthColumn  bool
}

type ColumnCommentTag struct {
	State   string
	Tag     string
	Remarks string
}

type Resource struct {
	ResourceName         string
	SimpleResourceName   string
	HandlerName          string
	TableName            string
	ModelName            string
	Properties           []Property
	IsAuthTable          bool
	PasswordPropertyName string
	PasswordColumnName   string
	AppPkg               string
	HasId                bool
}

func newResource(tableName, authTable, passwordColumn string, props []Property) Resource {
	modelName := strcase.ToCamel(tableName)
	modelName = inflection.Singular(modelName)
	resourceName := strcase.ToKebab(modelName)
	simpleResourceName := resourceName
	if simpleResourceName[0:3] == "tb-" {
		simpleResourceName = simpleResourceName[3:]
	}
	handlerName := strcase.ToLowerCamel(modelName)

	isAuthTable := tableName == authTable
	passwordPropName := strcase.ToCamel(passwordColumn)
	hasId := false
	for _, prop := range props {
		if strings.ToLower(prop.ColumnName) == "id" {
			hasId = true
			break
		}
	}
	return Resource{ModelName: modelName, TableName: tableName, ResourceName: resourceName, SimpleResourceName: simpleResourceName, HandlerName: handlerName, IsAuthTable: isAuthTable, PasswordColumnName: passwordColumn, PasswordPropertyName: passwordPropName, Properties: props, HasId: hasId}
}
func transformToResources(cols []ColumnSchema, authTable, passwordColumn string) ([]Resource, error) {
	tableMap := map[string][]Property{}
	for _, col := range cols {
		p := col.toProperty(authTable, passwordColumn)

		if props, ok := tableMap[col.TableName]; ok {
			tableMap[col.TableName] = append(props, p)
		} else {
			tableMap[col.TableName] = []Property{p}
		}

	}
	var list []Resource
	for tableName, props := range tableMap {
		resource := newResource(tableName, authTable, passwordColumn, props)
		list = append(list, resource)
	}
	return list, nil
}
