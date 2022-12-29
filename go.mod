module main

go 1.19

require (
	connections v1.2.3
	controllers v1.2.3
	entities v1.2.3
	repositories v1.2.3
	services v1.2.3
	docs v1.2.3
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.2.0 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/microsoft/go-mssqldb v0.17.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/swaggo/files v1.0.0 // indirect
	github.com/swaggo/http-swagger v1.3.3 // indirect
	github.com/swaggo/swag v1.8.9 // indirect
	github.com/urfave/cli/v2 v2.23.7 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/sqlserver v1.4.1 // indirect
	gorm.io/gorm v1.24.2 // indirect
)

replace entities v1.2.3 => ./data/entities

replace controllers v1.2.3 => ./presentation/controllers

replace services v1.2.3 => ./business/services

replace repositories v1.2.3 => ./data/repositories

replace connections v1.2.3 => ./data/connections

replace docs v1.2.3 => ./docs
