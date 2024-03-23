module stvsljl.com/CSMS/db

go 1.22.1

require gorm.io/gorm v1.25.8

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require stvsljl.com/CSMS/utils v0.0.0

replace stvsljl.com/CSMS/utils v0.0.0 => ../utils

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.5.6
)
