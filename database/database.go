package database

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/sqlite"

 var(
	DBConn *gorm.DB
 )