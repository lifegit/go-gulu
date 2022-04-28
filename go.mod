module github.com/lifegit/go-gulu/v2

go 1.16

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1168
	github.com/aliyun/aliyun-oss-go-sdk v2.1.8+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/chromedp/cdproto v0.0.0-20210526005521-9e51b9051fd0
	github.com/chromedp/chromedp v0.7.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/getkin/kin-openapi v0.66.0
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-openapi/spec v0.19.14
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.10.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gocolly/colly/v2 v2.1.0
	github.com/imdario/mergo v0.3.12
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.2.3 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.8.1
	github.com/speps/go-hashids v1.0.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/wenzhenxi/gorsa v0.0.0-20210524035706-528c7050d703
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	gorm.io/datatypes v1.0.6
	gorm.io/driver/mysql v1.3.2
	gorm.io/driver/postgres v1.3.1
	gorm.io/driver/sqlite v1.3.1
	gorm.io/driver/sqlserver v1.3.1
	gorm.io/gorm v1.23.2
	gorm.io/plugin/soft_delete v1.1.0
)

replace gorm.io/datatypes v1.0.6 => github.com/lifegit/datatypes v1.0.7
