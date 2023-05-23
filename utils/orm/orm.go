package orm

import "RudderMaster/database"

func DataBaseAutoMigrates(dst ...interface{}) {
	database.DB.AutoMigrate(dst...)
}
